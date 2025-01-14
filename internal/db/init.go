package db

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func Init(conn *clickhouse.Conn) error {
	if err := (*conn).Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS data (
		ticker String,
		price Float32,
		timestamp DateTime DEFAULT now(),
	) Engine = MergeTree()
	PARTITION BY toYYYYMMDD(timestamp)
	ORDER BY (ticker, timestamp);`); err != nil {
		return err
	}

	return nil
}
