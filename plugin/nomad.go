package main

import (
	"github.com/albertollamaso/nomad-plugin/pkg/nomad"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins/extractor"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins/source"
)

// The plugin must be registered to the SDK in the init() function.
// Registering the plugin using both source.Register and extractor.Register
// declares to the SDK a plugin with both sourcing and extraction features
// enabled. The order in which the two Register functions are called is not
// relevant.
// This requires our plugin to implement the source.Plugin interface, so
// compilation will fail if the mandatory methods are not implemented.
func init() {
	// plugins are registered in the Plugin SDK by defining a factory function
	// to be used by the SDK whenever Falco initializes a new plugin
	plugins.SetFactory(func() plugins.Plugin {
		// creating an instance of our plugin so that the Plugin SDK
		// knows its type definition
		p := &nomad.Plugin{}

		// declares that our plugin supports the the event sourcing capability
		// (see: https://falco.org/docs/plugins/plugin-api-reference/#event-sourcing-capability-api)
		source.Register(p)

		// declares that our plugin supports the the field extraction capability
		// (see: https://falco.org/docs/plugins/plugin-api-reference/#field-extraction-capability-api)
		extractor.Register(p)
		return p
	})
}

// main is required just because the plugin is compiled in the main package,
// but it's not actually used in by the Plugin SDK
func main() {}
