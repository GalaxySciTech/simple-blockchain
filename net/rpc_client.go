package net

import (
	"encoding/json"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"simple-blockchain/config"
	"simple-blockchain/core"
)

type Pie struct {
}

func (pie *Pie) GetBlockChainInfo(arg string, reply *string) error {
	block := core.Bc
	jsonBytes, _ := json.Marshal(block)
	*reply = string(jsonBytes)
	return nil
}

func Start() {

	rpc.RegisterName("Pie", new(Pie))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":"+config.Settings.Port, nil)
}
