package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/yuya-takeyama/argf"
	"gopkg.in/yaml.v3"
)

const AppName = "json2yaml"

type Options struct {
	ShowVersion bool `short:"v" long:"version" description:"Show version"`
}

var opts Options

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = AppName
	parser.Usage = "[OPTIONS] FILES..."

	args, err := parser.Parse()
	if err != nil {
		fmt.Print(err)
		return
	}

	r, err := argf.From(args)
	if err != nil {
		panic(err)
	}

	err = json2yaml(r, os.Stdout, os.Stderr, opts)
	if err != nil {
		panic(err)
	}
}

func json2yaml(r io.Reader, stdout io.Writer, stderr io.Writer, opts Options) error {
	if opts.ShowVersion {
		io.WriteString(stdout, fmt.Sprintf("%s v%s, build %s\n", AppName, Version, GitCommit))
		return nil
	}

	decoder := json.NewDecoder(r)
	var d interface{}

	for {
		if err := decoder.Decode(&d); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		yml, err := yaml.Marshal(d)
		if err != nil {
			panic(err)
		}

		stdout.Write([]byte("---\n"))
		stdout.Write(yml)
	}

	return nil
}
