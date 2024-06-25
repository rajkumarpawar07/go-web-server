package tools

import (
	log "github.com/sirupsen/logrus"
)


// DatabaseInterface is an interface that defines the methods that a database should implement
type LoginDetails struct {
	AuthToken string
	Username string
}
type CoinDetails struct {
	Coins int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserUserCoins(username string) *CoinDetails
	SetupDatabase() error
}


func NewDatabase() (*DatabaseInterface, error) {
	
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}