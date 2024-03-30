package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pilseong/tomonagent/models"
)

var db *sql.DB

// setDB initializes the sql object
func setDB(dbPool *sql.DB) {
	db = dbPool
}

type Models struct {
	Item models.Item
}

func (app *Config) BulkCreate(items []models.Item) (int, error) {
	// size := 500
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		log.Println("cannot start transaction. rollbacked : " + err.Error())
		return 0, err
	}

	// 배치 저장
	stmt := "INSERT INTO tmon.batch (id) VALUES ($1)"

	batchId := time.Now().UnixMilli()
	_, err = tx.Exec(stmt, batchId)
	if err != nil {
		tx.Rollback()
		log.Println("cannot insert batch metadata. rollbacked : " + err.Error())
		return 0, err
	}

	valueStrings := []string{}

	valueArgs := make([]interface{}, 0, 13)
	for index, item := range items {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, batchId)
		valueArgs = append(valueArgs, index+1)
		valueArgs = append(valueArgs, item.StartDate)
		valueArgs = append(valueArgs, item.EndDate)
		valueArgs = append(valueArgs, item.Title)
		valueArgs = append(valueArgs, item.DealNo)
		valueArgs = append(valueArgs, item.DeliveryInfo.DeliveryFeePolicy)
		valueArgs = append(valueArgs, item.DeliveryInfo.DeliveryFee)
		valueArgs = append(valueArgs, item.CategoryInfo.CategoryName)
		valueArgs = append(valueArgs, item.ImageInfo.Pc3ColImageURL)
		valueArgs = append(valueArgs, item.DiscountPrice.OriginalPrice)
		valueArgs = append(valueArgs, item.DiscountPrice.Price)
		valueArgs = append(valueArgs, item.DiscountPrice.DiscountRate)
	}

	stmt = fmt.Sprintf("INSERT INTO tmon.items ("+
		"batchId, "+
		"rank, "+
		"startDate, "+
		"endDate, "+
		"title, "+
		"dealNo, "+
		"deliveryFeePolicy, "+
		"deliveryFee, "+
		"categoryName, "+
		"pc3ColImageUrl, "+
		"originalPrice, "+
		"price, "+
		"discountRate "+
		") VALUES %s", strings.Join(valueStrings, ","))

	stmt = ReplaceSQL(stmt)
	res, err := tx.Exec(stmt, valueArgs...)

	if err != nil {
		tx.Rollback()
		log.Println("cannot insert item information. rollbacked : " + err.Error())
		return 0, err
	}

	size, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Println("cannot parse result. rollbacked : " + err.Error())
		return 0, err
	}

	// 최신 데이터 배치키 저장
	stmt = `UPDATE tmon.latest 
					SET batchid = $1
					WHERE id = 1	`

	_, err = tx.Exec(stmt, batchId)
	if err != nil {
		tx.Rollback()
		log.Println("cannot update latest batch id. rollbacked : " + err.Error())
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("cannot commit changes. rollbacked : " + err.Error())
		return 0, err
	}

	return int(size), nil
}

func ReplaceSQL(stmt string) string {
	n := 0

	for strings.IndexByte(stmt, '?') != -1 {
		n++
		param := "$" + strconv.Itoa(n)
		stmt = strings.Replace(stmt, "?", param, 1)
	}
	return strings.TrimSuffix(stmt, ",")
}
