package gitlab

import "project/xihe-statistics/domain/platform"

type cloneTotalResult struct {
	fetches `json:"fetches"`
}

type fetches struct {
	Total int64 `json:"total"`
}

func (c *cloneTotalResult) toCloneTotal() platform.CloneTotal {
	return platform.CloneTotal{
		Total: c.fetches.Total,
	}
}
