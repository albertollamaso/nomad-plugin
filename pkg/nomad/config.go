package nomad

// Defining a type for the plugin configuration.
// In this simple example, users can define the starting value the event
// counter. the `jsonschema` tags is used to automatically generate a
// JSON Schema definition, so that the framework can perform automatic
// validations.
type PluginConfig struct {
	Address   string `json:"address" jsonschema:"title=Nomad address,description=The address of the Nomad server.,default=http://localhost:4646"`
	Token     string `json:"token" jsonschema:"title=Nomad token,description=The token to use to connect to the Nomad server.,default="`
	Namespace string `json:"namespace" jsonschema:"title=Nomad namespace,description=The namespace to use to connect to the Nomad server.,default=*"`
}

// Resets sets the configuration to its default values
func (p *PluginConfig) Reset() {
	p.Address = "http://localhost:4646"
	p.Token = ""
	p.Namespace = "*"
}
