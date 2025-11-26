package protoc

import "google.golang.org/protobuf/compiler/protogen"

type Builder interface {
	Build(plugin *protogen.Plugin, options Options) error
}

type BuildFunc func(plugin *protogen.Plugin, options Options) error

func (fn BuildFunc) Build(plugin *protogen.Plugin, options Options) error {
	return fn(plugin, options)
}

type BuildChain []Builder

func (b BuildChain) Build(options Options) func(plugin *protogen.Plugin) error {
	return func(plugin *protogen.Plugin) error {
		for _, bu := range b {
			if err := bu.Build(plugin, options); err != nil {
				return err
			}
		}
		return nil
	}
}
