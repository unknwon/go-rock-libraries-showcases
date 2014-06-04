// Copyright 2014 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"log"
)

var printFn = func(idx int, bean interface{}) error {
	fmt.Printf("%d: %#v\n", idx, bean.(*Account))
	return nil
}

func main() {
	fmt.Println("Welcome bank of xorm!")

	count, err := getAccountCount()
	if err != nil {
		log.Fatalf("Fail to get account count: %v", err)
	}
	fmt.Println("Account count:", count)

	// 自动创建至 10 个账户
	for i := count; i < 10; i++ {
		if err = newAccount(fmt.Sprintf("joe%d", i), float64(i)*100); err != nil {
			log.Fatalf("Fail to create account: %v", err)
		}
	}

	// 迭代查询
	fmt.Println("Query all columns:")
	x.Iterate(new(Account), printFn)

	// 更灵活的迭代
	a := new(Account)
	rows, err := x.Rows(a)
	if err != nil {
		log.Fatalf("Fail to rows: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(a); err != nil {
			log.Fatalf("Fail get row: %v", err)
		}
		fmt.Printf("%#v\n", a)
	}

	// 查询特定字段
	fmt.Println("\nOnly query name:")
	x.Cols("name").Iterate(new(Account), printFn)

	// 排除特定字段
	fmt.Println("\nQuery all but name:")
	x.Omit("name").Iterate(new(Account), printFn)

	// 查询结果偏移
	fmt.Println("\nOffest 2 and limit 3:")
	x.Limit(3, 2).Iterate(new(Account), printFn)

	// 测试 LRU 缓存
	getAccount(1)
	getAccount(1)
}
