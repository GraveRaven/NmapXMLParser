package nmapXMLParser


type Ports struct {
	ExtraPorts ExtraPorts `xml:extraports"`
	Port       []Port     `xml:"port"`
}

//State gives the state of a port
func (ports *Ports) State(id int) string {
	for _, port := range ports.Port {
		if port.PortID == id {
			return port.State.State
		}
	}
	return "" //Not responding
}

type ExtraPorts struct {
	State string `xml:"state,attr"`
	Count int `xml:"count,attr"`

	ExtraReasons ExtraReasons `xml:"extrareasons"`
}

type ExtraReasons struct {
	Reason string `xml:"reason,attr"`
	Count  string `xml:"count,attr"`
}


type PortProtocol string
const(
	ProtoUnknown PortProtocol	= ""
	ProtoIP  					= "ip"
	ProtoTCP 					= "tcp"	
	ProtoUDP					= "udp"
	ProtoSCTP					= "sctp"
)

//Missing: owner
type Port struct {
	Protocol PortProtocol `xml:"protocol,attr"`
	PortID   int `xml:"portid,attr"`

	State   State   `xml:"state"`
	Service Service `xml:"service"`
	Script  Script  `xml:"script"`
}


type State struct {
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTR string `xml:"reason_ttl,attr"`
	ReasonIP  string `xml:"reason_ip,attr"`
}

type Service struct {
	Name       string `xml:"name,attr"`
	Product    string `xml:"product,attr"`
	Version    string `xml:"version,attr"`
	ExtraInfo  string `xml:"extrainfo,attr"`
	Tunnel     string `xml:"tunnel,attr"`
	Proto      string `xml:"proto,attr"`
	RpcNum     string `xml:"rpcnum,attr"`
	LowVer     string `xml:"lowver,attr"`
	HighVer    string `xml:"hihgver,attr"`
	Hostname   string `xml:"hostname,attr"`
	OSType     string `xml:"ostype,attr"`
	Method     string `xml:"method,attr"`
	Conf       string `xml:"conf,attr"`
	DeviceType string `xml:"devicetype,attr"`
	ServiceFp  string `xml:"servicefp,attr"`

	CPE string `xml:"cpe"`
}

type Script struct {
	ID     string `xml:"id,attr"`
	Output string `xml:"output,attr"`
}