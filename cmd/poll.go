package main

import (
	"time"

	"github.com/Courtcircuits/goplay/mod"
)

func Poll(interval time.Duration) {
	p := mod.NewPlayerctl()
	currentTrack, err := GetCurrentTrack(p)
	currentStatus := p.Status()
	if err == nil {
		Show(currentStatus, currentTrack)
	} else {
		ShowError(currentStatus)
	}

	for {
		track, err := GetCurrentTrack(p)
		status := p.Status()
		if err == nil {
			if track != currentTrack || status != currentStatus {
				Show(status, track)
				currentTrack = track
				currentStatus = status
			}
		} else {
			if status != currentStatus {
				ShowError(status)
				currentStatus = status
			}
		}
		time.Sleep(1 * time.Second)
	}
}
