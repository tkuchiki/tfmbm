package resource

type BindingResource struct {
	Project string
	Role    string
	Members []string
}

type BindingResources []*BindingResource

func NewBindingResource() *BindingResource {
	return &BindingResource{}
}
