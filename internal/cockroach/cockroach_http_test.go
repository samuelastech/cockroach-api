package cockroach_test

import (
	"encoding/json"
	"errors"
	"github.com/samulastech/cockroach/internal/cockroach"
	"github.com/samulastech/cockroach/internal/cockroach/mocks"
	"github.com/samulastech/cockroach/internal/entities"
	"github.com/samulastech/cockroach/pkg/common"
	"github.com/samulastech/cockroach/pkg/testutils"
	testMock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type SuiteCockroachHTTPHandler struct {
	suite.Suite
	handler       cockroach.CockroachHandler
	usecaseCreate *mocks.CockroachUsecaseCreateMock
}

func (s *SuiteCockroachHTTPHandler) SetupTest() {
	s.usecaseCreate = mocks.NewCockroachUsecaseCreateMock()
	s.handler = cockroach.NewCockroachHTTPHandler(s.usecaseCreate)
}

func (s *SuiteCockroachHTTPHandler) TearDownTest() {
	s.usecaseCreate.AssertExpectations(s.T())
}

func TestCockroachHTTPHandler(t *testing.T) {
	suite.Run(t, new(SuiteCockroachHTTPHandler))
}

func (s *SuiteCockroachHTTPHandler) TestCreateCockroach() {
	type (
		want struct {
			body string
			code int
		}

		args = func() (*httptest.ResponseRecorder, *http.Request)

		mock = func()

		test struct {
			name string
			args args
			want want
			mock mock
		}
	)

	tests := []test{
		{
			name: "should create a cockroach HTTP 200",
			want: want{code: http.StatusOK, body: testutils.ObjectToJSON(common.Response{
				Code:    http.StatusOK,
				Message: "success",
				Data:    "🪳",
			})},
			args: func() (res *httptest.ResponseRecorder, req *http.Request) {
				body := testutils.ObjectToJSON(entities.CreateCockroachDTO{})
				res = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/cockroach", strings.NewReader(body))
				return
			},
			mock: func() {
				s.usecaseCreate.
					On("DataProcessing", testMock.AnythingOfType("*entities.CreateCockroachDTO")).
					Return(nil)
			},
		},
		{
			name: "should return an error if the payload is bad HTTP 400",
			want: want{code: http.StatusBadRequest, body: testutils.ObjectToJSON(common.Response{
				Code:    http.StatusBadRequest,
				Message: "invalid body",
				Data:    nil,
				Error:   true,
			})},
			args: func() (res *httptest.ResponseRecorder, req *http.Request) {
				res = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/cockroach", strings.NewReader("wrong body"))
				return
			},
			mock: func() {},
		},
		{
			name: "should return an error if something wrong happened HTTP 500",
			want: want{code: http.StatusInternalServerError, body: testutils.ObjectToJSON(common.Response{
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
				Data:    nil,
				Error:   true,
			})},
			args: func() (res *httptest.ResponseRecorder, req *http.Request) {
				body, _ := json.Marshal(entities.CreateCockroachDTO{})
				res = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/cockroach", strings.NewReader(string(body)))
				return
			},
			mock: func() {
				s.usecaseCreate.
					On("DataProcessing", testMock.AnythingOfType("*entities.CreateCockroachDTO")).
					Return(errors.New("unknown error"))
			},
		},
	}

	for _, scenario := range tests {
		s.Run(scenario.name, func() {
			scenario.mock()
			res, req := scenario.args()
			s.handler.CreateCockroach(res, req)
			s.JSONEq(scenario.want.body, res.Body.String())
			s.Equal(scenario.want.code, res.Code)
		})
	}
}
