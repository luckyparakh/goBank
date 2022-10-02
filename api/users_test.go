package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/luckyparakh/goBank/db/mock"
	db "github.com/luckyparakh/goBank/db/sqlc"
	"github.com/luckyparakh/goBank/utils"
	"github.com/stretchr/testify/require"
)

// type eqCreateUserParamMatcher struct {
// 	args db.CreateUserParams
// 	password string
// }

// func (e eqCreateUserParamMatcher) Matches(x interface{}) bool {
// 	return reflect.DeepEqual(e.x, x)
// }

//	func (e eqCreateUserParamMatcher) String() string {
//		return fmt.Sprintf("is equal to %v", e.x)
//	}
func TestCreateUser(t *testing.T) {
	user, password := randomUser(t)
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "ok",
			body: gin.H{
				"username":  user.Username,
				"email":     user.Email,
				"password":  password,
				"full_name": user.FullName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
	}
	for id := range testCases {
		tc := testCases[id]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			body, err := json.Marshal(tc.body)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func randomUser(t *testing.T) (db.User, string) {
	password := utils.GenerateRandomString(6)
	hp, err := utils.HashedPassword(password)
	require.NoError(t, err)

	return db.User{
		Username:       utils.GenerateRandomString(6),
		Email:          utils.GenerateRandomEmail(),
		FullName:       utils.GenerateRandomOwner(),
		HashedPassword: hp,
	}, password
}
func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	var gotUser db.User
	require.NoError(t, err)
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
