package sshutils

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func MD5FromLocal(localPath string) (string, error) {
	cmd := fmt.Sprintf("md5sum %s | cut -d\" \" -f1", localPath)
	c := exec.Command("sh", "-c", cmd)
	var bout, berr bytes.Buffer
	c.Stdout, c.Stderr = &bout, &berr
	err := c.Run()
	if err != nil {
		return "", err
	}
	stderr := berr.String()
	if stderr != "" {
		return "", errors.New(stderr)
	}
	md5 := bout.String()
	md5 = strings.ReplaceAll(md5, "\n", "")
	md5 = strings.ReplaceAll(md5, "\r", "")
	return md5, nil
}

func (ss *SSH) MD5FromRemote(host, remoteFilePath string) (string, error) {
	cmd := fmt.Sprintf("md5sum %s | cut -d\" \" -f1", remoteFilePath)
	ret, err := SSHCmdWithSudo(ss, host, cmd)
	if err != nil {
		return "", err
	}
	if err = ret.Error(); err != nil {
		return "", err
	}
	md5 := ret.StdoutToString("")
	return md5, nil
}

func (ss *SSH) ValidateMd5sumLocalWithRemote(host, localFile, remoteFile string) (bool, error) {
	localMD5, err := MD5FromLocal(localFile)
	if err != nil {
		return false, err
	}
	remoteMD5, err := ss.MD5FromRemote(host, remoteFile)
	if err != nil {
		return false, err
	}
	return localMD5 == remoteMD5, nil
}
