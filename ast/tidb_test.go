/*
 * Copyright 2018 Xiaomi, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	"fmt"
	"testing"

	"github.com/XiaoMi/soar/common"
)

func TestPrintPrettyStmtNode(t *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	sqls := []string{
		`select 1`,
		`select * f`, // syntax error case
	}
	err := common.GoldenDiff(func() {
		for _, sql := range sqls {
			PrintPrettyStmtNode(sql, "", "")
		}
	}, t.Name(), update)
	if nil != err {
		t.Fatal(err)
	}
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}

func TestStmtNode2JSON(t *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	sqls := []string{
		`select 1`,
		`select * f`, // syntax error case
	}
	err := common.GoldenDiff(func() {
		for _, sql := range sqls {
			fmt.Println(StmtNode2JSON(sql, "", ""))
		}
	}, t.Name(), update)
	if nil != err {
		t.Fatal(err)
	}
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}

func TestSchemaMetaInfo(t *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	err := common.GoldenDiff(func() {
		for _, sql := range common.TestSQLs {
			fmt.Println(sql)
			fmt.Println(SchemaMetaInfo(sql, "sakila"))
		}
	}, t.Name(), update)
	if nil != err {
		t.Fatal(err)
	}
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}
