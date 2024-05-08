package api


import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"io"
	"encoding/json"
	gomock "go.uber.org/mock/gomock"
	"github.com/stretchr/testify/require"
	db "github.com/LeonDavidZipp/go_simple_bank/db/sqlc"
	mockdb "github.com/LeonDavidZipp/go_simple_bank/db/mock"
	"github.com/LeonDavidZipp/go_simple_bank/util"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()
	
	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, account)
}

func randomAccount() db.Account {
	return db.Account {
		ID: util.RandomInt(1, 1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
