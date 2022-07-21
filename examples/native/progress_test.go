package examples

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func testProgressProfile() error {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})

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
	}))

	rows, err := conn.Query(ctx, "SELECT number from system.numbers LIMIT 10000000")
	if err != nil {
		return err
	}
	var (
		col1 uint64
	)
	for rows.Next() {
	}

	fmt.Printf("Total Rows: %d\n", totalRows)
	fmt.Printf("Total: %d\n", col1)
	rows.Close()
	return rows.Err()
}

func TestProgress(t *testing.T) {
	require.NoError(t, testProgressProfile())
}
