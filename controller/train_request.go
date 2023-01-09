package controller

import "project/xihe-statistics/app"

type TrainIncreaseRequest struct {
	StartTime string `json:"start_time"`
	EndTime string	`json:"end_time"`
}

func (req TrainIncreaseRequest) toCmd() (
	cmd app.TrainIncreaseCmd,
	err error,
) {
	startTime := req.StartTime	// add time verify
	endTime := req.EndTime

	cmd = app.TrainIncreaseCmd{
		StartTime: startTime,
		EndTime: endTime,
	}

	return
}
