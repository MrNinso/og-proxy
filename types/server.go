package types

type ServerHTTPS struct {
	Type        string     `json:"type"`
	Hosts       []HostHTTP `json:"hosts"`
	ServerName  string     `json:"name"`
	ServerPort  uint       `json:"port"`
	Certificate string     `json:"ssl-crt"`
	Key         string     `json:"ssl-key"`
}

type ServerTCP struct {
	Type  string       `json:"type"`
	Hosts []HostTCPUDP `json:"hosts"`
	Port  uint         `json:"port"`
}

type ServerUDP struct {
	Type  string       `json:"type"`
	Hosts []HostTCPUDP `json:"hosts"`
	Port  uint         `json:"port"`
}
