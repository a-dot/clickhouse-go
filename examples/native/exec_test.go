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
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func createCreateDrop() error {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
	})
	if err != nil {
		return err
	}
	defer func() {
		conn.Exec(context.Background(), "DROP TABLE example")
	}()
	conn.Exec(context.Background(), `DROP TABLE IF EXISTS example`)
	err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS example (
			Col1 UInt8,
			Col2 String
		) engine=Memory
	`)
	if err != nil {
		return err
	}
	return conn.Exec(context.Background(), "INSERT INTO example VALUES (1, 'test-1')")
}

func TestExec(t *testing.T) {
	require.NoError(t, createCreateDrop())
}
