package db

import "github.com/ClickHouse/clickhouse-go/v2"

func Connect() (*clickhouse.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"host.docker.internal:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
	})
	if err != nil {
		return nil, err
	}

	return &conn, nil
}
