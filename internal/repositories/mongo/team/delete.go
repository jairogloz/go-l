package team

import (
	"context"
	"errors"
)

var (
	ErrTeamNotFound   = errors.New("player not found")
	ErrDeleteTeam     = errors.New("error deleting player")
	ErrIncorrectHexID = errors.New("incorrect mongo Hex format")
)

func (r Repository) Delete(ctx context.Context, id string) (err error) {
	//TODO implement me
	panic("implement me")
}
