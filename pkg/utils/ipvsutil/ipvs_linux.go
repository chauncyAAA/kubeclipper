//go:build linux
// +build linux

package ipvsutil

import (
	"net"
	"syscall"

	"github.com/moby/ipvs"
)

const (
	rr = "rr"
)

func defaultIPVS(address string, port uint16) *ipvs.Service {
	vs := &ipvs.Service{
		AddressFamily: uint16(syscall.AF_INET),
		Address:       net.ParseIP(address),
		Protocol:      uint16(syscall.IPPROTO_TCP),
		Port:          port,
		SchedName:     rr,
		Flags:         0,
		Timeout:       0,
	}

	if ip4 := vs.Address.To4(); ip4 != nil {
		vs.AddressFamily = syscall.AF_INET
		vs.Netmask = 0xffffffff
	} else {
		vs.AddressFamily = syscall.AF_INET6
		vs.Netmask = 128
	}
	return vs
}

func defaultRS(address string, port uint16) *ipvs.Destination {
	rs := &ipvs.Destination{
		Address: net.ParseIP(address),
		Port:    port,
		Weight:  1,
	}
	return rs
}

func CreateIPVS(vs *VirtualServer, dryRun bool) error {
	if dryRun {
		return nil
	}
	handle, err := ipvs.New("")
	if err != nil {
		return err
	}
	dvs := defaultIPVS(vs.Address, vs.Port)
	err = handle.NewService(dvs)
	if err != nil {
		return err
	}
	for _, rs := range vs.RealServers {
		if err = handle.NewDestination(dvs, defaultRS(rs.Address, rs.Port)); err != nil {
			return err
		}
	}
	return nil
}

func DeleteIPVS(vs *VirtualServer, dryRun bool) error {
	if dryRun {
		return nil
	}
	handle, err := ipvs.New("")
	if err != nil {
		return err
	}
	return handle.DelService(defaultIPVS(vs.Address, vs.Port))
}

func ListIPVS(dryRun bool) ([]VirtualServer, error) {
	if dryRun {
		return nil, nil
	}

	handle, err := ipvs.New("")
	if err != nil {
		return nil, err
	}

	services, err := handle.GetServices()
	if err != nil {
		return nil, err
	}

	result := make([]VirtualServer, 0)
	for _, service := range services {
		destinations, err := handle.GetDestinations(service)
		if err != nil {
			return nil, err
		}
		des := make([]RealServer, 0)
		for _, de := range destinations {
			des = append(des, RealServer{
				Address: de.Address.String(),
				Port:    de.Port,
			})
		}
		result = append(result, VirtualServer{
			Address:     service.Address.String(),
			Port:        service.Port,
			RealServers: des,
		})
	}

	return result, nil
}

func GetIPVS(addr string, port uint16, dryRun bool) (*VirtualServer, error) {
	if dryRun {
		return nil, nil
	}

	handle, err := ipvs.New("")
	if err != nil {
		return nil, err
	}

	service, err := handle.GetService(defaultIPVS(addr, port))
	if err != nil {
		return nil, err
	}

	destinations, err := handle.GetDestinations(service)
	if err != nil {
		return nil, err
	}
	des := make([]RealServer, 0)
	for _, de := range destinations {
		des = append(des, RealServer{
			Address: de.Address.String(),
			Port:    de.Port,
		})
	}

	return &VirtualServer{
		Address:     service.Address.String(),
		Port:        service.Port,
		RealServers: des,
	}, nil
}

func Clear(dryRun bool) error {
	if dryRun {
		return nil
	}
	handle, err := ipvs.New("")
	if err != nil {
		return err
	}
	return handle.Flush()
}
