package ports

import "github.com/jairogloz/go-l/pkg/domain"

type PlayerService interface {
	Create(player *domain.Player) (err error)
	Get(id string) (player *domain.Player, err error)
	Delete(id string) (err error)
	GetAll() (players []*domain.Player, err error)
}

type PlayerRepository interface {
	Insert(player *domain.Player) (err error)
	Get(id string) (player *domain.Player, err error)
	GetPlayersByTeamID(id string) (players []*domain.Player, err error)
	Delete(id string) (err error)
	GetAll() (players []*domain.Player, err error)
}
