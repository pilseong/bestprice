package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// const INTERVAL_PERIOD time.Duration = 2 * time.Hour

const INTERVAL_PERIOD time.Duration = 10 * time.Second

var counts int64

type jobTicker struct {
	t *time.Timer
}

type Config struct {
	DB *sql.DB
}

func getNextTickDuration() time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)

	// 현재 시간 보다 이전일 경우 주기를 더한다.
	if nextTick.Before(now) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	return time.Until(nextTick)
}

func NewJobTicker() jobTicker {
	log.Println("Created New Scheduler")
	return jobTicker{time.NewTimer(getNextTickDuration())}
}

func (jt jobTicker) updateJobTicker() {
	log.Println("Set the next schdule")
	jt.t.Reset(getNextTickDuration())
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

func connectToDB() *sql.DB {
	dataSourceName := os.Getenv("DSN")

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

func main() {

	// manually set time zone
	setTimezone()

	conn := connectToDB()
	if conn == nil {
		log.Fatal("cannot connect to DB")
	}

	app := Config{
		DB: conn,
	}

	setDB(conn)

	// app.mockData()
	jt := NewJobTicker()
	for {
		<-jt.t.C
		log.Println("Starting data fetching from tmon")

		app.getData()

		log.Println("Data fetching has finished successfully")
		jt.updateJobTicker()
	}
}
