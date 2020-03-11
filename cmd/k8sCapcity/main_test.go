package main

import (
	"os"
	"testing"
)

func TestHomeDirLinux(t *testing.T) {
	h := homeDir()
	if h == "" {
		t.Errorf("Could not find homedir\n")
	}
}

func TestHomeDirWindows(t *testing.T) {
	os.Setenv("HOME", "")
	os.Setenv("USERPROFILE", "C:\\Program Files\\Masters of Orion")
	h := homeDir()
	if h == "" {
		t.Errorf("Could not find homedir\n")
	}
}
