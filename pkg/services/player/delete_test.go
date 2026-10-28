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

func TestService_Delete(t *testing.T) {
	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockPlayerRepository)
		playerId      string
		assertionFunc func(subTest *testing.T, err error)
	}{
		"not found error": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Delete("abc").Return(domain.ErrNotFound)
			},
			playerId: "abc",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeNotFound, appErr.Code)
					assert.Equal(subTest, "Not found: Error deleting player", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},
		"generic error": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Delete("abc").Return(errors.New("db error"))
			},
			playerId: "abc",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "internal_server_error")
			},
		},
		"success": {
			setup: func(mockRepo *mocks.MockPlayerRepository) {
				mockRepo.EXPECT().Delete("abc").Return(nil)
			},
			playerId: "abc",
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
			err := s.Delete(tc.playerId)

			tc.assertionFunc(subTest, err)
		})
	}
}
