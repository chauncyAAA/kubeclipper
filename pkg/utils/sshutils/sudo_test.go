package sshutils

import "testing"

func Test_printCmd(t *testing.T) {
	cmd := printCmd("root", "echo 'root'|sudo -S id -u")
	if cmd != "echo '$PASSWD'|sudo -S id -u" {
		t.Failed()
	}
}

func TestBuildEcho(t *testing.T) {
	echo := WrapEcho("hello world", "/usr/lib/systemd/system/kc-etcd.service")
	if echo != `/bin/bash -c "echo 'hello world' > /usr/lib/systemd/system/kc-etcd.service"` {
		t.Failed()
	}
}
