package main

import (
	"context"
	"time"
)

type Meet struct {
	ID           int
	OwnerID      int
	Name         string
	Description  string
	Country      string
	City         string
	StartDate    time.Time
	FinishDate   time.Time
	Longitude    float32
	Latitude     float32
	Participants []User
}

type User struct {
	ID    int
	Name  string
	Email string
}

type MeetStorer interface {
	CreateMeet(ctx context.Context, meet Meet) (*Meet, error)
	DeleteMeet(ctx context.Context, id int) error
	UpdateMeet(ctx context.Context, meet Meet) (*Meet, error)
	GetMeet(ctx context.Context, id int) (*Meet, error)
}
