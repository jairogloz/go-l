package player_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/jairogloz/go-l/pkg/services/player"
)

func TestService_Get(t *testing.T) {

	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockPlayerRepository)
		playerId      string
		assertionFunc func(subTest *testing.T, p *domain.Player, err error)
	}{
		"empty id": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {

			},
			playerId: "",
			assertionFunc: func(subTest *testing.T, p *domain.Player, err error) {
				assert.Nil(subTest, p)
				assert.EqualError(subTest, err, "id is required")
			},
		},
		"not found error": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Get("abc").Return(nil /*player*/, domain.ErrNotFound)
			},
			playerId: "abc",
			assertionFunc: func(subTest *testing.T, p *domain.Player, err error) {
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeNotFound, appErr.Code)
					assert.Equal(subTest, "player with id 'abc' not found", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},

		"timeout error": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Get("abc").Return(nil /*player*/, domain.ErrTimeout)
			},
			playerId: "abc",
			assertionFunc: func(subTest *testing.T, p *domain.Player, err error) {
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
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Get("abc").Return(nil /*player*/, errors.New("generic error"))
			},
			playerId: "abc",
			assertionFunc: func(subTest *testing.T, p *domain.Player, err error) {
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "unexpected error getting player:")
			},
		},

		"success": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Get("abc").Return(&domain.Player{ID: "abc"}, nil)
			},
			playerId: "abc",
			assertionFunc: func(subTest *testing.T, p *domain.Player, err error) {
				assert.NotNil(subTest, p)
				assert.Equal(subTest, "abc", p.ID)
				assert.NoError(t, err)
			},
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {

			ctrl := gomock.NewController(subTest)
			defer ctrl.Finish()

			mockPlayerRepo := mocks.NewMockPlayerRepository(ctrl)

			s := &player.Service{
				Repo: mockPlayerRepo,
			}

			test.setup(mockPlayerRepo)

			p, err := s.Get(test.playerId)

			test.assertionFunc(subTest, p, err)

		})
	}

}

func TestService_GetAll(t *testing.T) {
	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockPlayerRepository)
		assertionFunc func(subTest *testing.T, players []*domain.Player, err error)
	}{
		"success - players found": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().GetAll().Return([]*domain.Player{
					{ID: "1", FirstName: "John", LastName: "Doe"},
					{ID: "2", FirstName: "Jane", LastName: "Smith"},
				}, nil)
			},
			assertionFunc: func(subTest *testing.T, players []*domain.Player, err error) {
				assert.NoError(subTest, err)
				assert.NotNil(subTest, players)
				assert.Len(subTest, players, 2)
				assert.Equal(subTest, "John", players[0].FirstName)
				assert.Equal(subTest, "Jane", players[1].FirstName)
			},
		},
		"timeout error": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().GetAll().Return(nil, domain.ErrTimeout)
			},
			assertionFunc: func(subTest *testing.T, players []*domain.Player, err error) {
				assert.Nil(subTest, players)
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
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().GetAll().Return(nil, errors.New("generic error"))
			},
			assertionFunc: func(subTest *testing.T, players []*domain.Player, err error) {
				assert.Nil(subTest, players)
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "generic error")
			},
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {

			ctrl := gomock.NewController(subTest)
			defer ctrl.Finish()

			mockPlayerRepo := mocks.NewMockPlayerRepository(ctrl)

			s := &player.Service{
				Repo: mockPlayerRepo,
			}

			test.setup(mockPlayerRepo)

			players, err := s.GetAll()

			test.assertionFunc(subTest, players, err)
		})
	}
}
