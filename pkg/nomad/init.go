package nomad

import (
	"encoding/json"

	"github.com/alecthomas/jsonschema"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
)

// InitSchema is gets called by the SDK before initializing the plugin.
// This returns a schema representing the configuration expected by the
// plugin to be passed to the Init() method. Defining InitSchema() allows
// the framework to automatically validate the configuration, so that the
// plugin can assume that it to be always be well-formed when passed to Init().
// This is ignored if the return value is nil. The returned schema must follow
// the JSON Schema specific. See: https://json-schema.org/
// This method is optional.
func (p *Plugin) InitSchema() *sdk.SchemaInfo {
	// We leverage the jsonschema package to autogenerate the
	// JSON Schema definition using reflection from our config struct.
	schema, err := jsonschema.Reflect(&PluginConfig{}).MarshalJSON()
	if err == nil {
		return &sdk.SchemaInfo{
			Schema: string(schema),
		}
	}
	return nil
}

// Init initializes this plugin with a given config string.
// Since this plugin defines the InitSchema() method, we can assume
// that the configuration is pre-validated by the framework and
// always well-formed according to the provided schema.
// This method is mandatory.
func (p *Plugin) Init(config string) error {
	// This is where any state variables can be set and initialize
	p.Config.Reset()

	// Deserialize the config json. Ignoring the error
	// and not validating the config values is possible
	// due to the schema defined through InitSchema(),
	// for which the framework performas a pre-validation.
	return json.Unmarshal([]byte(config), &p.Config)
}

// Destroy is gets called by the SDK when the plugin gets deinitialized.
// This is useful to release any open resource used by the plugin.
// This method is optional.
func (p *Plugin) Destroy() {
	// here we can cleanup the plugin state when it gets destroyed
}
