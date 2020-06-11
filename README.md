# At Doom Fortigate

```
 	

=================     ===============     ===============   ========  ========
\\ . . . . . . .\\   //. . . . . . .\\   //. . . . . . .\\  \\. . .\\// . . //
||. . ._____. . .|| ||. . ._____. . .|| ||. . ._____. . .|| || . . .\/ . . .||
|| . .||   ||. . || || . .||   ||. . || || . .||   ||. . || ||. . . . . . . ||
||. . ||   || . .|| ||. . ||   || . .|| ||. . ||   || . .|| || . | . . . . .||
|| . .||   ||. _-|| ||-_ .||   ||. . || || . .||   ||. _-|| ||-_.|\ . . . . ||
||. . ||   ||-'  || ||  `-||   || . .|| ||. . ||   ||-'  || ||  `|\_ . .|. .||
|| . _||   ||    || ||    ||   ||_ . || || . _||   ||    || ||   |\ `-_/| . ||
||_-' ||  .|/    || ||    \|.  || `-_|| ||_-' ||  .|/    || ||   | \  / |-_.||
||    ||_-'      || ||      `-_||    || ||    ||_-'      || ||   | \  / |  `||
||    `'         || ||         `'    || ||    `'         || ||   | \  / |   ||
||            .===' `===.         .==='.`===.         .===' /==. |  \/  |   ||
||         .=='   \_|-_ `===. .==='   _|_   `===. .===' _-|/   `==  \/  |   ||
||      .=='    _-'    `-_  `='    _-'   `-_    `='  _-'   `-_  /|  \/  |   ||
||   .=='    _-'          `-__\._-'         `-_./__-'         `' |. /|  |   ||
||.=='    _-'                                                     `' |  /==.||
=='    _-'                                                            \/   `==
\   _-'                                                                `-_   /
 `''                                                                      ``'
  
```

Tool to search for vulnerable Fortigate hosts in Rapid7 Project Sonar data anonymously through The Tor network.

#### CVE-2018-13379

More infomration on Orange Tsai's Blog: https://blog.orange.tw/2019/08/attacking-ssl-vpn-part-2-breaking-the-fortigate-ssl-vpn.html

#### How to use
1. Visit https://youtu.be/q657rEkgfKs
2. Download Rapid7 data in json format for port 10443 from https://opendata.rapid7.com/sonar.https/
3. Place json file in /data directory (or configure input path in config.go file and recompile)
4. Download and run Tor Browser
5. Run app
6. Profit (see results.txt output file or configure output file path in config.go and recompile)

### Building from source code
To build from source execute the commands below (Go needs to be installed and properly configured, see https://golang.org/doc/install)

```
git clone https://github.com/jpiechowka/at-doom-fortigate.git
cd at-doom-fortigate
go build -v -a .
```
