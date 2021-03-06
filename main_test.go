package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("capze --version", " ")

	status := cli.Run(args)
	if status != 0 {
		t.Errorf("expected %d to eq %d", status, 0)
	}

	expected := fmt.Sprintf("capze version %s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_parseError(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("capze --not-exist", " ")

	status := cli.Run(args)
	if status != 10 {
		t.Errorf("expected %d to eq %d", status, 10)
	}

	expected := "flag provided but not defined"
	if !strings.Contains(errStream.String(), expected) {
		t.Fatalf("expected %q to contain %q", errStream.String(), expected)
	}
}
