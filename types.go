package main

type Track struct {
	ID              int
	Name            string
	MaxProgress     int
	CurrentProgress int
}

type TrackData struct {
	Tracks []Track
}
