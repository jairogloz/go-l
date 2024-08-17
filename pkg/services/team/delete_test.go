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

func TestService_Delete(t *testing.T) {

	testTable := map[string]struct {
		setup         func(mockRepo *mocks.MockTeamRepository)
		teamId        string
		assertionFunc func(subTest *testing.T, err error)
	}{
		"empty id": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Delete(context.TODO(), "").Return(domain.ErrIncorrectID)
			},
			teamId: "",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeInvalidParams, appErr.Code)
					assert.Equal(subTest, "Incorrect id: unexpected error deleting team", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},
		"delete not found error": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Delete(context.TODO(), "abc").Return(domain.ErrNotFound)
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
				var appErr domain.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, domain.ErrCodeNotFound, appErr.Code)
					assert.Equal(subTest, "Not found: unexpected error deleting team", appErr.Msg)
				} else {
					subTest.Errorf("expected AppError, got %v", err)
				}
			},
		},

		"team generic error": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Delete(context.TODO(), "abc").Return(errors.New("generic error"))
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)
			},
		},

		"succes": {
			setup: func(mockRepo *mocks.MockTeamRepository) {
				mockRepo.EXPECT().Delete(context.TODO(), "abc").Return(nil)
			},
			teamId: "abc",
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Nil(subTest, err)
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

			err := s.Delete(context.TODO(), test.teamId)

			test.assertionFunc(subTest, err)

		})
	}

}
