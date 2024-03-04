package autodetection

import (
	"errors"
	"fmt"
	"net"
)

// DefaultInterfacesToExclude Default interfaces to exclude for any logic following the first-found
// autodetect IP method
var DefaultInterfacesToExclude = []string{
	"docker.*", "cbr.*", "dummy.*",
	"virbr.*", "lxcbr.*", "veth.*", "lo",
	"vxlan.calico", "cali.*", "tunl.*", "flannel.*", "kube-ipvs.*", "cni.*",
}

// FilteredEnumeration performs basic IP and IPNetwork discovery by enumerating
// all interfaces and filtering in/out based on the supplied filter regex.
//
// The incl and excl slice of regex strings may be nil.
func FilteredEnumeration(incl, excl []string, cidrs []net.IPNet, version int) (*Interface, *net.IPNet, error) {
	interfaces, err := GetInterfaces(incl, excl, version)
	if err != nil {
		return nil, nil, err
	}
	if len(interfaces) == 0 {
		return nil, nil, errors.New("no valid host interfaces found")
	}

	// Find the first interface with a valid matching IP address and network.
	// We initialise the IP with the first valid IP that we find just in
	// case we don't find an IP and network.
	for _, i := range interfaces {
		for _, c := range i.Cidrs {
			if c.IP.IsGlobalUnicast() && matchCIDRs(c.IP, cidrs) {
				return &i, &c, nil
			}
		}
	}

	return nil, nil, fmt.Errorf("no valid IPv%d addresses found on the host interfaces", version)
}

// matchCIDRs matches an IP address against a list of cidrs.
// If the list is empty, it always matches.
func matchCIDRs(ip net.IP, cidrs []net.IPNet) bool {
	if len(cidrs) == 0 {
		return true
	}
	for _, cidr := range cidrs {
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}
