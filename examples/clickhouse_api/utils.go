package clickhouse_api

import (
	"crypto/tls"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests"
)

func GetConnection(settings clickhouse.Settings, tlsConfig *tls.Config) (driver.Conn, error) {
	port := clickhouse_tests.GetEnv("CLICKHOUSE_PORT", "9000")
	host := clickhouse_tests.GetEnv("CLICKHOUSE_HOST", "localhost")
	username := clickhouse_tests.GetEnv("CLICKHOUSE_USERNAME", "default")
	password := clickhouse_tests.GetEnv("CLICKHOUSE_PASSWORD", "")
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr:     []string{fmt.Sprintf("%s:%s", host, port)},
		Settings: settings,
		Auth: clickhouse.Auth{
			Database: "default",
			Username: username,
			Password: password,
		},
		TLS: tlsConfig,
	})
	return conn, err
}
