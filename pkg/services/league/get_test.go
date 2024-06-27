package league_test

import (
	"errors"
	"testing"

	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/jairogloz/go-l/pkg/services/league"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Get(t *testing.T) {

	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockLeagueRepository)
		leagueId      string
		assertionFunc func(subTest *testing.T, p *domain.League, err error)
	}{
		"empty id": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {

			},
			leagueId: "",
			assertionFunc: func(subTest *testing.T, p *domain.League, err error) {
				assert.Nil(subTest, p)
				assert.EqualError(subTest, err, "id is required")
			},
		},
		"not found error": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Get("america cup").Return(nil /*league*/, domain.ErrNotFound)
			},
			leagueId: "america cup",
			assertionFunc: func(subTest *testing.T, p *domain.League, err error) {
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeNotFound, appErr.Code)
					assert.Equal(subTest, "league with id 'america cup' not found", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},

		"timeout error": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Get("america cup").Return(nil /*player*/, domain.ErrTimeout)
			},
			leagueId: "america cup",
			assertionFunc: func(subTest *testing.T, p *domain.League, err error) {
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeTimeout, appErr.Code)
					assert.Equal(subTest, "timeout error, try again later", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},

		"generic error": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Get("america cup").Return(nil /*player*/, errors.New("generic error"))
			},
			leagueId: "america cup",
			assertionFunc: func(subTest *testing.T, p *domain.League, err error) {
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error getting league:")
			},
		},

		"success": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Get("america cup").Return(&domain.League{ID: "america cup"}, nil)
			},
			leagueId: "america cup",
			assertionFunc: func(subTest *testing.T, p *domain.League, err error) {
				assert.NotNil(subTest, p)
				assert.Equal(subTest, "america cup", p.ID)
				assert.NoError(t, err)
			},
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {

			ctrl := gomock.NewController(subTest)
			defer ctrl.Finish()

			mockLeagueRepo := mocks.NewMockLeagueRepository(ctrl)

			s := &league.Service{
				Repo: mockLeagueRepo,
			}

			test.setup(mockLeagueRepo)

			p, err := s.Get(test.leagueId)

			test.assertionFunc(subTest, p, err)

		})
	}

}
