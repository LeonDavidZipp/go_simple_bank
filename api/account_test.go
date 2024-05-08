package api


func randomAccount() db.Account {
	return db.Account {
		ID: util.RandomInt(1, 1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomBalance(),
		Currency: util.RandomCurrency(),
		CreatedAt: time.Now(),
	}
}

func TestGetAccountAPI(t *testing.T) {

}