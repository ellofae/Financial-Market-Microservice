module github.com/ellofae/Financial-Market-Microservice/ClientSide

go 1.20

replace github.com/ellofae/Financial-Market-Microservice/CurrencyRates => ../CurrencyRates

require (
	github.com/ellofae/Financial-Market-Microservice/CurrencyRates v0.0.0-00010101000000-000000000000
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-hclog v1.5.0
	google.golang.org/grpc v1.55.0
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
