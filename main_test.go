package main

import (
	"os"
	"testing"
)

func TestUID(t *testing.T) {
	uid := os.Geteuid()
	if uid != 0 {
		t.Skip("run as root")
	}
	t.Log("not skip: run as root")
}
