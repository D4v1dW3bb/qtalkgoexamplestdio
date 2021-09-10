package shell

import "github.com/d4v1dw3bb/qtalkgoexample/handle"

type Case struct {
	handle.Handle

	Text      string
	UpperCase bool
}
