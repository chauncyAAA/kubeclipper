//go:build !linux
// +build !linux

package sysutil

import (
	"strings"

	"github.com/shirou/gopsutil/v3/net"
)

func netInfo() ([]Net, error) {
	inters, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	nets := make([]Net, 0)

	for _, v := range inters {
		if v.Name == Loopback {
			continue
		}
		nets = append(nets, Net{
			Index:        v.Index,
			Name:         v.Name,
			MTU:          v.MTU,
			HardwareAddr: v.HardwareAddr,
			Addrs:        AddrsConvert(v.Addrs),
		})
	}
	return nets, err
}

func AddrsConvert(list net.InterfaceAddrList) []InterfaceAddr {
	inters := make([]InterfaceAddr, len(list))
	for i, item := range list {
		inters[i].Addr = item.Addr
		if strings.Count(item.Addr, ":") < 2 {
			inters[i].Family = "ipv4"
		} else {
			inters[i].Family = "ipv6"
		}
	}
	return inters
}

func CUPInfo() (CPU, error) {
	return CPU{}, nil
}
