package nomad

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins/source"
	"github.com/hashicorp/nomad/api"
)

// Open opens the plugin source and starts a new capture session (e.g. stream
// of events), creating a new plugin instance. The state of each instance can
// be initialized here. This method is mandatory for the event sourcing capability.
func (m *Plugin) Open(params string) (source.Instance, error) {

	ctx, cancel := context.WithCancel(context.Background())

	// Define Nomad client
	client, err := api.NewClient(&api.Config{
		Address:   m.Config.Address,
		SecretID:  m.Config.Token,
		Namespace: m.Config.Namespace,
	})
	if err != nil {
		return nil, err
	}

	topics := map[api.Topic][]string{
		api.TopicAll: {"*"},
	}

	streamCh, err := client.EventStream().Stream(ctx, topics, 0, &api.QueryOptions{
		WaitIndex: 0,
		WaitTime:  30 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	eventCh := make(chan source.PushEvent)
	go func() {
		defer close(eventCh)

		for event := range streamCh {
			m.ParseEventsAndPush(event, eventCh)
		}
	}()
	return source.NewPushInstance(eventCh, source.WithInstanceClose(cancel))
}

func (m *Plugin) ParseEventsAndPush(events *api.Events, output chan<- source.PushEvent) {

	if events.Err != nil {
		output <- source.PushEvent{
			Err: events.Err,
		}
		return
	}

	var buffer bytes.Buffer
	for _, event := range events.Events {
		buffer.Reset()
		json.NewEncoder(&buffer).Encode(event)
		output <- source.PushEvent{
			Err:       nil,
			Data:      buffer.Bytes(),
			Timestamp: time.Now(),
		}
	}
}
