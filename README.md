# GoRH
Simple Alphavantage stock market API scraper written in Go


`go get github.com/anaskhan96/soup`
`gun test2.go | grep -Eo '\\"price\\",\\"[0-9]+\.[0-9]+' --color=yes | less` < -- Debug stuff
`gun test2.go | grep -Eo '\\"price\\",\\"[0-9]+\.[0-9]+' | awk -F '"' '{print $(NF)}'`
