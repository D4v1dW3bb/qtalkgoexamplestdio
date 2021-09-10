package child

import (
	"log"
	"os"
	"runtime"

	"github.com/d4v1dw3bb/qtalkgoexample/bridge"
	"github.com/progrium/qtalk-go/mux"
)

const Version = "0.1.0"

func init() {
	runtime.LockOSThread()
}

func Run() {
	sess, err := mux.DialIO(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	srv := bridge.NewServer()
	go srv.Respond(sess)
}
