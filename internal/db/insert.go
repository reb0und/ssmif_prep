package db

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2"

	"ssmif_prep/internal/yf"
)

func Insert(conn *clickhouse.Conn, data *yf.Data) error {
	if err := (*conn).Exec(context.Background(), `
	INSERT INTO data (ticker, price) VALUES (?, ?)`, data.Config.Ticker, data.Price); err != nil {
		return err
	}

	return nil
}
