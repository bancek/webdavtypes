package webdavtypes

import "encoding/xml"

type LockScope struct {
	Value xml.Name `xml:",any"`
}

type LockType struct {
	Value xml.Name `xml:",any"`
}

type LockEntry struct {
	LockScope LockScope `xml:"lockscope"`
	LockType  LockType  `xml:"locktype"`
}

type Supportedlock struct {
	Lockentry LockEntry `xml:"lockentry"`
}

type Owner struct {
	Href string `xml:"href"`
}

type LockToken struct {
	Href string `xml:"href"`
}

type LockRoot struct {
	Href string `xml:"href"`
}

type ResourceType struct {
	Value xml.Name `xml:",any"`
}

type ActiveLock struct {
	LockType  LockType  `xml:"locktype"`
	LockScope LockScope `xml:"lockscope"`
	Depth     string    `xml:"depth"`
	Owner     Owner     `xml:"owner"`
	Timeout   string    `xml:"timeout"`
	LockToken LockToken `xml:"locktoken"`
	LockRoot  LockRoot  `xml:"lockroot"`
}

type LockDiscovery struct {
	ActiveLock ActiveLock `xml:"activelock"`
}

type Prop struct {
	Displayname           string        `xml:"displayname"`
	Getlastmodified       string        `xml:"getlastmodified"`
	Getcontentlength      int64         `xml:"getcontentlength"`
	Getcontenttype        string        `xml:"getcontenttype"`
	Getetag               string        `xml:"getetag"`
	Supportedlock         Supportedlock `xml:"supportedlock"`
	Resourcetype          ResourceType  `xml:"resourcetype"`
	LockDiscovery         LockDiscovery `xml:"lockdiscovery"`
	Win32LastModifiedTime string        `xml:"Win32LastModifiedTime"`
}

type Propstat struct {
	Prop   Prop   `xml:"prop"`
	Status string `xml:"status"`
}

type Response struct {
	Href     string   `xml:"href"`
	Propstat Propstat `xml:"propstat"`
}

type Multistatus struct {
	Response []Response `xml:"response"`
}
