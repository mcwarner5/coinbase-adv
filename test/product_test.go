package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/mcwarner5/BlockBot8000/libraries/coinbase-adv/client"
)

func Test_GetProduct(t *testing.T) {
	ctx := context.Background()
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	product := "BTC-USD"

	cln := client.NewClient(&creds)

	rsp, err := cln.GetProduct(ctx, product)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	fmt.Printf("%s,  price = %f, vol = %f,  status = %s", product, *rsp.Price, *rsp.Volume24h, *rsp.Status)

}

func Test_GetProduct_CheckRateLimitRetry(t *testing.T) {
	ctx := context.Background()
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	product := "BTC-USD"

	cln := client.NewClient(&creds)

	// wipe this to zero, we're testing handling of httpStatus = 429 (too many requests)
	// the client should backoff for 1s and retry transparently.
	cln.SetRateLimit(0)

	iter := 20
	for iter > 0 {
		_, err := cln.GetProduct(ctx, product)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		// fmt.Printf("%s,  price = %f, vol = %f,  status = %s", product, *rsp.Price, *rsp.Volume24h, *rsp.Status)
		iter--
	}

}
