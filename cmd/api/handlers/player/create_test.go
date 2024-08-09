package player_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/handlers/player"
	"github.com/jairogloz/go-l/mocks"
	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreatePlayer(t *testing.T) {

	testTable := map[string]struct {
		reqBodyString string
		setup         func(playerServiceMock *mocks.MockPlayerService)
		assertionFunc func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		"invalid json": {
			reqBodyString: `{"invalid":"json"`,
			setup: func(playerServiceMock *mocks.MockPlayerService) {

			},
			assertionFunc: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, 400, w.Code)
			},
		},

		"success": {
			reqBodyString: `{
				"first_name": "lionel",
				"last_name": "messi",
				"team_info": {
					"team_id": "a",
					"jersey_number": 10
				},
				"date_of_birth": {
					"day": 10,
					"month": 5,
					"year": 1990
				}
			}`,
			setup: func(playerServiceMock *mocks.MockPlayerService) {
				playerServiceMock.EXPECT().Create(gomock.Any()).Return(nil)
			},
			assertionFunc: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, 200, w.Code)
				var responsePlayer domain.Player
				err := json.Unmarshal(w.Body.Bytes(), &responsePlayer)
				if err != nil {
					t.Fatal("error unmarshalling response", err.Error())
				}
				assert.Equal(t, "messi", responsePlayer.LastName)
				assert.Equal(t, "a", responsePlayer.TeamInfo.TeamID)
			},
		},

		"error duplicate key": {
			reqBodyString: `{
				"first_name": "lionel",
				"last_name": "messi",
				"team_info": {
					"team_id": "a",
					"jersey_number": 10
				},
				"date_of_birth": {
					"day": 10,
					"month": 5,
					"year": 1990
				}
			}`,
			setup: func(playerServiceMock *mocks.MockPlayerService) {
				playerServiceMock.EXPECT().Create(gomock.Any()).Return(domain.NewAppError(domain.ErrCodeDuplicateKey, "duplicate key"))
			},
			assertionFunc: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, 409, w.Code)
				var appErr domain.AppError
				err := json.Unmarshal(w.Body.Bytes(), &appErr)
				if err != nil {
					t.Fatal("error unmarshalling response", err.Error())
				}
				assert.Equal(t, domain.ErrCodeDuplicateKey, appErr.Code)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playerServiceMock := mocks.NewMockPlayerService(ctrl)

	h := player.Handler{
		PlayerService: playerServiceMock,
	}

	gin.SetMode(gin.TestMode)

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {

			req, err := http.NewRequest(http.MethodPost, "/players", bytes.NewBuffer([]byte(test.reqBodyString)))
			if err != nil {
				t.Fatal("unexpected error creating request: ", err.Error())
			}
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			test.setup(playerServiceMock)

			h.CreatePlayer(c)

			test.assertionFunc(subTest, w)

		})
	}

}
