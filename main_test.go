package main

import (
	"bytes"
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
		t.Errorf("stdout doen't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
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
		t.Errorf("stdout doen't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
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
		t.Errorf("stdout doen't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
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
		t.Errorf("stdout doen't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
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
		t.Errorf("stdout doen't match\nExpected:\n%s\nActual:\n%s", expected, stdout.String())
	}
}
