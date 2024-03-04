//go:build linux
// +build linux

package netutil

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultIP(t *testing.T) {
	// ip -4 route get 114.114.114.114
	cmd := exec.Command("bash", "-c", "ip -4 route get 114.114.114.114")
	stdOutBuf := &bytes.Buffer{}
	cmd.Stdout = stdOutBuf
	if err := cmd.Run(); err != nil {
		assert.FailNowf(t, "failed to exec ip r command", err.Error())
	}
	list := strings.Split(stdOutBuf.String(), " ")
	ip, err := GetDefaultIP(true, "")
	if err != nil {
		assert.FailNowf(t, "failed to get default IP", err.Error())
	}
	assert.Equal(t, list[6], ip.To4().String())
}

func TestGetDefaultGateway(t *testing.T) {
	// ip -4 route get 114.114.114.114
	cmd := exec.Command("bash", "-c", "ip -4 route get 114.114.114.114")
	stdOutBuf := &bytes.Buffer{}
	cmd.Stdout = stdOutBuf
	if err := cmd.Run(); err != nil {
		assert.FailNowf(t, "failed to exec ip r command", err.Error())
	}
	list := strings.Split(stdOutBuf.String(), " ")
	gw, err := GetDefaultGateway(true)
	if err != nil {
		assert.FailNowf(t, "failed to get default gateway", err.Error())
	}
	assert.Equal(t, list[2], gw.To4().String())
}
