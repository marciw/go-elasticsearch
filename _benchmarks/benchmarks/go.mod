module github.com/elastic/go-elasticsearch/v8/benchmarks

go 1.22

toolchain go1.22.0

replace github.com/elastic/go-elasticsearch/v8 => ../../

require (
	github.com/elastic/elastic-transport-go/v8 v8.6.1
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20200408073057-6f36a473b19f
	github.com/fatih/color v1.7.0
	github.com/montanaflynn/stats v0.6.3
	github.com/tidwall/gjson v1.9.3
)

require github.com/mattn/go-colorable v0.1.6 // indirect
