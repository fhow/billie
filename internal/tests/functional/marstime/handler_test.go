package marstime

import (
	"github.com/billie/internal/services/marstime/converter"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "false: UTC time unset",
			in:             httptest.NewRequest("GET", "/convert", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: 400,
			expectedBody: `{"error":"You should provide UTC time in format 2006-01-02T15:04:05Z07:00"}
`,
		},
		{
			name:           "false: UTC time unset",
			in:             httptest.NewRequest("GET", "/convert?UTC=2120p-08-30T04:24:27Z", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: 400,
			expectedBody: `{"error":"parsing time \"2120p-08-30T04:24:27Z\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"p-08-30T04:24:27Z\" as \"-\""}
`,
		},
		{
			name:           "true: Time converted",
			in:             httptest.NewRequest("GET", "/convert?UTC=2120-08-30T04:24:27Z", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: 200,
			expectedBody: `{"MSD":87683.1616857292,"MTC":"3:52:50"}
`,
		},
	}
	req := require.New(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			converter.Handler(test.out, test.in)

			req.Equal(test.out.Code, test.expectedStatus)
			req.Equal(test.out.Body.String(), test.expectedBody)

		})
	}
}
