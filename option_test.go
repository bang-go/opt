package opt_test

import (
	"github.com/bang-go/opt"
	"log"
	"testing"
)

func TestOption(t *testing.T) {
	optionList := []opt.Option[options]{WithName("test")}
	o := &options{}
	opt.Each[options](o, optionList...)
	log.Println(o.Name)
}

type options struct {
	Name string
}

func WithName(name string) opt.Option[options] {
	return opt.OptionFunc[options](func(o *options) {
		o.Name = name
	})
}
