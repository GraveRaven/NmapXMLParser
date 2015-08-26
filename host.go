package nmapXMLParser

import(
	"errors"
	"net"
	"encoding/xml"
	"fmt"
)

//Missing: hostscript
type Host struct {
	Starttime string `xml:"starttime,attr"`
	Endtime   string `xml:"endtime,attr"`
	Comment   string `xml:"comment,attr"`

	Status        Status        `xml:"status"`
	Address       []Address     `xml:"address"`
	Hostnames     Hostnames     `xml:"hostnames"`
	Smurf         Smurf         `xml:"smurf"`
	Ports         Ports         `xml:"ports"`
	OS            OS            `xml:"os"`
	Distance      Distance      `xml:"distance"`
	Uptime        Uptime        `xml:"uptime"`
	TcpSequence   TcpSequence   `xml:"tcpsequence"`
	IPIDSequence  IPIDSequence  `xml:"ipidsequence"`
	TCPtsSequence TCPtsSequence `xml:"tcptssequence"`
	Trace         Trace         `xml:"trace"`
	Times         Times         `xml:"times"`
}

type HostState string
const(
	HostStateUnknown HostState	= "unknown"
	HostStateUp					= "up"
	HostStateDown				= "down"
	HostStateSkipped			= "skipped"
)

type Status struct {
	State     HostState `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL string `xml:"reason_ttl,attr"`
}

type AddressType string
const(
	AddressUnknown AddressType	= ""
	AddressIPv4					= "ipv4"
	AddressIPv6 				= "ipv6"
	AddressMac					= "mac"
)

type Address struct {
	Address net.IP  `xml:"addr,attr"`
	Mac net.HardwareAddr `xml:"addr,attr"`
	AddressType AddressType `xml:"addrtype,attr"`

	Vendor string `xml:"vendor"`
}

func (a *Address) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error{
	var aux struct{
		Address     string `xml:"addr,attr"`
		AddressType string `xml:"addrtype,attr"`

		Vendor string `xml:"vendor"`
	}

	if err := d.DecodeElement(&aux, &start); err != nil{
		return	fmt.Errorf("decode Address: %v", err)
	}

	switch aux.AddressType{
		case "ipv4":
			a.AddressType = AddressIPv4
			a.Address = net.ParseIP(aux.Address)
		case "ipv6":
			a.AddressType = AddressIPv6
			a.Address = net.ParseIP(aux.Address)
		case "mac":
			a.AddressType = AddressMac
			a.Mac, _ = net.ParseMAC(aux.Address)
				
		default:
			a.AddressType = AddressUnknown
	}
	
	a.Vendor = aux.Vendor
	
	return nil
}

type Hostnames struct {
	Hostname []Hostname `xml:"hostname"`
}

type Hostname struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

type Smurf struct {
	Responses string `xml:"responses,attr"`
}

type OS struct {
	PortUsed      PortUsed      `xml:"portused"`
	OSMatch       OSMatch       `xml:"osmatch"`
	OSFingerprint OSFingerprint `xml:"osfingerprint"`
}

type Uptime struct {
	Seconds  string `xml:"seconds,attr"`
	Lastboot string `xml:"lastboot,attr"`
}

type Distance struct {
	Value string `xml:"value,attr"`
}

type TcpSequence struct {
	Index      string `xml:"index,attr"`
	Difficulty string `xml:"difficuly,attr"`
	Values     string `xml:"values,attr"`
}

type IPIDSequence struct {
	Class  string `xml:"class,attr"`
	Values string `xml:"values,attr"`
}

type TCPtsSequence struct {
	Class  string `xml:"class,attr"`
	Values string `xml:"values,attr"`
}

type Trace struct {
	Proto string `xml:"proto,attr"`
	Port  string `xml:"port,attr"`

	Hop Hop `xml:"hop"`
}

type Hop struct {
	TTL    string `xml:"ttl,attr"`
	RTTL   string `xml:"rttl,attr"`
	IPAddr string `xml:"ipaddr,attr"`
	Host   string `xml:"host,attr"`
}

type PortUsed struct {
	State  string `xml:"state,attr"`
	Proto  string `xml:"proto,attr"`
	PortID string `xml:"portid,attr"`
}

type OSMatch struct {
	Name     string `xml:"name,attr"`
	Accuracy string `xml:"accuracy,attr"`
	Line     string `xml:"line,attr"`

	OSClass OSClass `xml:"osclass"`
}

type OSClass struct {
	Vendor   string `xml:"vendor,attr"`
	OSGen    string `xml:"osgen,attr"`
	Type     string `xml:"type,attr"`
	Accuracy string `xml:"accuracy,attr"`
	OSFamily string `xml:"osfamily,attr"`

	CPE string `xml:"cpe"`
}

type OSFingerprint struct {
	Fingerprint string `xml:"fingerprint,attr"`
}

type Times struct {
	Srtt   string `xml:"srtt,attr"`
	Rttvar string `xml:"rttvar,attr"`
	To     string `xml:"to,attr"`
}

//IPs return all IPs associated with that host. Both v4 and v6
func (h *Host) IPs() (ips []net.IP, err error){
	for _, ip := range h.Address {
		if ip.AddressType == AddressIPv4 || ip.AddressType == AddressIPv6{
			ips = append(ips, ip.Address)
		}
	}	
	
	if len(ips) == 0{
		return nil, errors.New("No ip address")
	}
	
	return ips, nil
}

func (h *Host) IPv4() (net.IP, error) {
	for _, ip := range h.Address {
		if ip.AddressType == AddressIPv4 {
			return ip.Address, nil
		}
	}

	return nil, errors.New("No ipv4 address")
}

func (h *Host) IPv6() (net.IP, error) {
	for _, ip := range h.Address {
		if ip.AddressType == AddressIPv6 {
			return ip.Address, nil
		}
	}

	return nil, errors.New("No ipv6 address")
}

func (h *Host) PortsOpen() (ports []int) {
	for _, p := range h.Ports.Port {
		if p.State.State == "open" {
			ports = append(ports, p.PortID)
		}
	}
	return ports
}