package team_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/jairogloz/go-l/pkg/services/team"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Get(t *testing.T) {

	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockTeamRepository)
		teamId        string
		assertionFunc func(subTest *testing.T, t *domain.Team, p []domain.Player, err error)
	}{
		"empty id": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "").Return(nil /*team*/, domain.ErrIncorrectID)
			},
			teamId: "",
			assertionFunc: func(subTest *testing.T, t *domain.Team, p []domain.Player, err error) {
				assert.Nil(subTest, t)
				assert.Nil(subTest, p)
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
			assertionFunc: func(subTest *testing.T, t *domain.Team, p []domain.Player, err error) {
				assert.Nil(subTest, t)
				assert.Nil(subTest, p)
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
			assertionFunc: func(subTest *testing.T, t *domain.Team, p []domain.Player, err error) {
				assert.Nil(subTest, t)
				assert.Nil(subTest, p)
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
			assertionFunc: func(subTest *testing.T, t *domain.Team, p []domain.Player, err error) {
				assert.Nil(subTest, t)
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
			},
		},

		"players not found error": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "abc").Return(&domain.Team{ID: "abc"}, nil)
				mockRepo.EXPECT().GetPlayers(context.TODO(), "abc").Return(nil /*players*/, domain.ErrNotFound)
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, t *domain.Team, p []domain.Player, err error) {
				assert.Nil(subTest, t)
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeNotFound, appErr.Code)
					assert.Equal(subTest, "Not found: unexpected error getting players", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},

		"players timeout error": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "abc").Return(&domain.Team{ID: "abc"}, nil)
				mockRepo.EXPECT().GetPlayers(context.TODO(), "abc").Return(nil /*players*/, domain.ErrTimeout)
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, t *domain.Team, p []domain.Player, err error) {
				assert.Nil(subTest, t)
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeTimeout, appErr.Code)
					assert.Equal(subTest, "Timeout: unexpected error getting players", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},

		"players generic error": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "abc").Return(&domain.Team{ID: "abc"}, nil)
				mockRepo.EXPECT().GetPlayers(context.TODO(), "abc").Return(nil /*players*/, errors.New("generic error"))
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, t *domain.Team, p []domain.Player, err error) {
				assert.Nil(subTest, t)
				assert.Nil(subTest, p)
				assert.NotNil(subTest, err)
			},
		},
		"succes": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Get(context.TODO(), "abc").Return(&domain.Team{ID: "abc"}, nil)
				mockRepo.EXPECT().GetPlayers(context.TODO(), "abc").Return([]domain.Player{{ID: "1"}}, nil)
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, t *domain.Team, p []domain.Player, err error) {
				assert.NotNil(subTest, t)
				assert.NotNil(subTest, p)
				assert.Nil(subTest, err)
				assert.Equal(subTest, "abc", t.ID)
				assert.Equal(subTest, 1, len(p))
				assert.Equal(subTest, "1", p[0].ID)
			},
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {

			ctrl := gomock.NewController(subTest)
			defer ctrl.Finish()

			mockTeamRepo := mocks.NewMockTeamRepository(ctrl)

			s := &team.Service{
				Repo: mockTeamRepo,
			}

			test.setup(mockTeamRepo)

			team, players, err := s.Get(context.TODO(), test.teamId)

			test.assertionFunc(subTest, team, players, err)

		})
	}

}
