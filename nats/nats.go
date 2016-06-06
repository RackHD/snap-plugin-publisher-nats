package nats

import (
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/ctypes"
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
	return nil
}

// GetConfigPolicy ...
func (p *Publisher) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	return nil, nil
}
