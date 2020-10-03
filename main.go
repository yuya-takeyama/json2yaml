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

const appName = "json2yaml"

type options struct {
	ShowVersion bool `short:"v" long:"version" description:"Show version"`
}

func main() {
	var opts options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = appName
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

	err = json2yaml(r, os.Stdout, opts)
	if err != nil {
		panic(err)
	}
}

func json2yaml(r io.Reader, stdout io.Writer, opts options) error {
	if opts.ShowVersion {
		_, _ = io.WriteString(stdout, fmt.Sprintf("%s v%s, build %s\n", appName, Version, GitCommit))
		return nil
	}

	decoder := json.NewDecoder(r)

	var d interface{}

	for {
		if err := decoder.Decode(&d); err == io.EOF {
			break
		} else if err != nil {
			return err
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
