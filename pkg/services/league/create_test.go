package league_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/jairogloz/go-l/pkg/services/league"
)

func TestService_Create(t *testing.T) {
	ctx := context.TODO()

	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockLeagueRepository)
		ctx           context.Context
		league        domain.League
		assertionFunc func(subTest *testing.T, league *domain.League, err error)
	}{
		"success": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Insert(ctx, gomock.Any()).
					Do(func(ctx context.Context, league *domain.League) {
						league.ID = "123"
					}).
					Return(nil)
			},
			ctx: ctx,
			league: domain.League{
				Name:        "league test",
				Description: "a league for test",
			},
			assertionFunc: func(subTest *testing.T, league *domain.League, err error) {
				assert.Nil(subTest, err)
				assert.Equal(subTest, "123", league.ID)
			},
		},
		"error duplicated key": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Insert(ctx, gomock.Any()).
					Return(fmt.Errorf("%w: error inserting league", domain.ErrDuplicateKey))
			},
			ctx: ctx,
			league: domain.League{
				Name:        "league test",
				Description: "a league for test",
			},
			assertionFunc: func(subTest *testing.T, league *domain.League, err error) {
				assert.NotNil(subTest, err)

				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeDuplicateKey, appErr.Code)
					assert.Equal(subTest, "error creating league: duplicate key error", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},
		"generic error in repository": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Insert(ctx, gomock.Any()).
					Return(errors.New("unexpected error inserting league"))
			},
			ctx: ctx,
			league: domain.League{
				Name:        "league test",
				Description: "a league for test",
			},
			assertionFunc: func(subTest *testing.T, league *domain.League, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error inserting league")
			},
		},
	}

	for name, tc := range testTable {
		t.Run(name, func(subTest *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockLeagueRepository(ctrl)
			tc.setup(mockRepo)

			s := league.Service{Repo: mockRepo}
			err := s.Create(tc.ctx, &tc.league)
			tc.assertionFunc(subTest, &tc.league, err)
		})
	}
}
