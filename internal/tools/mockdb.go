package tools

import (
	"time"
)

type mockDB struct {}


var mockLoginDetails = map[string]LoginDetails{
	"alex":{
		AuthToken: "1234",
		Username: "alex",
	},
	"john":{
		AuthToken: "5678",
		Username: "john",
	},
	"jane":{
		AuthToken: "91011",
		Username: "jane",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex":{
		Coins: 100,
		Username: "alex",
	},
	"john":{
		Coins: 200,
		Username: "john",
	},
	"jane":{
		Coins: 300,
		Username: "jane",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails{
	time.Sleep(1 * time.Second)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	} 
	return &clientData
}

func (d *mockDB) GetUserUserCoins(username string) *CoinDetails{
	time.Sleep(1 * time.Second)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	} 
	return &clientData
}

func (d *mockDB) SetupDatabase() error{
	return nil
}
