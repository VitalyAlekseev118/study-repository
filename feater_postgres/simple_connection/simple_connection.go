package simpleconnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	con, err := pgx.Connect(ctx, "host=localhost port=5432 user=postgres password=NeroVit777$ dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := con.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подклюение к базе данных произошло успешно!")
}
