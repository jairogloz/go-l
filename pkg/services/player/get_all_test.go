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
