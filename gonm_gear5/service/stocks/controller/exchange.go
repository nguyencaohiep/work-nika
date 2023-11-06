package controller

import (
	"context"
	"fmt"
	"gonm_service/service/stocks/dao"

	"cloud.google.com/go/datastore"
	"github.com/komem3/gonm"
)

var gm *gonm.Gonm

func init() {
	dsClient, err := datastore.NewClient(context.Background(), "test-gear5")
	if err != nil {
		fmt.Println("init dsclient", err)
	}
	gm = gonm.FromContext(context.Background(), dsClient)
	fmt.Println("successfully")
}

func GetExchange() {
	fmt.Println(1)
	exchange := &dao.Exchange{ExchangeCode: "US"}
	err := gm.Get(exchange)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exchange)
	fmt.Println(2)
}
