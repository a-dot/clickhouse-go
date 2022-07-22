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
	"github.com/ClickHouse/clickhouse-go/v2/examples/clickhouse_api"
	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"path"
	"strings"
	"testing"
)

const defaultClickHouseVersion = "latest"

func GetClickHouseTestVersion() string {
	return clickhouse_tests.GetEnv("CLICKHOUSE_VERSION", defaultClickHouseVersion)
}

func TestMain(m *testing.M) {
	useDocker := strings.ToLower(clickhouse_tests.GetEnv("CLICKHOUSE_USE_DOCKER", "true"))
	if useDocker == "false" {
		fmt.Printf("Using external ClickHouse for IT tests -  %s:%s\n",
			clickhouse_tests.GetEnv("CLICKHOUSE_PORT", "9000"), clickhouse_tests.GetEnv("CLICKHOUSE_HOST", "localhost"))
		os.Exit(m.Run())
	}
	// create a ClickHouse container
	ctx := context.Background()
	// attempt use docker for CI
	provider, err := testcontainers.ProviderDocker.GetProvider()
	if err != nil {
		fmt.Printf("Docker is not running and no clickhouse connections details were provided. Skipping tests: %s\n", err)
		os.Exit(0)
	}
	err = provider.Health(ctx)
	if err != nil {
		fmt.Printf("Docker is not running and no clickhouse connections details were provided. Skipping IT tests: %s\n", err)
		os.Exit(0)
	}
	fmt.Printf("Using Docker for IT tests\n")
	cwd, err := os.Getwd()
	if err != nil {
		// can't test without container
		panic(err)
	}
	req := testcontainers.ContainerRequest{
		Image:        fmt.Sprintf("clickhouse/clickhouse-server:%s", GetClickHouseTestVersion()),
		ExposedPorts: []string{"9000/tcp", "8123/tcp"},
		WaitingFor:   wait.ForLog("Ready for connections"),
		Mounts: []testcontainers.ContainerMount{
			testcontainers.BindMount(path.Join(cwd, "../tests/resources/custom.xml"), "/etc/clickhouse-server/config.d/custom.xml"),
			testcontainers.BindMount(path.Join(cwd, "../tests/resources/admin.xml"), "/etc/clickhouse-server/users.d/admin.xml"),
		},
	}
	clickhouseContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		// can't test without container
		panic(err)
	}

	p, _ := clickhouseContainer.MappedPort(ctx, "9000")
	os.Setenv("CLICKHOUSE_PORT", p.Port())
	os.Setenv("CLICKHOUSE_HOST", "localhost")
	defer clickhouseContainer.Terminate(ctx) //nolint
	os.Exit(m.Run())
}

func TestJSON(t *testing.T) {
	require.NoError(t, clickhouse_api.InsertReadJSON())
	require.NoError(t, clickhouse_api.ReadComplexJSON())
}

func TestOpenTelemetry(t *testing.T) {
	require.NoError(t, clickhouse_api.OpenTelemetry())
}

func TestTuples(t *testing.T) {
	require.NoError(t, clickhouse_api.TupleInsertRead())
}

func TestAppendStruct(t *testing.T) {
	require.NoError(t, clickhouse_api.AppendStruct())
}

func TestArrayInsertRead(t *testing.T) {
	require.NoError(t, clickhouse_api.ArrayInsertRead())
}

func TestAsyncInsert(t *testing.T) {
	require.NoError(t, clickhouse_api.AsyncInsert())
}

func TestBatchInsert(t *testing.T) {
	require.NoError(t, clickhouse_api.BatchInsert())
}

func TestAuthConnect(t *testing.T) {
	require.NoError(t, clickhouse_api.AuthVersion())
}

func TestBigInt(t *testing.T) {
	require.NoError(t, clickhouse_api.ReadWriteBigInt())
}

func TestBind(t *testing.T) {
	require.NoError(t, clickhouse_api.BindParameters())
}

func TestSpecialCaseBind(t *testing.T) {
	require.NoError(t, clickhouse_api.SpecialBind())
}

func TestColumnInsert(t *testing.T) {
	require.NoError(t, clickhouse_api.ColumnInsert())
}

func TestConnect(t *testing.T) {
	require.NoError(t, clickhouse_api.Version())
}

func TestZSTDCompression(t *testing.T) {
	require.NoError(t, clickhouse_api.Compress())
}

func TestConnectWithSettings(t *testing.T) {
	require.NoError(t, clickhouse_api.PingWithSettings())
}

func TestDecimal(t *testing.T) {
	require.NoError(t, clickhouse_api.ReadWriteDecimal())
}

func TestContext(t *testing.T) {
	require.NoError(t, clickhouse_api.UseContext())
}

func TestDynamicScan(t *testing.T) {
	require.NoError(t, clickhouse_api.DynamicScan())
}

func TestExternalTable(t *testing.T) {
	require.NoError(t, clickhouse_api.ExternalData())
}

func TestExec(t *testing.T) {
	require.NoError(t, clickhouse_api.CreateCreateDrop())
}

func TestGeo(t *testing.T) {
	require.NoError(t, clickhouse_api.GeoInsertRead())
}

func TestMapInsertRead(t *testing.T) {
	require.NoError(t, clickhouse_api.MapInsertRead())
}

func TestMultiHostConnect(t *testing.T) {
	require.NoError(t, clickhouse_api.MultiHostVersion())
	require.NoError(t, clickhouse_api.MultiHostRoundRobinVersion())
}

func TestNested(t *testing.T) {
	require.NoError(t, clickhouse_api.NestedUnFlattened())
	require.NoError(t, clickhouse_api.NestedFlattened())
}

func TestProgress(t *testing.T) {
	require.NoError(t, clickhouse_api.ProgressProfileLogs())
}

func TestScanStruct(t *testing.T) {
	require.NoError(t, clickhouse_api.ScanStruct())
}

func TestQueryRow(t *testing.T) {
	require.NoError(t, clickhouse_api.QueryRow())
}

func TestSelectStruct(t *testing.T) {
	require.NoError(t, clickhouse_api.SelectStruct())
}

func TestTypeConvert(t *testing.T) {
	require.NoError(t, clickhouse_api.ConvertedInsert())
}

func TestUUID(t *testing.T) {
	require.NoError(t, clickhouse_api.UUIDInsertRead())
}

func TestNullable(t *testing.T) {
	require.NoError(t, clickhouse_api.NullableInsertRead())
}

func TestQueryRows(t *testing.T) {
	require.NoError(t, clickhouse_api.QueryRows())
}

func TestSSL(t *testing.T) {
	require.NoError(t, clickhouse_api.SSLVersion())
}

func TestSSLCustomCerts(t *testing.T) {
	require.NoError(t, clickhouse_api.SSLCustomCertsVersion())
}

func TestSSLNoVerify(t *testing.T) {
	require.NoError(t, clickhouse_api.SSLNoVerifyVersion())
}
