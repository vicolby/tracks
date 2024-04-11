package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
)

type MeetRepository struct {
	meets map[string]Meet
	mu    sync.Mutex
}

func NewMeetRepository() *MeetRepository {
	return &MeetRepository{meets: make(map[string]Meet)}
}

func (m *MeetRepository) CreateMeet(ctx context.Context, meet *Meet) (*Meet, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := strconv.Itoa(meet.ID)

	m.meets[key] = *meet

	return meet, nil
}

func (m *MeetRepository) DeleteMeet(ctx context.Context, id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := strconv.Itoa(id)

	delete(m.meets, key)

	return nil
}

func (m *MeetRepository) UpdateMeet(ctx context.Context, meet Meet) (*Meet, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := strconv.Itoa(meet.ID)

	m.meets[key] = meet

	return &meet, nil
}

func (m *MeetRepository) GetMeet(ctx context.Context, id int) (*Meet, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := strconv.Itoa(id)

	meet, exists := m.meets[key]
	if !exists {
		return nil, fmt.Errorf("meet with ID %d not found", id)
	}

	return &meet, nil
}

func (m *MeetRepository) GetMeets(ctx context.Context) ([]*Meet, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	meets := []*Meet{}

	for _, value := range m.meets {
		meet := value
		meets = append(meets, &meet)
	}

	return meets, nil
}
