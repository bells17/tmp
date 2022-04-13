package main

import (
	"os"
	"testing"
)

func TestUID(t *testing.T) {
	uid := os.Getuid()
	if uid != 0 {
		t.Skip("run as root")
	}
	t.Log("not skip: run as root")
}
