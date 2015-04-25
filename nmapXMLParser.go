package nmapXMLParser

import (
	"encoding/xml"
)

//Missing: target,taskbegin,taskprocess,taskend,prescript,postscript,output
type NmapRun struct {
	Scanner          string `xml:"scanner,attr"`
	Args             string `xml:"args,attr"`
	Start            string `xml:"start,attr"`
	StartStr         string `xml:"startstr,attr"`
	Version          string `xml:"version,attr"`
	ProfileName      string `xml:"profile_name,attr"`
	XmlOutputVersion string `xml:"xmloutputversion,attr"`

	ScanInfo  ScanInfo  `xml:"scaninfo"`
	Verbose   Verbose   `xml:"verbose"`
	Debugging Debugging `xml:debugging`
	Host      []Host    `xml:"host"`
	RunStats  RunStats  `xml:"runstats"`
}

type ScanInfo struct {
	Type        string `xml:"type,attr"`
	ScanFlags   string `xml:"scanflags,attr"`
	Protocol    string `xml:"protocol,attr"`
	Numservices string `xml:"numservices,attr"`
	Services    string `xml:"services,attr"`
}

type Verbose struct {
	Level string `xml:"level,attr"`
}

type Debugging struct {
	Level string `xml:"level,attr"`
}

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

type Status struct {
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL string `xml:"reason_ttl,attr"`
}

type Address struct {
	Address     string `xml:"addr,attr"`
	AddressType string `xml:"addrtype,attr"`

	Vendor string `xml:"vendor"`
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

type Ports struct {
	ExtraPorts ExtraPorts `xml:extraports"`
	Port       []Port     `xml:"port"`
}

type ExtraPorts struct {
	State string `xml:"state,attr"`
	Count string `xml:"count,attr"`

	ExtraReasons ExtraReasons `xml:"extrareasons"`
}

type ExtraReasons struct {
	Reason string `xml:"reason,attr"`
	Count  string `xml:"count,attr"`
}

//Missing: owner
type Port struct {
	Protocol string `xml:"protocol,attr"`
	PortID   string `xml:"portid,attr"`

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

type RunStats struct {
	Finished Finished `xml:"finished"`
	Hosts    Hosts    `xml:"hosts"`
}

type Finished struct {
	Time     string `xml:"time,attr"`
	Timestr  string `xml:"timestr,attr"`
	Elapsed  string `xml:"elapsed,attr"`
	Summary  string `xml:"summary,attr"`
	Exit     string `xml:"exit,attr"`
	ErrorMsg string `xml:"errormsg,attr"`
}

type Hosts struct {
	Up    string `xml:"up,attr"`
	Down  string `xml:"down,attr"`
	Total string `xml:"total,attr"`
}

func ParseNmap(data []byte) (NmapRun, error) {
	var n NmapRun

	err := xml.Unmarshal(data, &n)
	return n, err
}
