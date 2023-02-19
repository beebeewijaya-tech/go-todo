package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	db "database/sql"

	"beebeewijaya.com/db/sql"
	"beebeewijaya.com/db/sql/mock_sql"
	"beebeewijaya.com/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	failedSecretKey = "failedSecret"
)

func TestUserCreate(t *testing.T) {
	// Correct and expected body request
	successfulBody := createUserParams{
		Email:    util.RandomEmail(),
		Password: util.RandomString(8),
		Fullname: util.RandomString(32),
	}

	// Bad Request - Invalid missing full name
	invalidMissingField := createUserParams{
		Email:    util.RandomEmail(),
		Password: util.RandomString(32),
	}

	// Bad Request - Invalid password longer than 72 chars
	badPassword := createUserParams{
		Email:    util.RandomEmail(),
		Password: util.RandomString(80),
		Fullname: util.RandomString(32),
	}

	userResponse := sql.User{
		Email:    successfulBody.Email,
		Password: successfulBody.Password,
		Fullname: successfulBody.Fullname,
	}

	testCases := []struct {
		name          string
		body          createUserParams
		mockSql       func(*mock_sql.MockQueries)
		checkResponse func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: successfulBody,
			mockSql: func(mq *mock_sql.MockQueries) {
				mq.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(userResponse, nil)
			},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, response.Code)
				checkResponseUserCreated(t, response.Body, userResponse)
			},
		},
		{
			name:    "BadRequest",
			body:    invalidMissingField,
			mockSql: func(mq *mock_sql.MockQueries) {},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, response.Code)
			},
		},
		{
			name:    "BadPassword",
			body:    badPassword,
			mockSql: func(mq *mock_sql.MockQueries) {},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, response.Code)
			},
		},
		{
			name: "InternalServerError",
			body: successfulBody,
			mockSql: func(mq *mock_sql.MockQueries) {
				mq.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.User{}, db.ErrTxDone)
			},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, response.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSql := mock_sql.NewMockQueries(ctrl)
			tc.mockSql(mockSql)

			server := NewServer(mockSql, c)
			w := httptest.NewRecorder()

			b, err := json.Marshal(tc.body)
			require.NoError(t, err)
			body := bytes.NewReader(b)

			url := "/user/create"
			request, err := http.NewRequest(http.MethodPost, url, body)
			require.NoError(t, err)

			server.router.ServeHTTP(w, request)
			tc.checkResponse(t, w)
		})
	}
}

func TestLogin(t *testing.T) {
	// Correct and expected body request
	successfulBody := loginUserParams{
		Email:    util.RandomEmail(),
		Password: util.RandomString(8),
	}

	hashed, err := util.HashPassword(successfulBody.Password)
	require.NoError(t, err)

	// Bad Request - Invalid missing full name
	invalidMissingField := loginUserParams{
		Password: util.RandomString(32),
	}

	testCases := []struct {
		name          string
		body          loginUserParams
		mockSql       func(*mock_sql.MockQueries)
		checkResponse func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name:    "BadRequest",
			body:    invalidMissingField,
			mockSql: func(mq *mock_sql.MockQueries) {},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, response.Code)
			},
		},
		{
			name: "UserNotFound",
			body: successfulBody,
			mockSql: func(mq *mock_sql.MockQueries) {
				mq.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.User{}, db.ErrNoRows)
			},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, response.Code)
			},
		},
		{
			name: "InvalidPassword",
			body: successfulBody,
			mockSql: func(mq *mock_sql.MockQueries) {
				mq.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.User{
						Email:    successfulBody.Email,
						Fullname: util.RandomString(32),
						Password: util.RandomString(8),
					}, nil)
			},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, response.Code)
			},
		},
		{
			name: "OK",
			body: successfulBody,
			mockSql: func(mq *mock_sql.MockQueries) {
				mq.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.User{
						Email:    successfulBody.Email,
						Fullname: util.RandomString(32),
						Password: hashed,
					}, nil)
			},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, response.Code)
				checkResponseUserLogin(t, response.Body)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSql := mock_sql.NewMockQueries(ctrl)
			tc.mockSql(mockSql)

			server := NewServer(mockSql, c)
			w := httptest.NewRecorder()

			b, err := json.Marshal(tc.body)
			require.NoError(t, err)
			body := bytes.NewReader(b)

			url := "/user/login"
			request, err := http.NewRequest(http.MethodPost, url, body)

			require.NoError(t, err)

			server.router.ServeHTTP(w, request)
			tc.checkResponse(t, w)
		})
	}
}

func checkResponseUserCreated(t *testing.T, body *bytes.Buffer, user sql.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser *sql.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)

	require.Equal(t, gotUser.Fullname, user.Fullname)
	require.Equal(t, gotUser.Email, user.Email)
}

func checkResponseUserLogin(t *testing.T, body *bytes.Buffer) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser *loginUserResponse
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)

	require.NotEmpty(t, gotUser)
	require.NotEmpty(t, gotUser.Email)
	require.GreaterOrEqual(t, len(gotUser.Token), 1)
}
