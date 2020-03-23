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
	home := os.Getenv("HOME")
	windowsHome := os.Getenv("USERPROFILE")
	os.Setenv("HOME", "")
	os.Setenv("USERPROFILE", "C:\\Program Files\\Masters of Orion")
	h := homeDir()
	if h == "" {
		os.Setenv("HOME", home)
		os.Setenv("USERPROFILE", windowsHome)
		t.Errorf("Could not find homedir\n")
	}
	os.Setenv("HOME", home)
	os.Setenv("USERPROFILE", windowsHome)
}

func TestRunMain(t *testing.T) {
	main()
}
