package main

import (
	"flag"
	"os"
	"testing"
)

func setUpTest() {
	os.Args = append(os.Args, "-test.timeout=10m0s")
}

func tearDownTest() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = nil
}

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
	setUpTest()
	main()
	tearDownTest()
}

func TestRunMainJSON(t *testing.T) {
	setUpTest()
	os.Args = append(os.Args, "--json")
	main()
	tearDownTest()
}

func TestRunMainNodeLabelJSON(t *testing.T) {
	setUpTest()
	os.Args = append(os.Args, "--nodelabel=kubernetes.io/hostname=kind-control-plane")
	os.Args = append(os.Args, "--json")
	main()
	tearDownTest()
}

func TestRunMainCheck(t *testing.T) {
	setUpTest()
	os.Args = append(os.Args, "--check")
	main()
	tearDownTest()
}

func TestRunMainNamespace(t *testing.T) {
	setUpTest()
	os.Args = append(os.Args, "--namespace=kube-system")
	main()
	tearDownTest()
}

func TestRunMainNamespaceJSON(t *testing.T) {
	setUpTest()
	os.Args = append(os.Args, "--namespace=kube-system")
	os.Args = append(os.Args, "--json")
	main()
}
