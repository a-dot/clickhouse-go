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
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOpenTelemetry(t *testing.T) {
	require.NoError(t, openTelemetry())
}

func TestTuples(t *testing.T) {
	require.NoError(t, tupleInsertRead())
}

func TestAppendStruct(t *testing.T) {
	require.NoError(t, appendStruct())
}

func TestArrayInsertRead(t *testing.T) {
	require.NoError(t, arrayInsertRead())
}

func TestAsyncInsert(t *testing.T) {
	require.NoError(t, asyncInsert())
}

func TestBatchInsert(t *testing.T) {
	require.NoError(t, batchInsert())
}

func TestAuthConnect(t *testing.T) {
	_, err := authVersion()
	require.NoError(t, err)
}

func TestBigInt(t *testing.T) {
	require.NoError(t, readWriteBigInt())
}

func TestBind(t *testing.T) {
	require.NoError(t, bindParameters())
}

func TestSpecialCaseBind(t *testing.T) {
	require.NoError(t, specialBind())
}

func TestColumnInsert(t *testing.T) {
	require.NoError(t, columnInsert())
}

func TestConnect(t *testing.T) {
	_, err := version()
	require.NoError(t, err)
}

func TestZSTDCompression(t *testing.T) {
	require.NoError(t, compress())
}

func TestConnectWithSettings(t *testing.T) {
	require.NoError(t, pingWithSettings())
}

func TestDecimal(t *testing.T) {
	require.NoError(t, readWriteDecimal())
}

func TestContext(t *testing.T) {
	require.NoError(t, useContext())
}

func TestDynamicScan(t *testing.T) {
	require.NoError(t, dynamicScan())
}

func TestExternalTable(t *testing.T) {
	require.NoError(t, externalData())
}

func TestExec(t *testing.T) {
	require.NoError(t, createCreateDrop())
}

func TestGeo(t *testing.T) {
	require.NoError(t, testGeo())
}

func TestJSON(t *testing.T) {
	require.NoError(t, insertReadJSON())
	require.NoError(t, readComplexJSON())
}

func TestMapInsertRead(t *testing.T) {
	require.NoError(t, mapInsertRead())
}

func TestMultiHostConnect(t *testing.T) {
	_, err := multiHostVersion()
	require.NoError(t, err)
	_, err = multiHostRoundRobinVersion()
	require.NoError(t, err)
}

func TestNested(t *testing.T) {
	require.NoError(t, nestedUnFlattened())
	require.NoError(t, nestedFlattened())
}

func TestProgress(t *testing.T) {
	require.NoError(t, progressProfileLogs())
}

func TestScanStruct(t *testing.T) {
	require.NoError(t, scanStruct())
}

func TestQueryRow(t *testing.T) {
	require.NoError(t, queryRow())
}

func TestSelectStruct(t *testing.T) {
	require.NoError(t, selectStruct())
}

func TestSSL(t *testing.T) {
	_, err := sslVersion()
	require.NoError(t, err)
}

func TestSSLCustomCerts(t *testing.T) {
	_, err := sslCustomCertsVersion()
	require.NoError(t, err)
}

func TestTypeConvert(t *testing.T) {
	require.NoError(t, convertedInsert())
}

func TestUUID(t *testing.T) {
	require.NoError(t, testUUID())
}

func TestNullable(t *testing.T) {
	require.NoError(t, testNullable())
}

func TestQueryRows(t *testing.T) {
	require.NoError(t, queryRows())
}

func TestSSLNoVerify(t *testing.T) {
	_, err := sslNoVerifyVersion()
	require.NoError(t, err)
}
