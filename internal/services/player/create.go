package player

import (
	"fmt"
	"github.com/jairogloz/go-l/internal/domain"
	"log"
	"time"
)

func (s Service) Create(player domain.Player) (id interface{}, err error) {
	// Set creation time
	// Save to repo
	// Responder con el id del recurso creado
	player.CreationTime = time.Now().UTC()

	// ========= =repo
	insertedId, err := s.Repo.Insert(player)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating player: %w", err)
	}
	// ========= =repo

	return insertedId, nil
}
