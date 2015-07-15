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
	"log"

	// 导入 goconfig 包
	"github.com/Unknwon/goconfig"
)

func main() {
	// 创建并获取一个 ConfigFile 对象，以进行后续操作
	// 文件名支持相对和绝对路径
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}

	// 加载完成后所有数据均已存入内存，任何对文件的修改操作都不会影响到已经获取到的对象

	// >>>>>>>>>>>>>>> 基本读写操作 >>>>>>>>>>>>>>>

	// 对默认分区进行普通读取操作
	value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "key_default")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "key_default", err)
	}
	log.Printf("%s > %s: %s", goconfig.DEFAULT_SECTION, "key_default", value)

	// 对已有的键进行值重写操作，返回值为 bool 类型，表示是否为插入操作
	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION, "key_default", "这是新的值")
	log.Printf("设置键值 %s 为插入操作：%v", "key_default", isInsert)

	// 对不存在的键进行插入操作
	isInsert = cfg.SetValue(goconfig.DEFAULT_SECTION, "key_new", "这是新插入的键")
	log.Printf("设置键值 %s 为插入操作：%v", "key_new", isInsert)

	// 传入空白字符串也可直接操作默认分区
	value, err = cfg.GetValue("", "key_default")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "key_default", err)
	}
	log.Printf("%s > %s: %s", goconfig.DEFAULT_SECTION, "key_default", value)

	// 获取冒号为分隔符的键值
	value, err = cfg.GetValue("super", "key_super2")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "key_super2", err)
	}
	log.Printf("%s > %s: %s", "super", "key_super2", value)

	// <<<<<<<<<<<<<<< 基本读写操作 <<<<<<<<<<<<<<<

	// >>>>>>>>>>>>>>> 对注释进行读写操作 >>>>>>>>>>>>>>>

	// 获取某个分区的注释
	comment := cfg.GetSectionComments("super")
	log.Printf("分区 %s 的注释：%s", "super", comment)

	// 获取某个键的注释
	comment = cfg.GetKeyComments("super", "key_super")
	log.Printf("键 %s 的注释：%s", "key_super", comment)

	// 设置某个键的注释，返回值为 true 时表示注释被插入或删除（空字符串），false 表示注释被重写
	v := cfg.SetKeyComments("super", "key_super", "# 这是新的键注释")
	log.Printf("键 %s 的注释被插入或删除：%v", "key_super", v)

	// 设置某个分区的注释，返回值效果同上
	v = cfg.SetSectionComments("super", "# 这是新的分区注释")
	log.Printf("分区 %s 的注释被插入或删除：%v", "super", v)

	// <<<<<<<<<<<<<<< 对注释进行读写操作 <<<<<<<<<<<<<<<

	// >>>>>>>>>>>>>>> 自动转换和 Must 系列方法 >>>>>>>>>>>>>>>

	// 自动转换类型读取操作，直接返回指定类型，error 类型用于指示是否发生错误
	vInt, err := cfg.Int("must", "int")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "int", err)
	}
	log.Printf("%s > %s: %v", "must", "int", vInt)

	// Must 系列方法，一定返回某个类型的值；如果失败则返回零值
	vBool := cfg.MustBool("must", "bool")
	log.Printf("%s > %s: %v", "must", "bool", vBool)

	// 若键不存在则返回零值，此例应返回 false
	vBool = cfg.MustBool("must", "bool404")
	log.Printf("%s > %s: %v", "must", "bool404", vBool)

	// <<<<<<<<<<<<<<< 自动转换和 Must 系列方法 <<<<<<<<<<<<<<<

	// 删除键值，返回值用于表示是否删除成功
	ok := cfg.DeleteKey("must", "string")
	log.Printf("删除键值 %s 是否成功：%v", "string", ok)

	// 保存 ConfigFile 对象到文件系统，保存后的键顺序与读取时的一样
	err = goconfig.SaveConfigFile(cfg, "conf_save.ini")
	if err != nil {
		log.Fatalf("无法保存配置文件：%s", err)
	}
}
