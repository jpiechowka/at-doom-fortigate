package networking

import (
	"at-doom-fortigate/config"
	"at-doom-fortigate/logging"
	"crypto/tls"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var torHttpClient *http.Client = nil

func init() {
	// Logger initialization here because this init() block is called before main() function body
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	torHttpClient = initializeHttpClientWithTorTransport()
}

func GetRequestThroughTor(url string) *MiniResponseObject {
	// We only perform HTTPS requests
	if !strings.Contains(strings.ToLower(url), strings.ToLower("https://")) {
		url = "https://" + url
	}

	// Create empty response object
	getMiniResponseObject := MiniResponseObject{}

	// Perform GET request via our TOR HTTP client
	resp, reqError := torHttpClient.Get(url)

	if reqError != nil {
		log.Error().Err(reqError).Msg("Error when performing HTTP GET request")
		getMiniResponseObject = MiniResponseObject{
			RequestError: reqError,
		}

		return &getMiniResponseObject
	}

	defer func() {
		respBodyCloseErr := resp.Body.Close()
		logging.CheckFatalError(respBodyCloseErr, "Error when closing HTTP response body in defer block")
	}()

	// Read response body
	respBody, respBodyReadError := ioutil.ReadAll(resp.Body)
	logging.CheckFatalError(respBodyReadError, "Error when reading HTTP response body")

	getMiniResponseObject = MiniResponseObject{
		RequestUrl:     url,
		HttpStatusCode: resp.StatusCode,
		Headers:        resp.Header,
		ResponseBody:   respBody,
	}

	return &getMiniResponseObject
}

func initializeHttpClientWithTorTransport() *http.Client {
	log.Info().Msg("Initializing HTTP client with TOR proxy for the first time")
	log.Info().Msg("Creating TOR transport for HTTP client")

	// Parse Tor proxy URL string to a URL type
	log.Info().
		Str("networking-proxy-string", config.TorProxyString).
		Msg("Parsing TOR proxy string")
	torProxyUrl, urlParseErr := url.Parse(config.TorProxyString)
	logging.CheckFatalError(urlParseErr, "Error when parsing TOR proxy string")

	log.Warn().Msg("Disabling TLS certificate checks! Fortigate certs are not considered legitimate/valid so we " +
		"will not validate them.")

	log.Info().Msg("Disabling compression and HTTP keep-alives because of TOR")

	torHttpTransport := &http.Transport{
		Proxy:              http.ProxyURL(torProxyUrl),
		DisableKeepAlives:  true,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	log.Info().
		Dur("timeout-millis", config.HttpClientTimeout).
		Msg("Setting HTTP client timeout")

	torHttpClient = &http.Client{
		Timeout:   config.HttpClientTimeout,
		Transport: torHttpTransport,
	}

	log.Info().Msg("HTTP client initialized")

	return torHttpClient
}
