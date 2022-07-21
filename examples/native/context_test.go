// Licensed to ClickHouse, Inc. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. ClickHouse, Inc. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package examples

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
	"time"
)

func testContext() error {
	var (
		dialCount int

		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"127.0.0.1:9000"},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: "",
			},
			DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
				dialCount++
				var d net.Dialer
				return d.DialContext(ctx, "tcp", addr)
			},
		})
	)
	if err != nil {
		return err
	}
	if err := clickhouse_tests.CheckMinServerVersion(conn, 22, 6, 1); err != nil {
		return nil
	}

	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"allow_experimental_object_type": "1",
	}))
	conn.Exec(ctx, "DROP TABLE IF EXISTS example")

	// to create a JSON column we need allow_experimental_object_type=1
	if err = conn.Exec(ctx, `
		CREATE TABLE example (
				Col1 JSON
			) 
			Engine Memory
		`); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		cancel()
	}()
	start := time.Now()
	// query is cancelled with context
	if err = conn.QueryRow(ctx, "SELECT sleep(3)").Scan(); err == nil {
		return fmt.Errorf("expected cancel")
	}
	fmt.Printf("cancelled after %v and %d dial\n", time.Since(start), dialCount)

	// set a deadline for a query
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer cancel()
	if err := conn.Ping(ctx); err == nil {
		return fmt.Errorf("expected deadline exceeeded")
	}
	fmt.Printf("deadline exceeded %s\n", err)

	return nil
}

func TestContext(t *testing.T) {
	require.NoError(t, testContext())
}
