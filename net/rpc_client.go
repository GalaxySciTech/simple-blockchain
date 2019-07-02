package net

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"simple-blockchain/config"
	"simple-blockchain/core"
)

type Pie struct {
}

func (pie *Pie) GetBlockChainInfo(arg string, reply *core.BlockChain) error {
	*reply = *core.Bc
	return nil
}

func Start() {

	rpc.RegisterName("Pie", new(Pie))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
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
