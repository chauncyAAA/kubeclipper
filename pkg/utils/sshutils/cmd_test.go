package sshutils

import (
	"testing"
)

func TestCmdV2(t *testing.T) {
	ret, err := CmdToString("ls", "-al")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestIsFileExistV2(t *testing.T) {
	exists, err := IsFileExist("123")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(exists)
}

func TestWhoami(t *testing.T) {
	whoami := Whoami()
	t.Log(whoami)
}

func TestRunCmdAsSSH(t *testing.T) {
	ret, err := RunCmdAsSSH("whoami1")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.String())
}
