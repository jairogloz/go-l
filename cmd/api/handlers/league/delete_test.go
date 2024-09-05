package league_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/handlers/league"
	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_DeleteLeague(t *testing.T) {
	testTable := map[string]struct {
		leagueID      string
		setup         func(leagueServiceMock *mocks.MockLeagueService)
		assertionFunc func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		"invalid id": {
			leagueID: "invalid-id",
			setup: func(leagueServiceMock *mocks.MockLeagueService) {
				leagueServiceMock.EXPECT().Delete("invalid-id").Return(domain.NewAppError(domain.ErrCodeInvalidParams, "invalid id"))
			},
			assertionFunc: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, 400, w.Code)
				var appErr domain.AppError
				err := json.Unmarshal(w.Body.Bytes(), &appErr)
				if err != nil {
					t.Fatal("error unmarshalling response", err.Error())
				}
				assert.Equal(t, domain.ErrCodeInvalidParams, appErr.Code)
			},
		},

		"success": {
			leagueID: "valid-id",
			setup: func(leagueServiceMock *mocks.MockLeagueService) {
				leagueServiceMock.EXPECT().Delete("valid-id").Return(nil)
			},
			assertionFunc: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, 200, w.Code)
			},
		},

		"league not found": {
			leagueID: "not-found-id",
			setup: func(leagueServiceMock *mocks.MockLeagueService) {
				leagueServiceMock.EXPECT().Delete("not-found-id").Return(domain.NewAppError(domain.ErrCodeNotFound, "league not found"))
			},
			assertionFunc: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, 404, w.Code)
				var appErr domain.AppError
				err := json.Unmarshal(w.Body.Bytes(), &appErr)
				if err != nil {
					t.Fatal("error unmarshalling response", err.Error())
				}
				assert.Equal(t, domain.ErrCodeNotFound, appErr.Code)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	leagueServiceMock := mocks.NewMockLeagueService(ctrl)

	h := league.Handler{
		LeagueService: leagueServiceMock,
	}

	gin.SetMode(gin.TestMode)

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {

			req, err := http.NewRequest(http.MethodDelete, "/leagues/"+test.leagueID, nil)
			if err != nil {
				t.Fatal("unexpected error creating request: ", err.Error())
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = gin.Params{{Key: "id", Value: test.leagueID}}

			test.setup(leagueServiceMock)

			h.DeleteLeague(c)

			test.assertionFunc(subTest, w)
		})
	}
}
