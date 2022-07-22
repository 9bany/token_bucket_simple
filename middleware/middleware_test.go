package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addUserType(
	t *testing.T,
	request *http.Request,
	userType string,
	userTypeValue string,
) {
	request.Header.Set(userType, userTypeValue)
}

func TestMiddleWare(t *testing.T) {
	testCases := []struct {
		name          string
		setupUserType func(t *testing.T, request *http.Request)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupUserType: func(t *testing.T, request *http.Request) {
				addUserType(t, request, "user-type", "gen-user")
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "user user type default",
			setupUserType: func(t *testing.T, request *http.Request) {
				
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			// setup server route
			router := gin.Default()
			authPath := "/auth"
			router.GET(authPath, Middleware, func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{})
			})

			// request
			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupUserType(t, request)
			router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
