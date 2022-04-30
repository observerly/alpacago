package alpacago

import (
	"fmt"
	"net"
)

const (
	ALPACA_DISCOVERY_VERSION = 1
	DEFAULT_PORT             = 11111
	DEFAULT_PORT_STR         = "11111"
)

type AlpacaDiscoveryMessage struct {
	Fixed    []byte // for discovery, this must be 'alpacadiscovery'
	Version  byte
	Reserved []byte
}

type AlpacaDiscoveryServer struct {
	Packet        *net.PacketConn
	UrlBase       string
	Protocol      string
	ClientId      uint32
	TransactionId uint32
	ErrorNumber   int
	ErrorMessage  string
}

func NewDiscoveryMessage(version int) *AlpacaDiscoveryMessage {
	a := AlpacaDiscoveryMessage{
		Fixed:    []byte("alpacadiscovery"),
		Version:  byte(fmt.Sprintf("%d", version)[0]),
		Reserved: make([]byte, 48),
	}

	return &a
}

func NewDiscoveryServer(clientId uint32, protocol string, domain string, ip string, port int32) *AlpacaDiscoveryServer {
	if port < 1 || port > 65535 {
		port = DEFAULT_PORT
	}

	var urlBase string = fmt.Sprintf("%s:%d", ip, port)

	if port == -1 && len(domain) > 0 {
		urlBase = domain
	}

	server := AlpacaDiscoveryServer{
		UrlBase:       urlBase,
		Protocol:      protocol,
		ClientId:      clientId,
		TransactionId: 0,
	}

	return &server
}
