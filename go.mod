module github.com/linbeier/authsys

go 1.13

require (
	github.com/coreos/go-etcd v2.0.0+incompatible // indirect
	github.com/cosmos/cosmos-sdk v0.38.4
	github.com/cpuguy83/go-md2man v1.0.10 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.4
	github.com/tendermint/tm-db v0.5.1
	github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8 // indirect
)

replace github.com/cosmos/cosmos-sdk v0.38.4 => github.com/cosmos/cosmos-sdk v0.34.4-0.20200511222341-80be50319ca5
