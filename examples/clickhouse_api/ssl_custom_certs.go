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

package clickhouse_api

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"io/ioutil"
)

func SSLCustomCertsVersion() error {
	t := &tls.Config{}
	caCert, err := ioutil.ReadFile("play.clickhouse.pem")
	if err != nil {
		return err
	}
	caCertPool := x509.NewCertPool()
	successful := caCertPool.AppendCertsFromPEM(caCert)
	if !successful {
		return err
	}
	t.RootCAs = caCertPool

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"play.clickhouse.com:9440"},
		TLS:  t,
		Auth: clickhouse.Auth{
			Username: "explorer",
		},
	})
	if err != nil {
		return err
	}
	v, err := conn.ServerVersion()
	if err != nil {
		return err
	}
	fmt.Println(v.String())
	return nil
}
