package main

import (
	"context"
	"time"
)

type Meet struct {
	ID           int       `json:"id"`
	OwnerID      int       `json:"owner_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Country      string    `json:"country"`
	City         string    `json:"city"`
	StartDate    time.Time `json:"start_date"`
	FinishDate   time.Time `json:"finish_date"`
	Longitude    float32   `json:"longitude"`
	Latitude     float32   `json:"latitude"`
	Participants []User    `json:"participants"`
}

type User struct {
	ID    int
	Name  string
	Email string
}

type MeetStorer interface {
	CreateMeet(ctx context.Context, meet *Meet) (*Meet, error)
	DeleteMeet(ctx context.Context, id int) error
	UpdateMeet(ctx context.Context, meet Meet) (*Meet, error)
	GetMeet(ctx context.Context, id int) (*Meet, error)
	GetMeets(ctx context.Context) ([]*Meet, error)
}
