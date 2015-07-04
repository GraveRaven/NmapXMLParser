package nmapXMLParser

import(
	"time"
	"encoding/xml"
	"fmt"
)

//Missing: target,taskbegin,taskprocess,taskend,prescript,postscript,output
type NmapRun struct {
	Scanner          string `xml:"scanner,attr"`
	Args             string `xml:"args,attr"`
	Start            time.Time `xml:"start,attr"`
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

func (n *NmapRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error{

	var aux struct{
		Scanner          string `xml:"scanner,attr"`
		Args             string `xml:"args,attr"`
		Start            int64 `xml:"start,attr"`
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
	
	if err := d.DecodeElement(&aux, &start); err != nil{
		return	fmt.Errorf("decode NmapRun: %v", err)
	}

	n.Scanner = aux.Scanner
	n.Args = aux.Args
	n.Start = time.Unix(aux.Start, 0)
	n.StartStr = aux.StartStr
	n.Version = aux.Version
	n.ProfileName = aux.ProfileName
	n.XmlOutputVersion = aux.XmlOutputVersion
	
	n.ScanInfo = aux.ScanInfo
	n.Verbose = aux.Verbose
	n.Debugging = aux.Debugging
	n.Host = aux.Host
	n.RunStats = aux.RunStats
	
	return nil
}


type ScanInfo struct {
	Type        string `xml:"type,attr"`
	ScanFlags   string `xml:"scanflags,attr"`
	Protocol    string `xml:"protocol,attr"`
	Numservices int `xml:"numservices,attr"`
	Services    string `xml:"services,attr"`
}

type Verbose struct {
	Level int `xml:"level,attr"`
}

type Debugging struct {
	Level int `xml:"level,attr"`
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
