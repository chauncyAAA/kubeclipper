package sshutils

import (
	"fmt"
	"strings"
)

func fillCmd(sshConfig *SSH, cmd string) (string, error) {
	if sshConfig.User == "root" {
		return cmd, nil
	}
	// note: not work on some command,such as 'echo "&&"'
	split := strings.Split(cmd, "&&")
	list := make([]string, 0, len(split))
	for _, v := range split {
		var sudoCmd string
		if sshConfig.Password != "" {
			// use `echo '$passwd' |sudo -S $cmd` cmd avoid interactive enter passwd
			sudoCmd = fmt.Sprintf("echo '%s' | sudo -S %s", sshConfig.Password, v)
		} else {
			// no passwd maybe user configured NOPASSWD in sudoers.
			sudoCmd = "sudo " + v
		}
		list = append(list, sudoCmd)
	}
	return strings.Join(list, "&&"), nil
}

// printCmd replace sensitive information in cmd
func printCmd(passwd, cmd string) string {
	if passwd == "" {
		return cmd
	}
	return strings.ReplaceAll(cmd, fmt.Sprintf("echo '%s'", passwd), "echo $PASSWD")
}

// WrapEcho use sh -c to wrap echo and > in one cmd for sudo.
func WrapEcho(data, dest string) string {
	data = strings.ReplaceAll(data, `"`, `\"`)
	return fmt.Sprintf("/bin/bash -c \"%s\"", fmt.Sprintf("echo '%s' > %s", data, dest))
}

// WrapSh use sh -c to wrap cmd for sudo.
func WrapSh(cmd string) string {
	cmd = strings.ReplaceAll(cmd, `"`, `\"`)
	return fmt.Sprintf("/bin/bash -c \"%s\"", cmd)
}

func Combine(cmdList []string) string {
	return strings.Join(cmdList, " && ")
}
