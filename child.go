package main

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/progrium/qtalk-go/codec"
	"github.com/progrium/qtalk-go/fn"
	"github.com/progrium/qtalk-go/mux"
	"github.com/progrium/qtalk-go/rpc"
)

func init() {
	runtime.LockOSThread()
}

type child struct{}

// Change string to Upper case
func (c *child) Upper(s string) string {
	return strings.ToUpper(s)
}

func subprogram() {
	sess, err := mux.DialIO(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	srv := &rpc.Server{
		Codec:   codec.JSONCodec{},
		Handler: fn.HandlerFrom(new(child)),
	}

	go srv.Respond(sess)
}
