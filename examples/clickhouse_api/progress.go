package clickhouse_api

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
)

func ProgressProfileLogs() error {
	conn, err := GetConnection(clickhouse.Settings{
		"send_logs_level": "trace",
	}, nil)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	totalRows := uint64(0)
	// use context to pass a call back for progress and profile info
	ctx := clickhouse.Context(context.Background(), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
		totalRows += p.Rows
	}), clickhouse.WithProfileInfo(func(p *clickhouse.ProfileInfo) {
		fmt.Println("profile info: ", p)
	}), clickhouse.WithLogs(func(log *clickhouse.Log) {
		fmt.Println("log info: ", log)
	}))

	rows, err := conn.Query(ctx, "SELECT number from numbers(10000000) LIMIT 10000000")
	if err != nil {
		return err
	}
	for rows.Next() {
	}

	fmt.Printf("Total Rows: %d\n", totalRows)
	rows.Close()
	return rows.Err()
}
