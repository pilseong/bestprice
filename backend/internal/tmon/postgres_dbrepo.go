package tmon

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

// Connection returns underlying connection pool.
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllItems() ([]*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			batchid
		from
		tmon.latest
		where id = 1
	`

	row, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Failed to execute to get the latest batch id: " + err.Error())
		return nil, err
	}
	defer row.Close()

	var latest Latest
	for row.Next() {
		err = row.Scan(
			&latest.BatchId,
		)
	}

	log.Printf("batch id is : %d\n", latest.BatchId)

	if err != nil {
		log.Println("Failed to get the lastest batch id from result batch id: " + err.Error())
		return nil, err
	}

	query = `
		select
			id, rank, title, deliveryFeePolicy, 
			deliveryFee, categoryName, url, pc3ColImageUrl, originalPrice, price, discountRate
		from
			tmon.items
		where
		  batchid = $1
		order by
			rank
	`

	rows, err := m.DB.QueryContext(ctx, query, latest.BatchId)
	if err != nil {
		log.Printf("Failed to get the items of batch id %d : %v", latest.BatchId, err.Error())
		return nil, err
	}
	defer rows.Close()

	var items []*Item

	// for i := 0; rows.Next() && i < 50; i++ {
	for rows.Next() {
		var item Item
		err := rows.Scan(
			&item.ID,
			&item.Rank,
			&item.Title,
			&item.DeliveryFeePolicy,
			&item.DeliveryFee,
			&item.CategoryName,
			&item.URL,
			&item.Pc3ColImageURL,
			&item.OriginalPrice,
			&item.Price,
			&item.DiscountRate,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	return items, nil
}
