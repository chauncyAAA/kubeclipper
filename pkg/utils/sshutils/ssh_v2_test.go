package sshutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSSHCmd(t *testing.T) {
	defer func() {
		Cmd("rm", "/tmp/d.txt")
		Cmd("rm", "/tmp/dd.txt")
	}()
	ret, err := SSHCmd(nil, "", "bash -c \"echo 123\"")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "123", ret.StdoutToString(""))

	_, err = SSHCmd(nil, "", WrapEcho("12345", "/tmp/d.txt"))
	if err != nil {
		t.Fatal(err)
	}
	ret, err = RunCmdAsSSH("cat /tmp/d.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "12345", ret.StdoutToString(""))

	_, err = SSHCmdWithSudo(nil, "", WrapEcho("54321", "/tmp/dd.txt"))
	if err != nil {
		t.Fatal(err)
	}
	ret, err = RunCmdAsSSH("cat /tmp/dd.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "54321", ret.StdoutToString(""))
}

func TestSSHTable(t *testing.T) {
	ret, err := SSHCmdWithSudo(nil, "", "ls .")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret.Table())
}
