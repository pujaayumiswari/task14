package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect(){
	databaseUrl := "postgres://postgres:puja@localhost:5432/b47-s1"
	var  err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil{
	fmt.Fprintf(os.Stderr, "unable to connect to database:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("successfully connected to database")
}