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
)

func testNullable() error {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"127.0.0.1:9000"},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: "",
			},
		})
	)
	if err != nil {
		return err
	}
	conn.Exec(ctx, "DROP TABLE IF EXISTS example")

	if err = conn.Exec(ctx, `
		CREATE TABLE example (
				col1 Nullable(String),
				col2 String
			) 
			Engine Memory
		`); err != nil {
		return err
	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO example")
	if err != nil {
		return err
	}
	if err = batch.Append(
		nil,
		nil,
	); err != nil {
		return err
	}

	if err = batch.Send(); err != nil {
		return err
	}

	var (
		col1 *string
		col2 string
	)

	if err = conn.QueryRow(ctx, "SELECT * FROM example").Scan(&col1, &col2); err != nil {
		return err
	}
	fmt.Printf("col1=%v, col2=%v\n", col1, col2)
	return nil
}
