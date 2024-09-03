package player

import "github.com/jairogloz/go-l/pkg/domain"

func (s *Service) GetAll() (player []*domain.Player, err error) {

	player, err = s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return player, nil
}
