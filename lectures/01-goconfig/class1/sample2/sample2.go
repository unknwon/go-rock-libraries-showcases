// Copyright 2013-2014 Unknown
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
	"log"

	// 导入 goconfig 包
	"github.com/Unknwon/goconfig"
)

func main() {
	// 创建并获取一个 ConfigFile 对象，以进行后续操作
	// 文件名支持相对和绝对路径，可指定多个文件名进行覆盖加载
	cfg, err := goconfig.LoadConfigFile("conf.ini", "conf2.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}

	// 加载完成后所有数据均已存入内存，任何对文件的修改操作都不会影响到已经获取到的对象

	// 对默认分区进行普通读取操作
	value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "key_default")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "key_default", err)
	}
	log.Printf("%s > %s: %s", goconfig.DEFAULT_SECTION, "key_default", value)

	// 若外部文件发生修改，可通过调用方法进行快速重载
	err = cfg.Reload()
	if err != nil {
		log.Fatalf("无法重载配置文件：%s", err)
	}

	// 若在调用 Must 系列方法时发生错误，则可设置缺省值
	vBool := cfg.MustBool("must", "bool404", true)
	log.Printf("%s > %s: %v", "must", "bool404", vBool)

	// 可在操作中途追加配置文件
	err = cfg.AppendFiles("conf3.ini")
	if err != nil {
		log.Fatalf("无法追加配置文件：%s", err)
	}

	// 进行递归读取键值
	value, err = cfg.GetValue("", "search")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "search", err)
	}
	log.Printf("%s > %s: %s", goconfig.DEFAULT_SECTION, "search", value)

	// >>>>>>>>>>>>>>> 子孙分区覆盖读取 >>>>>>>>>>>>>>>

	// 以半角符号 . 为分隔符来表示多级别分区

	// 当子孙分区某个键存在时，会直接获取
	value, err = cfg.GetValue("parent.child", "age")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "age", err)
	}
	log.Printf("%s > %s: %s", "parent.child", "age", value)

	// 当子孙分区某个键不存在时，会向父区按级寻找
	value, err = cfg.GetValue("parent.child", "sex")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "sex", err)
	}
	log.Printf("%s > %s: %s", "parent.child", "sex", value)

	// <<<<<<<<<<<<<<< 子孙分区覆盖读取 <<<<<<<<<<<<<<<

	// 进行自增键名获取，凡是键名为半角符号 - 的在加载时均会被处理为自增
	// 自增范围限制在相同分区内
	// 为了方便展示，此处直接结合获取整个分区的功能并打印
	sec, err := cfg.GetSection("auto increment")
	if err != nil {
		log.Fatalf("无法获取分区：%s", err)
	}
	log.Printf("%s : %v", "auto increment", sec)

}
