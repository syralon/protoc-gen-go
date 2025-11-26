package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/syralon/protoc-gen-go/internal/protogen/errors"
	"github.com/syralon/protoc-gen-go/pkg/protoc"
	"google.golang.org/protobuf/compiler/protogen"
)

var VERSION = "v0.0.1"

func main() {
	v := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if v != nil && *v {
		_, file := path.Split(strings.ReplaceAll(os.Args[0], "\\", "/"))
		fmt.Printf("%s %s", strings.Split(file, ".")[0], VERSION)
		os.Exit(0)
	}

	opts := protoc.NewOptions()
	builders := protoc.BuildChain{
		&errors.ErrorBuilder{},
	}
	protogen.Options{ParamFunc: opts.Set}.Run(builders.Build(opts))
}
