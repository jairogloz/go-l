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

func TestService_Create(t *testing.T) {
	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockPlayerRepository)
		player        *domain.Player
		assertionFunc func(subTest *testing.T, err error)
	}{
		"duplicate key error": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Insert(gomock.Any()).Return(domain.ErrDuplicateKey)
			},
			player: &domain.Player{
				FirstName: "Lionel",
				LastName:  "Messi",
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeDuplicateKey, appErr.Code)
					assert.Equal(subTest, "error creating player: duplicate key error", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},
		"generic error": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Insert(gomock.Any()).Return(errors.New("db connection failure"))
			},
			player: &domain.Player{
				FirstName: "Lionel",
				LastName:  "Messi",
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "error creating player: db connection failure")
			},
		},
		"success": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Insert(gomock.Any()).Return(nil)
			},
			player: &domain.Player{
				FirstName: "Lionel",
				LastName:  "Messi",
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Nil(subTest, err)
			},
		},
	}

	for testName, tc := range testTable {
		t.Run(testName, func(subTest *testing.T) {
			ctrl := gomock.NewController(subTest)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockPlayerRepository(ctrl)
			tc.setup(mockRepo)

			s := &player.Service{
				Repo: mockRepo,
			}
			err := s.Create(tc.player)

			tc.assertionFunc(subTest, err)
			assert.NotNil(subTest, tc.player.CreatedAt)
		})
	}
}
