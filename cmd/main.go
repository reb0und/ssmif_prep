package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"

	"ssmif_prep/internal/db"
	"ssmif_prep/internal/yf"
)

func insert(conn *clickhouse.Conn, ticker string) error {
	data, err := yf.Fetch(&yf.TickerConfig{Ticker: ticker, Period: "1d", Interval: "1m"})
	if err != nil {
		return err
	}

	if err = db.Insert(conn, data); err != nil {
		return fmt.Errorf("failed to insert to ClickHouse: %v", err)
	}

	log.Printf("%s: %f", data.Config.Ticker, data.Price)

	return nil
}

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to ClickHouse: %v", err)
	}

	defer (*conn).Close()

	if err = (*conn).Ping(context.Background()); err != nil {
		log.Fatalf("failed to ping ClickHouse: %v", err)
	}

	if err = db.Init(conn); err != nil {
		log.Fatalf("failed to insert table: %v", err)
	}
	for {
		if err = insert(conn, "AAPL"); err != nil {
			log.Printf("Error: %v", err)
		}

		time.Sleep(2 * time.Minute)
	}
}
