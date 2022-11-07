package main

import (
	"context"
	"fmt"
	"log"
	customuser "test-postgres/custom_user"
	db "test-postgres/custom_user/db"
	"test-postgres/postgressql"
)

func main() {
	PostgresSQLClient, err := postgressql.New(context.TODO(), 3)
	if err != nil {
		log.Fatal(err)
	}
	r := db.NewRepository(PostgresSQLClient)
	var user customuser.CustomUser = customuser.CustomUser{Age: 19, Name: "Микита"}
	err = r.Create(context.TODO(), user)

	if err != nil {
		log.Fatal(err)
	}

	u, err := r.FindAll(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range u {
		fmt.Printf("%v\n", v)
	}
}
