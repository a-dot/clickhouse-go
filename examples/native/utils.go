package examples

import (
	"crypto/tls"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getConnection(settings clickhouse.Settings, tlsConfig *tls.Config) (driver.Conn, error) {
	port := GetEnv("CLICKHOUSE_PORT", "9000")
	host := GetEnv("CLICKHOUSE_HOST", "localhost")
	username := GetEnv("CLICKHOUSE_USERNAME", "default")
	password := GetEnv("CLICKHOUSE_PASSWORD", "")
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
