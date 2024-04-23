package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// const INTERVAL_PERIOD time.Duration = 2 * time.Hour

// const INTERVAL_PERIOD time.Duration = 10 * time.Second

var counts int64

type jobTicker struct {
	t *time.Timer
}

type Config struct {
	DB *sql.DB
}

func getNextTickDuration() time.Duration {
	interval := getInterval()

	log.Printf("Interval time is set to every %v \n", interval)

	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)

	// 현재 시간 보다 이전일 경우 주기를 더한다.
	if nextTick.Before(now) {
		nextTick = nextTick.Add(time.Duration(interval))
	}
	return time.Until(nextTick)
}

// 환경 변수에서 읽어온다.
// 환경 변수가 없으면 10초
// 환경변수는 기본 hour 단위
func getInterval() time.Duration {
	interval_str := os.Getenv("INTERVAL")

	var interval time.Duration
	if interval_str == "" {
		interval = 20 * time.Second
	} else {
		interval_int, err := strconv.Atoi(interval_str)
		if err != nil {
			log.Println("error occurred while getting the INTERVAL TIME :" + err.Error())
		}
		interval = time.Duration(interval_int) * time.Hour
	}
	return interval
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
			dataSourceName = "postgresql://postgres:qwe123@192.168.50.59:5432/bestprice?sslmode=disable&timezone=UTC+9&connect_timeout=5"
			log.Println("trying to connect to", dataSourceName)
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
		log.Println("Starting data fetching from 11st")

		app.getData()

		log.Println("Data fetching has finished successfully")
		jt.updateJobTicker()
	}
}
