package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/Golang-Tanzania/azampay"
)

func InitAzamPay() {
	appName :=  os.Getenv("AZAMPAY_APPNAME")
	clientId := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTID_SECR_KEY")
	tokenKey := os.Getenv("TOKEN")
	client, err := azampay.NewClient(appName, clientId, clientSecret, tokenKey)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = client.GetAccessToken(ctx)

	if err != nil {
		fmt.Println("Error in obtaining the tokens....")
		fmt.Println(err)
	}

}
