package player

import (
	"fmt"
	"github.com/jairogloz/go-l/internal/domain"
	"log"
	"time"
)

func (s *Service) Create(player domain.Player) (id interface{}, err error) {
	now := time.Now().UTC()
	player.CreatedAt = &now

	// ========= =repo
	insertedId, err := s.Repo.Insert(player)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating player: %w", err)
	}
	// ========= =repo

	return insertedId, nil
}
