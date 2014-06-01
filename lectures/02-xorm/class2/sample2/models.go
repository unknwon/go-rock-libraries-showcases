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
	"errors"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

// 银行账户
type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` // 乐观锁
}

func (a *Account) BeforeInsert() {
	log.Printf("before insert: %s", a.Name)
}

func (a *Account) AfterInsert() {
	log.Printf("after insert: %s", a.Name)
}

// ORM 引擎
var x *xorm.Engine

func init() {
	// 创建 ORM 引擎与数据库
	var err error
	x, err = xorm.NewEngine("sqlite3", "./bank.db")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}

	// 同步结构体与数据表
	if err = x.Sync(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}

	// 记录日志
	f, err := os.Create("sql.log")
	if err != nil {
		log.Fatalf("Fail to create log file: %v\n", err)
		return
	}
	x.Logger = xorm.NewSimpleLogger(f)
	x.ShowSQL = true

	// 设置默认 LRU 缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	x.SetDefaultCacher(cacher)
}

// 创建新的账户
func newAccount(name string, balance float64) error {
	// 对未存在记录进行插入
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
}

// 获取账户数量
func getAccountCount() (int64, error) {
	return x.Count(new(Account))
}

// 获取账户信息
func getAccount(id int64) (*Account, error) {
	a := &Account{}
	// 直接操作 ID 的简便方法
	has, err := x.Id(id).Get(a)
	// 判断操作是否发生错误或对象是否存在
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Account does not exist")
	}
	return a, nil
}
