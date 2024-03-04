//go:build !linux
// +build !linux

package ipvsutil

func defaultIPVS(address string, port uint16) *VirtualServer {
	return nil
}

func CreateIPVS(vs *VirtualServer, dryRun bool) error {
	return nil
}

func DeleteIPVS(vs *VirtualServer, dryRun bool) error {
	return nil
}

func ListIPVS(dryRun bool) ([]VirtualServer, error) {
	return nil, nil
}

func GetIPVS(addr string, port uint16, dryRun bool) (*VirtualServer, error) {
	return nil, nil
}

func Clear(dryRun bool) error {
	return nil
}
