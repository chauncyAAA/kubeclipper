package sshutils

import (
	"fmt"
	"path"
	"strings"
)

func (ss *SSH) IsFileExistV2(host, remoteFilePath string) (bool, error) {
	// if remote file is
	// ls -l | grep aa | wc -l
	remoteFileName := path.Base(remoteFilePath) // aa
	remoteFileDirName := path.Dir(remoteFilePath)
	// it's bug: if file is aa.bak, `ls -l | grep aa | wc -l` is 1 ,should use `ll aa 2>/dev/null |wc -l`
	// remoteFileCommand := fmt.Sprintf("ls -l %s| grep %s | grep -v grep |wc -l", remoteFileDirName, remoteFileName)
	remoteFileCommand := fmt.Sprintf("ls -l %s/%s 2>/dev/null |wc -l", remoteFileDirName, remoteFileName)
	ret, err := SSHCmdWithSudo(ss, host, remoteFileCommand)
	if err != nil {
		return false, err
	}
	if err = ret.Error(); err != nil {
		return false, err
	}
	data := ret.StdoutToString(" ")
	return strings.TrimSpace(data) != "0", nil
}
