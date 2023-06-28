package domain

import "errors"

const (
	resourceProject = "project"
	resourceDataset = "dataset"
	resourceModel   = "model"
)

// ResourceType
type ResourceType interface {
	ResourceType() string
}

func NewResourceType(v string) (ResourceType, error) {
	b := v == resourceProject ||
		v == resourceModel ||
		v == resourceDataset
	if b {
		return resourceType(v), nil
	}

	return nil, errors.New("invalid resource type")
}

type resourceType string

func (r resourceType) ResourceType() string {
	return string(r)
}
