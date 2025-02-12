package logutils

import (
	"bytes"
	"log"
	"regexp"
	"testing"
)

func TestGIDFilter(t *testing.T) {
	buf := new(bytes.Buffer)
	filter := &GIDFilter{
		Writer: buf,
	}

	logger := log.New(filter, "", 0)
	logger.Println("[GID]foo")

	result := buf.String()
	if match, _ := regexp.MatchString("\\[[0..9]\\]foo\\n", result); match {
		t.Fatalf("bad: %#v", result)
	}
}

func TestGIDFilterWithoutGID(t *testing.T) {
	buf := new(bytes.Buffer)
	filter := &GIDFilter{
		Writer:    buf,
		GIDString: "[NOGID]",
	}

	logger := log.New(filter, "", 0)
	logger.Println("[GID]foo")

	result := buf.String()
	if match, _ := regexp.MatchString("\\[[0..9]\\]foo\\n", result); match {
		t.Fatalf("bad: %#v", result)
	}
}
