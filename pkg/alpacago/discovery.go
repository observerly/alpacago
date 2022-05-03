package alpacago

import (
	"fmt"
	"log"
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
	Address       *net.UDPAddr
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

/*
	OpenSocket()

	@returns A bound open a socket for discovery on the udp4 network protocol, i.e., announce on the local network address.
	@see https://ascom-standards.org/Developer/ASCOM%20Alpaca%20API%20Reference.pdf
*/
func (s *AlpacaDiscoveryServer) OpenSocket() {
	packet, err := net.ListenPacket("udp4", ":32227")

	if err != nil {
		log.Fatal(fmt.Errorf("unable to open ASCOM Alpaca API discovery listen socket: %s", err.Error()))
	}

	defer packet.Close()

	s.Packet = &packet
}

/*
	ResolveUDPAddress()

	@returns A resolved UDP address broadcast by the server.
	@see https://ascom-standards.org/Developer/ASCOM%20Alpaca%20API%20Reference.pdf
*/
func (s *AlpacaDiscoveryServer) ResolveUDPAddress() {
	address, err := net.ResolveUDPAddr("udp4", ":32227")

	if err != nil {
		log.Fatal(fmt.Errorf("unable to resolve ASCOM Alpaca API discovery UDP broadcast address: %s", err.Error()))
	}

	s.Address = address
}
