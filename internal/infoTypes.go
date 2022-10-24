package internal

import "regexp"

type InfoType struct {
	Expr string
	Name string
}

func (i InfoType) Compiled() *regexp.Regexp {
	r, _ := regexp.Compile("p([a-z]+)ch")
	return r
}
