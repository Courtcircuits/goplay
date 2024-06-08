package mod

import (
	"errors"
	"fmt"
	"strings"
)

func HelloWorld() {
	fmt.Println("Hello, World!")
}

type Media struct {
	Artist string
	Title  string
	Album  string
}

type Player interface {
	Status() string
	Play() error
	Pause() error
	PlayPause() error
	Stop() error
	Next() error
	Previous() error
	Metadata() (Media, error) //must be playing
}

type Playerctl struct {
	current_track Media
}

type TrackStatus string

const (
	Paused   TrackStatus = "Paused"
	Playing              = "Playing"
	Stopped              = "Stopped"
	NoPlayer             = "NoPlayer"
)

func (t TrackStatus) String() string {
	return string(t)
}

func (t TrackStatus) Show() string {
	switch t {
	case Paused:
		return ""
	case Playing:
		return ""
	case Stopped:
		return ""
	default:
		return ""
	}
}

var NotTrackStatus = errors.New("not a track status")

func TrackStatusFromString(s string) (TrackStatus, error) {
	s_sanitized := strings.TrimSuffix(s, "\n")

	switch s_sanitized {
	case "Paused":
		return Paused, nil
	case "Playing":
		return Playing, nil
	case "Stopped":
		return Stopped, nil
	case "No players found":
		return NoPlayer, nil
	default:
		return "", NotTrackStatus
	}
}

func NewPlayerctl() *Playerctl {
	return &Playerctl{}
}

func (p *Playerctl) Status() TrackStatus {
	cmd_return, err := Exec("playerctl", "status")
	if err != nil {

	}
	status, err := TrackStatusFromString(cmd_return)
	if err != nil {
		return NoPlayer
	}
	return status
}

func (p *Playerctl) Play() error {
	_, err := Exec("playerctl", "play")
	return err
}

func (p *Playerctl) Pause() error {
	err := ExecSilent("playerctl", "pause")
	return err
}

func (p *Playerctl) PlayPause() error {
	err := ExecSilent("playerctl", "play-pause")
	return err
}

func (p *Playerctl) Stop() error {
	err := ExecSilent("playerctl", "stop")
	return err
}

func (p *Playerctl) Next() error {
	err := ExecSilent("playerctl", "next")
	return err
}

func (p *Playerctl) Previous() error {
	err := ExecSilent("playerctl", "previous")
	return err
}

var NotPlaying = errors.New("player is not playing")
var NoAlbum = errors.New("no album metadata")
var NoArtist = errors.New("no artist metadata")
var NoTitle = errors.New("no title metadata")

func (p *Playerctl) Metadata() (Media, error) {
	if p.Status() != Playing {
		return Media{}, NotPlaying
	}
	artist, err := Exec("playerctl", "metadata", "xesam:artist")
	if err != nil {
		return Media{}, NoArtist
	}
	title, err := Exec("playerctl", "metadata", "xesam:title")
	if err != nil {
		return Media{}, NoTitle
	}
	album, err := Exec("playerctl", "metadata", "xesam:album")
	if err != nil {
		return Media{}, NoAlbum
	}
	return Media{artist, title, album}, nil
}

/*--------*/
// Media //
/*-------*/

func (m *Media) ToString() string {
	toShow := ""
	artist := strings.TrimSuffix(m.Artist, "\n")
	title := strings.TrimSuffix(m.Title, "\n")
	album := strings.TrimSuffix(m.Album, "\n")

	if artist != "" {
		toShow += artist
	}

	if title != "" {
		if artist != "" {
			toShow += " - "
		}
		toShow += title
	}

	if album != "" {
		if artist != "" || title != "" {
			toShow += " ("
		}
		toShow += album
		if artist != "" || title != "" {
			toShow += ")"
		}
	}
	return toShow
}
