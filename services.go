package main

import (
	"context"
	"fmt"
)

type MeetService struct {
	meetStorer MeetStorer
}

func NewMeetService(meetStorer MeetStorer) *MeetService {
	return &MeetService{meetStorer: meetStorer}
}

func (m *MeetService) CreateMeet(ctx context.Context, meet *Meet) error {
	_, err := m.meetStorer.CreateMeet(ctx, meet)

	if err != nil {
		return fmt.Errorf("Error creating new meet: %s", err)
	}

	return nil
}

func (m *MeetService) DeleteMeet(ctx context.Context, id int) error {
	err := m.meetStorer.DeleteMeet(ctx, id)

	if err != nil {
		return fmt.Errorf("Error deleting the meet: %s", err)
	}

	return nil
}

func (m *MeetService) GetMeet(ctx context.Context, id int) (*Meet, error) {
	meet, err := m.meetStorer.GetMeet(ctx, id)

	if err != nil {
		return &Meet{}, fmt.Errorf("Error retrieving the meet: %s", err)
	}

	return meet, nil
}

func (m *MeetService) UpdateMeet(ctx context.Context, meet Meet) (*Meet, error) {
	newMeet, err := m.meetStorer.UpdateMeet(ctx, meet)

	if err != nil {
		return &Meet{}, fmt.Errorf("Error updating the meet: %s", err)
	}

	return newMeet, err
}

func (m *MeetService) GetMeets(ctx context.Context) ([]*Meet, error) {
	meets, err := m.meetStorer.GetMeets(ctx)

	if err != nil {
		return []*Meet{}, fmt.Errorf("Error getting meets: %s", err)
	}

	return meets, nil
}
