package types

type HostHTTP struct {
	Hostname string `json:"hostname"`
	Port     uint   `json:"port"`
	Path     string `json:"path"`
	Weight   uint8  `json:"weight"`
}

type HostTCPUDP struct {
	Hostname string `json:"hostname"`
	Port     uint   `json:"port"`
	Weight   uint8  `json:"weight"`
}