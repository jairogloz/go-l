package team

import (
	"context"
	"log"

	"github.com/jairogloz/go-l/internal/domain"
)

func (s Service) Delete(ctx context.Context, id string) (err error) {

	err = s.Repo.Delete(ctx, id)
	if err != nil {
		appErr := domain.ManageError(err, "")
		log.Println(appErr.Error())
		return appErr
	}

	return nil
}
