package ipvsutil

type VirtualServer struct {
	// loadblance host addr
	Address string
	// loadblance host port
	Port uint16

	RealServers []RealServer
}

type RealServer struct {
	// target host addr
	Address string
	// target host Port
	Port uint16
}
