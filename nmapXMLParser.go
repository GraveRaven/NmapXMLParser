package nmapXMLParser

import (
	"encoding/xml"
	"net"
)

// Struct for reporting
type HostReport struct {
	IP    net.IP
	Ports []int
}

func (n *NmapRun) IPv4() (ips []net.IP) {
	
	
	for _, host := range n.Host {
		for _, address := range host.Address {
			if address.AddressType == AddressIPv4 {
				ips = append(ips, address.Address)
			}
		}
	}
	return ips
}

func (n *NmapRun) IPv6() (ips []net.IP) {
	for _, host := range n.Host {
		for _, address := range host.Address {
			if address.AddressType == AddressIPv6{
				ips = append(ips, address.Address)
			}
		}
	}
	return ips
}

func (n *NmapRun) Hosts() (ips []net.IP){
	for _, host := range n.Host {
		for _, address := range host.Address{
			if address.AddressType == AddressIPv4 || address.AddressType == AddressIPv6{
				ips = append(ips, address.Address)
			}
		}
	}
	return ips
}


func (n *NmapRun) HTTPS() (r []HostReport) {
	for _, host := range n.Host {
		var ports []int
		for _, port := range host.Ports.Port {
			if (port.Service.Name == "http" || port.Service.Name == "https") && port.Service.Tunnel == "ssl" {
				ports = append(ports, port.PortID)
			}
		}
		if len(ports) > 0 {
			for _, ip := range host.Address {
				if ip.AddressType == AddressIPv4 || ip.AddressType == AddressIPv6 {
					r = append(r, HostReport{IP: ip.Address, Ports: ports})
				}
			}
		}
	}

	return r
}

func (n *NmapRun) SSL() (r []HostReport){
	for _, host := range n.Host{
		var ports []int
		for _, port := range host.Ports.Port{
			if port.Service.Tunnel == "ssl"{
				ports = append(ports, port.PortID)
			}
		}
		
		if len(ports) > 0{
			for _, ip := range host.Address{
				if ip.AddressType == AddressIPv4 || ip.AddressType == AddressIPv6{
					r = append(r, HostReport{IP: ip.Address, Ports: ports})
				}
			}
		}
	}
	
	return r
}

func (n *NmapRun) Port(p int) (hosts []net.IP) {
	for _, host := range n.Host {
		if host.Ports.State(p) == "open" {
			if ip, err := host.IPv4(); err == nil {
				hosts = append(hosts, ip)
			}
			if ip, err := host.IPv6(); err == nil{
				hosts = append(hosts, ip)
			}
		}
	}

	return hosts
}

func (n *NmapRun) Parse(data []byte) error {
	err := xml.Unmarshal(data, n)
	return err
}

func NewNmapRun(data []byte) (NmapRun, error) {
	var n NmapRun

	err := xml.Unmarshal(data, &n)
	return n, err
}

