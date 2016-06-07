package nats

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/ctypes"

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

	logger.Printf("%+v\n", metrics)

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

	cp.Add([]string{""}, config)

	return cp, nil
}
