package parent

import (
	"log"
	"os"

	"github.com/d4v1dw3bb/qtalkgoexample/shell"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Run program as a parent
func Run() {
	sh := shell.New(nil)
	sh.Debug = os.Stderr
	must(sh.Open())
	defer sh.Close()
	w := shell.Case{
		Text:      "TestThis",
		UpperCase: true,
	}
	must(sh.Sync(&w))
	if err := sh.Wait(); err != nil {
		log.Fatal(err)
	}
}
