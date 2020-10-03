package main

import (
	"bytes"
	"errors"
	"os/exec"
	"testing"
)

func TestSingleFileWithSingleJSON(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--", "fixtures/single.json")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		t.Errorf("failed: %v", err)
	}

	expected := "---\nfoo: FOO\n"
	if stdout.String() != expected {
		t.Errorf("stdout doesn't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
	}

	if stderr.String() != "" {
		t.Errorf("stderr should be empty:\n%s", stderr.String())
	}
}

func TestSingleFileWithMultipleJSONs(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--", "fixtures/multiple.json")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		t.Errorf("failed: %v", err)
	}

	expected := `---
foo: FOO
---
bar: BAR
---
baz: BAZ
`
	if stdout.String() != expected {
		t.Errorf("stdout doesn't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
	}

	if stderr.String() != "" {
		t.Errorf("stderr should be empty:\n%s", stderr.String())
	}
}

func TestMultipleFilesWithSingleJSON(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--", "fixtures/single.json", "fixtures/single.json", "fixtures/single.json")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		t.Errorf("failed: %v", err)
	}

	expected := `---
foo: FOO
---
foo: FOO
---
foo: FOO
`
	if stdout.String() != expected {
		t.Errorf("stdout doesn't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
	}

	if stderr.String() != "" {
		t.Errorf("stderr should be empty:\n%s", stderr.String())
	}
}

func TestMultipleFilesWithMultipleJSONs(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--", "fixtures/multiple.json", "fixtures/multiple.json", "fixtures/multiple.json")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		t.Errorf("failed: %v", err)
	}

	expected := `---
foo: FOO
---
bar: BAR
---
baz: BAZ
---
foo: FOO
---
bar: BAR
---
baz: BAZ
---
foo: FOO
---
bar: BAR
---
baz: BAZ
`
	if stdout.String() != expected {
		t.Errorf("stdout doesn't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
	}

	if stderr.String() != "" {
		t.Errorf("stderr should be empty:\n%s", stderr.String())
	}
}

func TestVersion(t *testing.T) {
	cmd := exec.Command("go", "run", "-ldflags", "-X main.version=1.2.3 -X main.gitCommit=deadbeef", "main.go", "--version")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	expected := "json2yaml v1.2.3, build deadbeef\n"
	if err := cmd.Run(); err != nil {
		t.Errorf("failed: %v", err)
	}

	if stdout.String() != expected {
		t.Errorf("stdout doesn't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
	}
}

func TestHelp(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--help")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	expected := `Usage:
  json2yaml [OPTIONS] FILES...

Application Options:
  -v, --version  Show version

Help Options:
  -h, --help     Show this help message
`
	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			t.Errorf("failed: %v", err)
		}
	}

	if stdout.String() != "" {
		t.Errorf("stdout should be empty")
	}

	if stderr.String() != expected {
		t.Errorf("stderr doesn't match\nExpected: \n%s\nActual:\n%s", expected, stderr.String())
	}
}

func TestUnknownOption(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--foo")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	expected := "json2yaml: flag parse error: unknown flag `foo'\nexit status 1\n"
	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			t.Errorf("failed: %v", err)
		}
	}

	if stdout.String() != "" {
		t.Errorf("stdout should be empty")
	}

	if stderr.String() != expected {
		t.Errorf("stderr doesn't match\nExpected: \n%s\nActual:\n%s", expected, stderr.String())
	}
}

func TestFileNotExists(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "not_exist.json")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	expected := "json2yaml: file loading error: open not_exist.json: no such file or directory\nexit status 1\n"
	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			t.Errorf("failed: %v", err)
		}
	}

	if stdout.String() != "" {
		t.Errorf("stdout should be empty")
	}

	if stderr.String() != expected {
		t.Errorf("stderr doesn't match\nExpected: \n%s\nActual:\n%s", expected, stderr.String())
	}
}

func TestInvalidJSON(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "fixtures/invalid.json")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	expected := "json2yaml: error: invalid character '0' looking for beginning of object key string\nexit status 1\n"
	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			t.Errorf("failed: %v", err)
		}
	}

	if stdout.String() != "" {
		t.Errorf("stdout should be empty")
	}

	if stderr.String() != expected {
		t.Errorf("stderr doesn't match\nExpected: \n%s\nActual:\n%s", expected, stderr.String())
	}
}
