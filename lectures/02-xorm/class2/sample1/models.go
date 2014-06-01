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
}

// 创建新的账户
func newAccount(name string, balance float64) error {
	// 对未存在记录进行插入
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
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

// 用户存款
func makeDeposit(id int64, deposit float64) (*Account, error) {
	a, err := getAccount(id)
	if err != nil {
		return nil, err
	}
	a.Balance += deposit
	// 对已有记录进行更新
	_, err = x.Update(a)
	return a, err
}

// 用户取款
func makeWithdraw(id int64, withdraw float64) (*Account, error) {
	a, err := getAccount(id)
	if err != nil {
		return nil, err
	}
	if a.Balance < withdraw {
		return nil, errors.New("Not enough balance")
	}
	a.Balance -= withdraw
	_, err = x.Update(a)
	return a, err
}

// 用户转账
func makeTransfer(id1, id2 int64, balance float64) error {
	a1, err := getAccount(id1)
	if err != nil {
		return err
	}

	a2, err := getAccount(id2)
	if err != nil {
		return err
	}

	if a1.Balance < balance {
		return errors.New("Not enough balance")
	}

	a1.Balance -= balance
	a2.Balance += balance
	// 创建 Session 对象
	sess := x.NewSession()
	defer sess.Close()
	// 启动事务
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Update(a1); err != nil {
		// 发生错误时进行回滚
		sess.Rollback()
		return err
	} else if _, err = sess.Update(a2); err != nil {
		sess.Rollback()
		return err
	}
	// 完成事务
	return sess.Commit()
}

// 按照 ID 正序排序返回所有账户
func getAccountsAscId() (as []Account, err error) {
	// 使用 Find 方法批量获取记录
	err = x.Find(&as)
	return as, err
}

// 按照存款倒序排序返回所有账户
func getAccountsDescBalance() (as []Account, err error) {
	// 使用 Desc 方法使结果呈倒序排序
	err = x.Desc("balance").Find(&as)
	return as, err
}

// 删除账户
func deleteAccount(id int64) error {
	// 通过 Delete 方法删除记录
	_, err := x.Delete(&Account{Id: id})
	return err
}
