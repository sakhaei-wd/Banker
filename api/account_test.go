package api

import (
	"testing"

	"github.com/golang/mock/gomock"
	db "github.com/sakhaei-wd/banker/db/sqlc"
	"github.com/sakhaei-wd/banker/util"
)


func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	//The first context argument could be any value, so we use gomock.Any() matcher for it. 
	//The second argument should equal to the ID of the random account we created above. So we use this matcher: gomock.Eq() and pass the account.ID to it.
	store.EXPECT().
	GetAccount(gomock.Any(), gomock.Eq(account.ID)).
	Times(1).
	Return(account, nil)
	//means : I expect the GetAccount() function of the store to be called with any context and this specific account ID arguments.

	
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
