package league_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/jairogloz/go-l/pkg/services/league"
)

func TestService_Delete(t *testing.T) {

	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockLeagueRepository)
		leagueId      string
		assertionFunc func(subTest *testing.T, err error)
	}{
		"empty id": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				// No se espera llamada al método Delete cuando el ID está vacío
			},
			leagueId: "",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.EqualError(subTest, err, "id is required")
			},
		},
		"not found error": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Delete("abc").Return(domain.ErrNotFound)
			},
			leagueId: "abc",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeNotFound, appErr.Code)
					assert.Equal(subTest, "Not found: Error deleting league", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},
		"generic error": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Delete("abc").Return(errors.New("generic error"))
			},
			leagueId: "abc",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, err.Error(), "internal_server_error")
			},
		},
		"success": {
			setup: func(mockRepo *mocks.MockLeagueRepository) {
				mockRepo.EXPECT().Delete("abc").Return(nil)
			},
			leagueId: "abc",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NoError(subTest, err)
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

			err := s.Delete(test.leagueId)

			test.assertionFunc(subTest, err)

		})
	}
}
