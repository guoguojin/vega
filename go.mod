module code.vegaprotocol.io/data-node

go 1.16

require (
	code.vegaprotocol.io/protos v0.41.1-0.20210817173950-833e9e05f293
	code.vegaprotocol.io/quant v0.2.5
	code.vegaprotocol.io/vega v0.41.1-0.20210813173752-c4e56dfda378
	github.com/99designs/gqlgen v0.13.0
	github.com/dgraph-io/badger/v2 v2.2007.2
	github.com/fsnotify/fsnotify v1.4.9
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/oasisprotocol/curve25519-voi v0.0.0-20210716083614-f38f8e8b0b84
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/cors v1.7.0
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/vektah/gqlparser/v2 v2.1.0
	github.com/zannen/toml v0.3.2
	go.elastic.co/apm/module/apmhttp v1.8.0
	go.nanomsg.org/mangos/v3 v3.2.1
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/shopspring/decimal => github.com/vegaprotocol/decimal v1.2.1-0.20210705145732-aaa563729a0a
