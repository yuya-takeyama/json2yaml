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

var (
	version   string = ""
	gitCommit        = ""
)

type options struct {
	ShowVersion bool `short:"v" long:"version" description:"Show version"`
}

func main() {
	var opts options
	parser := flags.NewParser(&opts, flags.Default^flags.PrintErrors)
	parser.Name = appName
	parser.Usage = "[OPTIONS] FILES..."

	args, err := parser.Parse()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				parser.WriteHelp(os.Stderr)
				return
			}
		}

		errorf("flag parse error: %s", err)
		os.Exit(1)
	}

	r, err := argf.From(args)
	if err != nil {
		errorf("ARGF error: %s", err)
		os.Exit(1)
	}

	err = json2yaml(r, os.Stdout, opts)
	if err != nil {
		errorf("error: %s", err)
		os.Exit(1)
	}
}

func json2yaml(r io.Reader, stdout io.Writer, opts options) error {
	if opts.ShowVersion {
		_, _ = io.WriteString(stdout, fmt.Sprintf("%s v%s, build %s\n", appName, version, gitCommit))
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
			return err
		}

		stdout.Write([]byte("---\n"))
		stdout.Write(yml)
	}

	return nil
}

func errorf(message string, args ...interface{}) {
	subMessage := fmt.Sprintf(message, args...)
	_, _ = fmt.Fprintf(os.Stderr, "json2yaml: %s\n", subMessage)
}
