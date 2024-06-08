package main

import (
	"fmt"
	"time"

	"github.com/Courtcircuits/goplay/mod"
)

func main() {
	Poll(1 * time.Second)
}

func ShowError(status mod.TrackStatus) {
	fmt.Println(status.Show())
}

func Show(status mod.TrackStatus, track mod.Media) {
	fmt.Printf("%s %s\n", status.Show(), track.ToString())
}

func GetCurrentTrack(p *mod.Playerctl) (mod.Media, error) {
	return p.Metadata()
}
