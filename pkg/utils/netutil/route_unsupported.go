//go:build !linux
// +build !linux

package netutil

import (
	"net"
)

func GetDefaultIP(ipv4 bool, method string) (net.IP, error) {
	return net.IPv4zero, nil
}

func GetDefaultGateway(ipv4 bool) (net.IP, error) {
	return net.IPv4zero, nil
}
