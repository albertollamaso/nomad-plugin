package nomad

import (
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
)

// Fields return the list of extractor fields exported by this plugin.
// This method is mandatory the field extraction capability.
// If the Fields method is defined, the framework expects an Extract method
// to be specified too.
func (p *Plugin) Fields() []sdk.FieldEntry {
	return []sdk.FieldEntry{
		{Type: "uint64", Name: "nomad.index", Display: "Event index", Desc: "the index of the nomad event."},
		{Type: "string", Name: "nomad.alloc.name", Display: "Allocation name", Desc: "the name of the nomad allocation."},
		{Type: "string", Name: "nomad.alloc.namespace", Display: "Allocation namespace", Desc: "the namespace of the allocation."},
		{Type: "string", Name: "nomad.alloc.jobID", Display: "Allocation Job ID", Desc: "the job ID of the allocation."},
		{Type: "string", Name: "nomad.alloc.clientStatus", Display: "Allocation client status", Desc: "the client status of the allocation."},
		{Type: "string", Name: "nomad.alloc.images", Display: "Allocation container images", Desc: "the list of container images on allocations.", IsList: true},
		{Type: "string", Name: "nomad.alloc.images.tags", Display: "Allocation container tags", Desc: "the tags of each container image on allocations.", IsList: true},
		{Type: "string", Name: "nomad.alloc.images.repositories", Display: "Allocation container repositories", Desc: "the container repositories used on allocations container images.", IsList: true},
		{Type: "string", Name: "nomad.alloc.taskStates.type", Display: "Allocation Task State", Desc: "the state of the task on the allocations.", IsList: true},
		{Type: "uint64", Name: "nomad.alloc.res.cpu", Display: "Allocation CPU Resources", Desc: "the CPU required to run this allocation in MHz."},
		{Type: "uint64", Name: "nomad.alloc.res.cores", Display: "Allocation CPU Cores Resources", Desc: "the number of CPU cores to reserve for the allocation."},
		{Type: "uint64", Name: "nomad.alloc.res.diskMB", Display: "Allocation Disk in MB Resources", Desc: "the amount of disk required for the allocation."},
		{Type: "uint64", Name: "nomad.alloc.res.iops", Display: "Allocation IOPS Resources", Desc: "the number of iops required for the allocation."},
		{Type: "uint64", Name: "nomad.alloc.res.memoryMB", Display: "Allocation Memory in MB Resources", Desc: "the memory required in MB for the allocation."},
		{Type: "uint64", Name: "nomad.alloc.res.memoryMaxMB", Display: "Allocation Max Memory in MB Resources", Desc: "the maximum memory the allocation may use."},
		{Type: "string", Name: "nomad.event.topic", Display: "Nomad Event Topic", Desc: "the topic of the nomad event."},
		{Type: "string", Name: "nomad.event.type", Display: "Nomad Event type", Desc: "the type of the nomad event."},
	}
}
