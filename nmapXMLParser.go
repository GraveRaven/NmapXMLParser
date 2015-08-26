package nmapXMLParser

import (
	"encoding/xml"
	"net"
	"strings"
)

// Struct for reporting
type HostReport struct {
	IP    net.IP
	Ports []Port
}

//IPv4 returns all IPv4 IP addresses
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

//IPv6 returns all IPv6 IP addresses
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

//IPs returns all hosts IPs, both IPv4 and IPv6
func (n *NmapRun) IPs() (ips []net.IP){
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
		var ports []Port
		for _, port := range host.Ports.Port {
			if (port.Service.Name == "http" || port.Service.Name == "https") && port.Service.Tunnel == "ssl" {
				ports = append(ports, port)
			}
		}
		if len(ports) > 0 {
			if ips, err := host.IPs(); err == nil{
				for _, ip := range ips{
					r = append(r, HostReport{IP: ip, Ports: ports})
				}
			}
		}
	}

	return r
}

func (n *NmapRun) SSL() (r []HostReport){
	for _, host := range n.Host{
		var ports []Port
		for _, port := range host.Ports.Port{
			if port.Service.Tunnel == "ssl"{
				ports = append(ports, port)
			}
		}
		
		if len(ports) > 0{
			if ips, err := host.IPs(); err == nil{
				for _, ip := range ips{
					r = append(r, HostReport{IP: ip, Ports: ports})
				}
			}
		}
	}
	
	return r
}

//Port returns every IP where that port is open
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

func (n *NmapRun) Service(s string) (hosts []HostReport){
	
	for _, host := range n.Host{
		var ports []Port
		for _, port := range host.Ports.Port{
			if strings.ToLower(port.Service.Name) == strings.ToLower(s){
				ports = append(ports, port)
			}
		}
		
		if len(ports) > 0{
			if ips, err := host.IPs(); err == nil{
				for _, ip := range ips{
					hosts = append(hosts, HostReport{ip, ports})
				}
			}
		}
	}
	
	return hosts
}

//Product returns every host that matches a specific product
func (n *NmapRun) Product(s string) (hosts []HostReport){
	for _, host := range n.Host {
		var ports []Port
		
		for _, port := range host.Ports.Port{
			if strings.Contains(strings.ToLower(port.Service.Product), strings.ToLower(s)){
				ports = append(ports, port)
			}
		}
		
		if len(ports) > 0{
			if ips, err := host.IPs(); err == nil{
				for _, ip := range ips{
					hosts = append(hosts, HostReport{ip, ports})
				}
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

