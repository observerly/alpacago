package alpacago

import "fmt"

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

func NewDiscoveryMessage(version int) *AlpacaDiscoveryMessage {
	a := AlpacaDiscoveryMessage{
		Fixed:    []byte("alpacadiscovery"),
		Version:  byte(fmt.Sprintf("%d", version)[0]),
		Reserved: make([]byte, 48),
	}

	return &a
}
