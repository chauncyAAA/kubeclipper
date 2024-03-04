//go:build linux
// +build linux

package netutil

import (
	"errors"
	"net"

	"github.com/vishvananda/netlink"

	"github.com/kubeclipper/kubeclipper/pkg/utils/autodetection"
)

func GetDefaultIP(ipv4 bool, method string) (net.IP, error) {
	version := autodetection.IPv4
	if !ipv4 {
		version = autodetection.IPv4
	}
	ipNet, err := autodetection.AutoDetectCIDR(method, version)
	if err != nil {
		return nil, err
	}
	return ipNet.IP, nil
}

func GetDefaultGateway(ipv4 bool) (ip net.IP, err error) {
	family := netlink.FAMILY_V4
	if !ipv4 {
		family = netlink.FAMILY_V6
	}
	rl, err := netlink.RouteList(nil, family)
	if err != nil {
		return nil, err
	}
	for _, r := range rl {
		if r.Gw != nil && r.Dst == nil {
			return r.Gw, nil
		}
	}
	return nil, errors.New("no default gateway")
}
