package ports

import "github.com/jairogloz/go-l/pkg/domain"

type PlayerService interface {
	Create(player *domain.Player) (err error)
	Get(id string) (player *domain.Player, err error)
	Delete(id string) (err error)
}

type PlayerRepository interface {
	Insert(player *domain.Player) (err error)
	Get(id string) (player *domain.Player, err error)
	Delete(id string) (err error)
}
