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

func Test_HTTPRegisterUe(t *testing.T) {
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
	server := sbi.NewServer(nfApp, "")

	testCases := []struct {
		name         string
		body         string
		expectedBody string
	}{
		{
			name:         "Invalid JSON body",
			body:         `{"supi":`,
			expectedBody: "Invalid JSON body",
		},
		{
			name:         "Supi without imsi prefix",
			body:         `{"supi":"208930000000001"}`,
			expectedBody: "supi must start with imsi-",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpRecorder := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(httpRecorder)

			var err error
			ginCtx.Request, err = http.NewRequest("POST", "/ue/", strings.NewReader(tc.body))
			if err != nil {
				t.Errorf("Failed to create request: %s", err)
				return
			}
			ginCtx.Request.Header.Set("Content-Type", "application/json")

			server.HTTPRegisterUe(ginCtx)

			if httpRecorder.Code != http.StatusBadRequest {
				t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, httpRecorder.Code)
			}
			if httpRecorder.Body.String() != tc.expectedBody {
				t.Errorf("Expected body %s, got %s", tc.expectedBody, httpRecorder.Body.String())
			}
		})
	}
}
