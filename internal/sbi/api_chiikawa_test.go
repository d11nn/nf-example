package sbi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Alonza0314/nf-example/internal/sbi"
	"github.com/Alonza0314/nf-example/pkg/factory"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func newChiikawaTestServer(t *testing.T) *sbi.Server {
	t.Helper()

	gin.SetMode(gin.TestMode)

	mockCtrl := gomock.NewController(t)
	nfApp := sbi.NewMocknfApp(mockCtrl)
	nfApp.EXPECT().Config().Return(&factory.Config{
		Configuration: &factory.Configuration{
			Sbi: &factory.Sbi{
				Port: 8000,
			},
		},
	}).AnyTimes()

	return sbi.NewServer(nfApp, "")
}

func Test_HTTPGetChiikawaVoice(t *testing.T) {
	server := newChiikawaTestServer(t)

	t.Run("Find Voice That Exists", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const EXPECTED_BODY = `{"name":"Usagi","voice":"Una una una! Yaha!"}`

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)
		ginCtx.Params = gin.Params{
			{
				Key:   "Name",
				Value: "Usagi",
			},
		}

		server.HTTPGetChiikawaVoice(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if strings.TrimSpace(httpRecorder.Body.String()) != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("Find Voice That Does Not Exist", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusNotFound
		const EXPECTED_BODY = "Kurimanju voice not found"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)
		ginCtx.Params = gin.Params{
			{
				Key:   "Name",
				Value: "Kurimanju",
			},
		}

		server.HTTPGetChiikawaVoice(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})
}

func Test_HTTPCreateChiikawaVoice(t *testing.T) {
	server := newChiikawaTestServer(t)

	t.Run("Create Voice", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusCreated
		const EXPECTED_BODY = `{"message":"Chiikawa voice saved","name":"Hachiware","voice":"Nanka wakannai kedo wakatta!"}`

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest(
			http.MethodPost,
			"/chiikawa/voice",
			strings.NewReader(`{"name":"Hachiware","voice":"Nanka wakannai kedo wakatta!"}`),
		)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}
		ginCtx.Request.Header.Set("Content-Type", "application/json")

		server.HTTPCreateChiikawaVoice(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if strings.TrimSpace(httpRecorder.Body.String()) != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusBadRequest
		const EXPECTED_BODY = "Invalid request body"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest(
			http.MethodPost,
			"/chiikawa/voice",
			strings.NewReader(`{"name":"Usagi"}`),
		)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}
		ginCtx.Request.Header.Set("Content-Type", "application/json")

		server.HTTPCreateChiikawaVoice(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})
}
