package transport

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rogudator/get-nrows-service/internal/service"
	mock_service "github.com/rogudator/get-nrows-service/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetRows(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockRowsGetter)

	testTable := []struct {
		name                 string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK with n=1",
			mockBehaviour: func(s *mock_service.MockRowsGetter) {
				s.EXPECT().GetRows().Return([]string{"John Doe"})
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"names\":[\"John Doe\"]}",
		},
		{
			name: "OK with n=3",
			mockBehaviour: func(s *mock_service.MockRowsGetter) {
				s.EXPECT().GetRows().Return([]string{"John Doe", "Sarah Williams", "Ava Smith"})
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"names\":[\"John Doe\",\"Sarah Williams\",\"Ava Smith\"]}",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			rows := mock_service.NewMockRowsGetter(c)
			testCase.mockBehaviour(rows)

			serv := &service.Service{
				RowsGetter: rows,
			}
			rest := NewTransport(serv)

			router := gin.New()
			router.GET("/rows", rest.getRows)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/rows?n=1", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
		})
	}
}
