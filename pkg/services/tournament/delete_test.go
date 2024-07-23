package tournament_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/jairogloz/go-l/pkg/services/tournament"
)

func TestService_Delete(t *testing.T) {
	ctx := context.TODO()

	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockTournamentRepository)
		ctx           context.Context
		id            string
		assertionFunc func(subTest *testing.T, err error)
	}{
		"success": {
			setup: func(mockRepo *mocks.MockTournamentRepository) {
				mockRepo.EXPECT().Delete(ctx, "667a09ac8c8a44e9c44ec248").
					Return(nil)
			},
			ctx: ctx,
			id:  "667a09ac8c8a44e9c44ec248",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Nil(subTest, err)
			},
		},
		"error incorrect id": {
			setup: func(mockRepo *mocks.MockTournamentRepository) {
				mockRepo.EXPECT().Delete(ctx, "incorrect-id").
					Return(domain.ErrIncorrectID)
			},
			ctx: ctx,
			id:  "incorrect-id",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)

				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeInvalidParams, appErr.Code)
					assert.Equal(subTest, "error deleting tournament: incorrect ID", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},
		"error not found": {
			setup: func(mockRepo *mocks.MockTournamentRepository) {
				mockRepo.EXPECT().Delete(ctx, "667a09ac8c8a44e9c44ec248").
					Return(domain.ErrNotFound)
			},
			ctx: ctx,
			id:  "667a09ac8c8a44e9c44ec248",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)

				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeNotFound, appErr.Code)
					assert.Equal(subTest, "error deleting tournament: not found", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},
		"generic error in repository": {
			setup: func(mockRepo *mocks.MockTournamentRepository) {
				mockRepo.EXPECT().Delete(ctx, "667a09ac8c8a44e9c44ec248").
					Return(errors.New("unexpected error deleting tournament"))
			},
			ctx: ctx,
			id:  "667a09ac8c8a44e9c44ec248",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error deleting tournament")
			},
		},
	}

	for name, tc := range testTable {
		t.Run(name, func(subTest *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockTournamentRepository(ctrl)
			tc.setup(mockRepo)

			s := tournament.Service{Repo: mockRepo}
			err := s.Delete(tc.ctx, tc.id)
			tc.assertionFunc(subTest, err)
		})
	}
}
