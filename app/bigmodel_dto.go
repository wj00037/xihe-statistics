package app

import "project/xihe-statistics/domain"

type BigModelDTO struct {
	BigModel string   `json:"bigmodel"`
	Users    []string `json:"users"`
	Calls    int64    `json:"counts"`      // calls
	Counts   int      `json:"user_counts"` // xihe user counts
	UpdateAt string   `json:"update_at"`
}

type BigModelAllDTO struct {
	Users             []string `json:"users"`
	Counts            int      `json:"user_counts"`        // xihe user counts
	Calls             int64    `json:"counts"`             // calls
	DedupliacteCounts int      `json:"deduplicate_counts"` // xihe user deduplicate_counts
	UpdateAt          string   `json:"update_at"`
}

type BigModelCountIncreaseDTO struct {
	BigModel string `json:"bigmodel"`
	Counts   int64  `json:"counts"`
}

type UserWithBigModelAddCmd struct {
	UserName domain.Account
	BigModel domain.BigModel
	CreatAt  int64
}

type BigModelCountIncreaseCmd struct {
	BigModel  domain.BigModel `json:"bigmodel"`
	StartTime string          `json:"start_time"`
	EndTime   string          `json:"end_time"`
}
