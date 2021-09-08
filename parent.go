package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/progrium/qtalk-go/codec"
	"github.com/progrium/qtalk-go/fn"
	"github.com/progrium/qtalk-go/mux"
	"github.com/progrium/qtalk-go/talk"
)

// Parent holds all parent properties
type Parent struct {
	Debug io.Writer
	peer  *talk.Peer
	cmd   *exec.Cmd
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Run program as a parent
func parentprogram() {
	cmd := exec.Command("qtalkgoexample", "-type=child")
	wc, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	rc, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	sess, err := mux.DialIO(wc, rc)
	if err != nil {
		panic(err)
	}

	p := &Parent{
		peer: talk.NewPeer(sess, codec.JSONCodec{}),
		cmd:  cmd,
	}
	p.Debug = os.Stderr
	if p.cmd != nil {
		p.cmd.Stderr = p.Debug
		if p.Debug != nil {
			p.cmd.Args = append(p.cmd.Args, "-debug")
		}
		must(p.cmd.Start())
	}
	go p.peer.Respond()
	ctx := context.Background()
	var ret interface{}
	_, err = p.peer.Call(ctx, "Upper", fn.Args{"hello world"}, &ret)
	must(err)
	log.Println(ret)
}
