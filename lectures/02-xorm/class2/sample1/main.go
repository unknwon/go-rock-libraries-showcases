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
)

const prompt = `Please enter number of operation:
1. Create new account
2. Show detail of account
3. Deposit
4. Withdraw
5. Make transfer
6. List exist accounts by Id
7. List exist accounts by balance
8. Delete account
9. Exit`

func main() {
	fmt.Println("Welcome bank of xorm!")
Exit:
	for {
		fmt.Println(prompt)

		// 获取用户输入并转换为数字
		var num int
		fmt.Scanf("%d\n", &num)

		// 判断用户选择
		switch num {
		case 1:
			fmt.Println("Please enter <name> <balance>:")
			var name string
			var balance float64
			fmt.Scanf("%s %f\n", &name, &balance)
			if err := newAccount(name, balance); err != nil {
				fmt.Println("Fail to create new account:", err)
			} else {
				fmt.Println("New account has been created")
			}
		case 2:
			fmt.Println("Please enter <id>:")
			var id int64
			fmt.Scanf("%d\n", &id)
			a, err := getAccount(id)
			if err != nil {
				fmt.Println("Fail to get account:", err)
			} else {
				fmt.Printf("%#v\n", a)
			}
		case 3:
			fmt.Println("Please enter <id> <deposit>:")
			var id int64
			var deposit float64
			fmt.Scanf("%d %f\n", &id, &deposit)
			a, err := makeDeposit(id, deposit)
			if err != nil {
				fmt.Println("Fail to deposit:", err)
			} else {
				fmt.Printf("%#v\n", a)
			}
		case 4:
			fmt.Println("Please enter <id> <withdraw>:")
			var id int64
			var withdraw float64
			fmt.Scanf("%d %f\n", &id, &withdraw)
			a, err := makeWithdraw(id, withdraw)
			if err != nil {
				fmt.Println("Fail to withdraw:", err)
			} else {
				fmt.Printf("%#v\n", a)
			}
		case 5:
			fmt.Println("Please enter <id> <balance> <id>:")
			var id1, id2 int64
			var balance float64
			fmt.Scanf("%d %f %d\n", &id1, &balance, &id2)
			if err := makeTransfer(id1, id2, balance); err != nil {
				fmt.Println("Fail to transfer:", err)
			} else {
				fmt.Println("Transfer has been made")
			}
		case 6:
			as, err := getAccountsAscId()
			if err != nil {
				fmt.Println("Fail to get accounts:", err)
			} else {
				for i, a := range as {
					fmt.Printf("%d: %#v\n", i+1, a)
				}
			}
		case 7:
			as, err := getAccountsDescBalance()
			if err != nil {
				fmt.Println("Fail to get accounts:", err)
			} else {
				for i, a := range as {
					fmt.Printf("%d: %#v\n", i+1, a)
				}
			}
		case 8:
			fmt.Println("Please enter <id>:")
			var id int64
			fmt.Scanf("%d\n", &id)
			if err := deleteAccount(id); err != nil {
				fmt.Println("Fail to delete account:", err)
			} else {
				fmt.Println("Account has been deleted")
			}
		case 9:
			fmt.Println("Thank you! Hope see you again soon!")
			break Exit
		default:
			fmt.Println("Unknown operation number:", num)
		}
		fmt.Println()
	}
}
