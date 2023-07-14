package nuclei

import (
	"errors"
	"github.com/projectdiscovery/nuclei/v2/pkg/templates/types"
	"strings"
)

var protocolMappings = map[types.ProtocolType]string{
	types.InvalidProtocol:   "invalid",
	types.DNSProtocol:       "dns",
	types.FileProtocol:      "file",
	types.HTTPProtocol:      "http",
	types.HeadlessProtocol:  "headless",
	types.NetworkProtocol:   "tcp",
	types.WorkflowProtocol:  "workflow",
	types.SSLProtocol:       "ssl",
	types.WebsocketProtocol: "websocket",
	types.WHOISProtocol:     "whois",
}

func normalizeProtocolValue(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func stringToProtocol(protocolString string) (types.ProtocolType, error) {
	normalizedValue := normalizeProtocolValue(protocolString)
	for key, currentValue := range protocolMappings {
		if normalizedValue == currentValue {
			return key, nil
		}
	}
	return -1, errors.New("Invalid protocol: " + protocolString)
}
