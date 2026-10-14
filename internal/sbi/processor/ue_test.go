package processor_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	nf_context "github.com/Alonza0314/nf-example/internal/context"
	"github.com/Alonza0314/nf-example/internal/sbi/processor"
	"github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

const testSupi = "imsi-208930000000001"

func newTestUeContext() *nf_context.NFContext {
	return &nf_context.NFContext{
		UeData: map[string]nf_context.UeInfo{
			testSupi: {Supi: testSupi, GnbId: "gnb-1", Dnn: "internet"},
		},
	}
}

func Test_GetUe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCtrl := gomock.NewController(t)
	processorNf := processor.NewMockProcessorNf(mockCtrl)
	p, err := processor.NewProcessor(processorNf)
	if err != nil {
		t.Errorf("Failed to create processor: %s", err)
		return
	}

	testCases := []struct {
		name           string
		supi           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "UE That Exists",
			supi:           testSupi,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"supi":"` + testSupi + `","gnbId":"gnb-1","dnn":"internet"}`,
		},
		{
			name:           "UE That Does Not Exist",
			supi:           "imsi-208930000000999",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "[imsi-208930000000999] not registered",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processorNf.EXPECT().Context().Return(newTestUeContext())

			httpRecorder := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(httpRecorder)
			p.GetUe(ginCtx, tc.supi)

			if httpRecorder.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, httpRecorder.Code)
			}
			if httpRecorder.Body.String() != tc.expectedBody {
				t.Errorf("Expected body %s, got %s", tc.expectedBody, httpRecorder.Body.String())
			}
		})
	}
}

func Test_RegisterUe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCtrl := gomock.NewController(t)
	processorNf := processor.NewMockProcessorNf(mockCtrl)
	p, err := processor.NewProcessor(processorNf)
	if err != nil {
		t.Errorf("Failed to create processor: %s", err)
		return
	}

	t.Run("Register New UE", func(t *testing.T) {
		const newSupi = "imsi-208930000000002"
		ctx := newTestUeContext()
		processorNf.EXPECT().Context().Return(ctx)

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)
		p.RegisterUe(ginCtx, nf_context.UeInfo{Supi: newSupi, GnbId: "gnb-2", Dnn: "internet"})

		if httpRecorder.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, httpRecorder.Code)
		}
		if _, ok := ctx.UeData[newSupi]; !ok {
			t.Errorf("Expected %s to be stored in context", newSupi)
		}
	})

	t.Run("Register Duplicate UE", func(t *testing.T) {
		const expectedBody = "[" + testSupi + "] already registered"
		processorNf.EXPECT().Context().Return(newTestUeContext())

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)
		p.RegisterUe(ginCtx, nf_context.UeInfo{Supi: testSupi})

		if httpRecorder.Code != http.StatusConflict {
			t.Errorf("Expected status code %d, got %d", http.StatusConflict, httpRecorder.Code)
		}
		if httpRecorder.Body.String() != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, httpRecorder.Body.String())
		}
	})
}
