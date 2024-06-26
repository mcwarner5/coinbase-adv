package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/mcwarner5/BlockBot8000/libraries/coinbase-adv/client"
)

func Test_GetAccount(t *testing.T) {
	ctx := context.Background()
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	limit := int32(1)
	p := client.ListAccountsParams{
		Limit: &limit,
	}

	rsp, err := cln.ListAccounts(ctx, &p)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	// should be at least 1
	if len(rsp.Accounts) < 1 {
		t.FailNow()
	}

	// Get account
	acct, err := cln.GetAccount(ctx, *rsp.Accounts[0].Uuid)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	fmt.Printf("[%s] [%s] %f\n", acct.GetUuid(), acct.GetCurrency(), acct.AvailableBalance.GetValue())
}
