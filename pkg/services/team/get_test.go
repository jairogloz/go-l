package team_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	teamService "github.com/jairogloz/go-l/pkg/services/team"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Get(t *testing.T) {

	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockTeamRepository)
		teamId        string
		assertionFunc func(subTest *testing.T, t *domain.Team, err error)
	}{
		"empty id": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "").Return(nil /*team*/, domain.ErrIncorrectID)
			},
			teamId: "",
			assertionFunc: func(subTest *testing.T, t *domain.Team, err error) {
				assert.Nil(subTest, t)
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeInvalidParams, appErr.Code)
					assert.Equal(subTest, "Incorrect id: unexpected error getting team", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},
		"team not found error": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "abc").Return(nil /*team*/, domain.ErrNotFound)
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, t *domain.Team, err error) {
				assert.Nil(subTest, t)
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeNotFound, appErr.Code)
					assert.Equal(subTest, "Not found: unexpected error getting team", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},

		"team timeout error": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "abc").Return(nil /*team*/, domain.ErrTimeout)
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, t *domain.Team, err error) {
				assert.Nil(subTest, t)
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeTimeout, appErr.Code)
					assert.Equal(subTest, "Timeout: unexpected error getting team", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},

		"team generic error": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "abc").Return(nil /*team*/, errors.New("generic error"))
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, t *domain.Team, err error) {
				assert.Nil(subTest, t)
				assert.NotNil(subTest, err)
			},
		},

		"success": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "abc").Return(&domain.Team{ID: "abc"}, nil)
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, t *domain.Team, err error) {
				assert.NotNil(subTest, t)
				assert.Nil(subTest, err)
				assert.Equal(subTest, "abc", t.ID)
			},
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {

			ctrl := gomock.NewController(subTest)
			defer ctrl.Finish()

			mockTeamRepo := mocks.NewMockTeamRepository(ctrl)

			s := &teamService.Service{
				Repo: mockTeamRepo,
			}

			test.setup(mockTeamRepo)

			team, err := s.Get(context.TODO(), test.teamId)

			test.assertionFunc(subTest, team, err)

		})
	}

}
