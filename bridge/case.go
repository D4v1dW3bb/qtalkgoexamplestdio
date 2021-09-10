package bridge

import (
	"strings"
	"unicode"

	"github.com/d4v1dw3bb/qtalkgoexample/shell"
)

func init() {
	register(&Case{})
}

type Case struct {
	shell.Case ` mapstructure:",squash"`

	Result string
}

func (c *Case) Resource() interface{} {
	return &c.Case
}

func (c *Case) Discard() error {
	c.Result = ""
	return nil
}

func (c *Case) Apply() error {
	if c.Result == "" {
		if c.UpperCase {
			c.Result = strings.ToUpper(c.Text)
		} else {
			c.Result = strings.ToLower(c.Text)
		}
	} else if strings.EqualFold(c.Result, c.Text) {
		if c.UpperCase && isLower(&c.Result) {
			c.Result = strings.ToUpper(c.Text)
		} else if !c.UpperCase && isUpper(&c.Result) {
			c.Result = strings.ToLower(c.Text)
		}
	} else {
		//log.Fatal("String missmatch: c.Result && C.Text do not contain simmilar text")
		c.Result = ""
	}

	//log.Println(c.Result)

	return nil
}

func isUpper(s *string) bool {
	for _, l := range *s {
		if unicode.IsLetter(l) && !unicode.IsUpper(l) {
			return false
		}
	}
	return true
}

func isLower(s *string) bool {
	for _, l := range *s {
		if unicode.IsLetter(l) && !unicode.IsLower(l) {
			return false
		}
	}
	return true
}
