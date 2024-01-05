package database

import (
	"context"
	"os"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Dbinstance struct{
	Db *pgxpool.Pool
}

var DB Dbinstance

func ConnectDb(){
	username :=  os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := "postgres"
	port := "5432"
	database :=  os.Getenv("POSTGRES_DB")
	// Connection string
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", username, password, host, port, database)

	con, err := pgxpool.New(context.Background(), connString)
	if(err!=nil){
		// log db error while conneting
		fmt.Printf("Error while connecting to DB %s", err)
	}

	defer con.Close()

	// check the connection
	err = con.Ping(context.Background())
	if err != nil {
		// connection not established
		fmt.Printf("Error while establishing connection to DB %s", err)
	}

	// message connected to Postgresql
	fmt.Printf("Connected to postgresql")

	DB = Dbinstance{
		Db: con,
	}
}