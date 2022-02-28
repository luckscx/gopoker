module main

go 1.17

require logger v0.0.0

require (
	github.com/onrik/logrus v0.9.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
)

replace logger => ../logger
