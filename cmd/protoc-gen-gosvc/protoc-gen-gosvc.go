package main

import (
	"github.com/jemtucker/protogengosvc"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(p *protogen.Plugin) error {
		for _, f := range p.Files {
			if f.Generate {
				gen := protogengosvc.Protogen{Generator: p}
				if err := gen.Generate(f); err != nil {
					return err
				}
			}
		}

		return nil
	})
}
