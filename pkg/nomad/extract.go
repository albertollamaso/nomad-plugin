package nomad

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
	"github.com/hashicorp/nomad/api"
)

// This method is mandatory the field extraction capability.
// If the Extract method is defined, the framework expects an Fields method
// to be specified too.
func (p *Plugin) Extract(req sdk.ExtractRequest, evt sdk.EventReader) error {
	var event api.Event
	encoder := json.NewDecoder(evt.Reader())
	if err := encoder.Decode(&event); err != nil {
		return err
	}

	switch req.Field() {
	case "nomad.index":
		req.SetValue(event.Index)
	case "nomad.alloc.name":
		alloc, err := event.Allocation()
		if err == nil {
			req.SetValue(string(alloc.Name))
		}
	case "nomad.alloc.namespace":
		alloc, err := event.Allocation()
		if err == nil {
			req.SetValue(string(alloc.Namespace))
		}
	case "nomad.alloc.jobID":
		alloc, err := event.Allocation()
		if err == nil {
			req.SetValue(string(alloc.JobID))
		}
	case "nomad.alloc.clientStatus":
		alloc, err := event.Allocation()
		if err == nil {
			req.SetValue(string(alloc.ClientStatus))
		}
	case "nomad.alloc.taskStates.type":
		var driver []string
		alloc, err := event.Allocation()
		if err == nil {
			for _, task := range alloc.TaskStates {
				for _, taskEvent := range task.Events {
					if taskEvent.Type == "Driver" {
						driver = append(driver, taskEvent.Type)
					}
				}
			}
		}
		req.SetValue(driver)
	case "nomad.alloc.res.cpu":
		alloc, err := event.Allocation()
		if err == nil {
			req.SetValue(uint64(*alloc.Resources.CPU))
		}
	case "nomad.alloc.res.cores":
		alloc, err := event.Allocation()
		if err == nil {
			req.SetValue(uint64(*alloc.Resources.Cores))
		}
	case "nomad.alloc.res.diskMB":
		alloc := event.Payload["Allocation"]
		valStr := fmt.Sprintf("%v", alloc.(map[string]interface{})["Resources"].(map[string]interface{})["DiskMB"]) // bug in nomad api. event.Allocation() returns <nil> for alloc.Resources.DiskMB
		value, _ := strconv.ParseUint(valStr, 10, 64)
		req.SetValue(value)
	case "nomad.alloc.res.iops":
		alloc, err := event.Allocation()
		if err == nil {
			req.SetValue(uint64(*alloc.Resources.IOPS))
		}
	case "nomad.alloc.res.memoryMB":
		alloc := event.Payload["Allocation"]
		valStr := fmt.Sprintf("%v", alloc.(map[string]interface{})["Resources"].(map[string]interface{})["MemoryMB"]) // bug in nomad api. event.Allocation() returns <nil> for alloc.Resources.MemoryMB
		value, _ := strconv.ParseUint(valStr, 10, 64)
		req.SetValue(value)
	case "nomad.alloc.res.memoryMaxMB":
		alloc := event.Payload["Allocation"]
		valStr := fmt.Sprintf("%v", alloc.(map[string]interface{})["Resources"].(map[string]interface{})["MemoryMaxMB"]) // bug in nomad api. event.Allocation() returns <nil> for alloc.Resources.MemoryMaxMB
		value, _ := strconv.ParseUint(valStr, 10, 64)
		req.SetValue(value)
	case "nomad.alloc.images":
		images, err := getAllocImages(&event)
		if err == nil {
			req.SetValue(images)
		}
	case "nomad.alloc.images.tags":
		tags, err := getAllocTags(&event)
		if err == nil {
			req.SetValue(tags)
		}
	case "nomad.alloc.images.repositories":
		repos, err := getAllocRepos(&event)
		if err == nil {
			req.SetValue(repos)
		}
	case "nomad.event.topic":
		req.SetValue(string(event.Topic))
	case "nomad.event.type":
		req.SetValue(event.Type)
	default:
		return fmt.Errorf("unsupported field: %s", req.Field())
	}

	return nil
}
