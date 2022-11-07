package main

import (
	"context"
	"fmt"
	"log"
	"test-postgres/config"
	"test-postgres/custom_user/db"
	"test-postgres/postgressql"
)

func main() {
	PostgresSQLClient, err := postgressql.New(context.TODO(), 3, config.GetPostgersConfig())
	if err != nil {
		log.Fatal(err)
	}
	r := db.NewRepository(PostgresSQLClient)
	u, err := r.FindAll(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range u {
		fmt.Printf("%v", v)
	}
}
