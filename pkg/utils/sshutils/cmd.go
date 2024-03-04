package sshutils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/kubeclipper/kubeclipper/pkg/cli/logger"
)

func IsFileExist(filepath string) (bool, error) {
	// ls -l $dir | grep $name | wc -l
	fileName := path.Base(filepath)
	fileDirName := path.Dir(filepath)
	fileCommand := fmt.Sprintf("ls -l %s | grep %s | wc -l", fileDirName, fileName)
	ret, err := CmdToString("/bin/sh", "-c", fileCommand)
	if err != nil {
		return false, err
	}
	if err = ret.Error(); err != nil {
		return false, err
	}
	data := strings.ReplaceAll(strings.TrimSpace(ret.Stdout), "\r", "")
	data = strings.ReplaceAll(data, "\n", "")
	return data != "0", nil
}

func CmdToString(name string, arg ...string) (Result, error) {
	var ret Result
	cmd := exec.Command(name, arg[:]...)
	ret.PrintCmd = cmd.String()
	cmd.Stdin = os.Stdin
	var bout, berr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &bout, &berr
	err := cmd.Run()
	if err != nil {
		return ret, err
	}
	ret.Stdout = bout.String()
	ret.Stderr = berr.String()

	logger.V(5).Info(ret.Table())
	return ret, ret.error()
}

func RunCmdAsSSH(cmdStr string) (Result, error) {
	var ret Result

	user := Whoami()
	ec := exec.Command("sh", []string{"-c", cmdStr}...)
	ec.Stdin = os.Stdin
	var bout, berr bytes.Buffer
	ec.Stdout, ec.Stderr = &bout, &berr
	err := ec.Run()
	ret = Result{
		User:     user,
		Host:     "localhost",
		Cmd:      ec.String(),
		PrintCmd: ec.String(),
		Stdout:   bout.String(),
		Stderr:   berr.String(),
	}
	if err != nil {
		ok, exitCode := ExtraExitCode(err)
		if !ok {
			return ret, err
		}
		// with exitCode,ignore error
		ret.ExitCode = exitCode
		return ret, nil
	}
	logger.V(5).Info(ret.Table())
	return ret, ret.error()
}

func Whoami() string {
	result, err := CmdToString("whoami")
	if err != nil {
		return ""
	}
	return result.StdoutToString("")
}

func Cmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
