package main

import (
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/zlx2019/ecom/cmd/api"
	"github.com/zlx2019/ecom/config"
	"github.com/zlx2019/ecom/store"
)

func main() {
	// init storage
	store, err := store.NewStore(mysql.Config{
		User:                 config.Envs.Store.Username,
		Passwd:               config.Envs.Store.Password,
		Addr:                 config.Envs.Store.Addr,
		DBName:               config.Envs.Store.DBName,
		Net:                  "tcp",
		Timeout:              time.Second * 5,
		ParseTime:            true,
		AllowNativePasswords: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := api.NewAPIServer(config.Envs.Addr(), store).Startup(); err != nil {
		log.Fatalf("APIServer startup fail: %v", err)
	}
}
