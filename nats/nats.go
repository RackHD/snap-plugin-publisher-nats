package nats

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/ctypes"
	"github.com/nats-io/nats"

	log "github.com/Sirupsen/logrus"
)

// Meta ...
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(name, version, pluginType, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

// Publisher ...
type Publisher struct{}

// NewPublisher ...
func NewPublisher() *Publisher {
	return &Publisher{}

}

const (
	name       = "nats"
	version    = 1
	pluginType = plugin.PublisherPluginType
)

// Publish ...
func (p *Publisher) Publish(contentType string, content []byte, config map[string]ctypes.ConfigValue) error {
	logger := log.New()
	var metrics []plugin.MetricType

	switch contentType {
	case plugin.SnapGOBContentType:
		decoder := gob.NewDecoder(bytes.NewBuffer(content))
		if err := decoder.Decode(&metrics); err != nil {
			logger.Printf("Error decoding: error=%v content=%v", err, content)
			return err
		}
	default:
		logger.Printf("Unknown content type '%v'.", contentType)
		return fmt.Errorf("Unknown content type '%s'.", contentType)
	}

	address := config["address"].(ctypes.ConfigValueStr).Value
	channel := config["channel"].(ctypes.ConfigValueStr).Value
	encoding := config["encoding"].(ctypes.ConfigValueStr).Value

	// Server Connection
	nc, err := nats.Connect(address)
	if err != nil {
		fmt.Printf("Could not connect to NATS server: %s\n\n", err)
		return err
	}
	defer nc.Close()

	// Connection encoding
	var c *nats.EncodedConn

	c, err = nats.NewEncodedConn(nc, encoding)
	if err != nil {
		fmt.Printf("Could not encode Nats connection with encoding type %s: %s \n\n", encoding, err)
		fmt.Printf("Using default encoding\n\n")

		c, err = nats.NewEncodedConn(nc, nats.DEFAULT_ENCODER)
		if err != nil {
			fmt.Printf("Could not encode Nats connection with default encoding type: %s \n\n", err)
			return err
		}
	}
	defer c.Close()

	//data := fmt.Sprintf("%v", metrics)
	for _, metric := range metrics {
		err = c.Publish(channel, metric)
		if err != nil {
			fmt.Printf("Could not publish to NATS server: %s\n\n", err)
			return err
		}
	}
	return nil
}

// GetConfigPolicy ...
func (p *Publisher) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	config := cpolicy.NewPolicyNode()

	address, err := cpolicy.NewStringRule("address", true)
	if err != nil {
		return nil, err
	}
	address.Description = "Nats URI"
	config.Add(address)

	channel, err := cpolicy.NewStringRule("channel", true)
	if err != nil {
		return nil, err
	}
	channel.Description = "Nats Channel"
	config.Add(channel)

	encoding, err := cpolicy.NewStringRule("encoding", true)
	if err != nil {
		return nil, err
	}
	encoding.Description = "Nats Encoding"
	config.Add(encoding)

	cp.Add([]string{""}, config)

	return cp, nil
}
