// This plugin is an example of plugin that supports both
// the event sourcing and the field extraction capabilities.
package nomad

import (
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins"
)

const (
	// note: 999 is for development only. Once released, plugins need to
	// get assigned an ID in the public Falcosecurity registry.
	// See: https://github.com/falcosecurity/plugins#plugin-registry
	PluginID          uint32 = 10
	PluginName               = "nomad"
	PluginDescription        = "Falcosecurity Nomad Plugin"
	PluginContact            = "github.com/albertollamaso/nomad-plugin"
	PluginVersion            = "0.2.0"
	PluginEventSource        = "nomad"
)

// Defining a type for the plugin.
// Composing the struct with plugins.BasePlugin is the recommended practice
// as it provides the boilerplate code that satisfies most of the interface
// requirements of the SDK.
//
// State variables to store in the plugin must be defined here.
type Plugin struct {
	plugins.BasePlugin
	Config PluginConfig
}

// Info returns a pointer to a plugin.Info struct,
// containing all the general information about this plugin.
// This method is mandatory.
func (m *Plugin) Info() *plugins.Info {
	return &plugins.Info{
		ID:          PluginID,
		Name:        PluginName,
		Description: PluginDescription,
		Contact:     PluginContact,
		Version:     PluginVersion,
		EventSource: PluginEventSource,
	}
}
