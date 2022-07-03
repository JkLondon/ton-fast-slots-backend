package models

import "Casino/pkg/slot"

type StartGameResult struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause,omitempty"`
}

type RoundSlotResult struct {
	Success      bool                `json:"success"`
	WinAmount    float64             `json:"winAmount"`
	ScrollResult []slot.ScrollResult `json:"scrollResult"`
	Cause        string              `json:"cause"`
}

type EndGameResult struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
}

type RoundSlotParams struct {
	Amount float64 `json:"amount"`
}
