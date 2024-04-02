package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/pilseong/bestprice/backend/internal/eleventst"
	"github.com/pilseong/bestprice/backend/internal/gmarket"
	"github.com/pilseong/bestprice/backend/internal/tmon"
	"github.com/pilseong/bestprice/backend/internal/wemakeprice"
)

const (
	webPort = "5006"
)

type application struct {
	GmarketDB     gmarket.DatabaseRepo
	TmonDB        tmon.DatabaseRepo
	WeMakePriceDB wemakeprice.DatabaseRepo
	ElevenStDB    eleventst.DatabaseRepo
}

func setTimezone() {
	var err error
	if tz := os.Getenv("TZ"); tz != "" {
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Printf("error loading location '%s': %v\n", tz, err)
		}
	} else {
		time.Local, err = time.LoadLocation("Asia/Seoul")
		if err != nil {
			log.Println("cannot start server because of timezone setting; ", err.Error())
		}
	}
}

func openDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

var counts int64

func connectToDB() *sql.DB {
	dataSourceName := os.Getenv("DSN")
	// dataSourceName := "postgresql://postgres:qwe123@192.168.50.141:5433/bestprice?sslmode=disable&timezone=UTC+9&connect_timeout=5"

	log.Println(dataSourceName)

	for {
		conn, err := openDB(dataSourceName)
		if err != nil {
			log.Println("cannot open postgres", err)
			counts++
		} else {
			log.Println("connected to postgres")
			return conn
		}
		if counts > 10 {
			log.Println(err)
			return nil
		}
		log.Println("Wait for 2 seconds")
		time.Sleep(2 * time.Second)

		continue
	}
}

func main() {
	setTimezone()

	conn := connectToDB()
	if conn == nil {
		log.Fatal("cannot connect to DB")
	}

	app := application{}
	app.GmarketDB = &gmarket.PostgresDBRepo{DB: conn}
	app.TmonDB = &tmon.PostgresDBRepo{DB: conn}
	app.WeMakePriceDB = &wemakeprice.PostgresDBRepo{DB: conn}
	app.ElevenStDB = &eleventst.PostgresDBRepo{DB: conn}

	defer app.GmarketDB.Connection().Close()
	defer app.TmonDB.Connection().Close()
	defer app.WeMakePriceDB.Connection().Close()
	defer app.ElevenStDB.Connection().Close()

	log.Printf("Starting best price service on port %s\n", webPort)

	// define http server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
