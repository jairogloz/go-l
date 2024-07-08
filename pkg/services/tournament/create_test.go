package tournament_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/jairogloz/go-l/pkg/services/tournament"
)

func TestService_Create(t *testing.T) {
	ctx := context.TODO()

	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockTournamentRepository)
		ctx           context.Context
		tournament    domain.Tournament
		assertionFunc func(subTest *testing.T, tm *domain.Tournament, err error)
	}{
		"success": {
			setup: func(mockRepo *mocks.MockTournamentRepository) {
				mockRepo.EXPECT().Insert(ctx, gomock.Any()).
					Do(func(ctx context.Context, tournament *domain.Tournament) {
						tournament.ID = "123"
					}).
					Return(nil)
			},
			ctx: ctx,
			tournament: domain.Tournament{
				Name:        "tournament test",
				Description: "a tournament for test",
				URL:         "https://tournament.test",
			},
			assertionFunc: func(subTest *testing.T, tm *domain.Tournament, err error) {
				assert.Nil(subTest, err)
				assert.Equal(subTest, "123", tm.ID)
			},
		},
		"error duplicated key": {
			setup: func(mockRepo *mocks.MockTournamentRepository) {
				mockRepo.EXPECT().Insert(ctx, gomock.Any()).
					Return(fmt.Errorf("%w: error inserting player", domain.ErrDuplicateKey))
			},
			ctx: ctx,
			tournament: domain.Tournament{
				Name:        "tournament test",
				Description: "a tournament for test",
				URL:         "https://tournament.test",
			},
			assertionFunc: func(subTest *testing.T, tm *domain.Tournament, err error) {
				assert.NotNil(subTest, err)

				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeDuplicateKey, appErr.Code)
					assert.Equal(subTest, "error creating tournament: duplicate key error", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},
		"generic error in repository": {
			setup: func(mockRepo *mocks.MockTournamentRepository) {
				mockRepo.EXPECT().Insert(ctx, gomock.Any()).
					Return(errors.New("unexpected error inserting tournament"))
			},
			ctx: ctx,
			tournament: domain.Tournament{
				Name:        "tournament test",
				Description: "a tournament for test",
				URL:         "https://tournament.test",
			},
			assertionFunc: func(subTest *testing.T, tm *domain.Tournament, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error inserting tournament")
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
			err := s.Create(tc.ctx, &tc.tournament)
			tc.assertionFunc(subTest, &tc.tournament, err)
		})
	}
}
