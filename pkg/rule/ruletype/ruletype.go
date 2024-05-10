package ruletype

type DictType struct {
	Wifi      []string
	Directory []string
	Url       []string
	WeakPwd   []string
	Other     []string
}

// 预定义一些特定场景下的规则 减少无效字典生成 精准字典

// Wifi

// Directory 目录字典

// login

// xss

// sql

// rce

// lfi

// ssrf

// xxe

// api

// 文件包含

// 命令执行

// 文件上传

// weak 端口 弱密码

// apply 应用 设备 弱密码

// 学校 特定学校 特定学院 特定专业 特定班级 学号

// 电话号码

// 身份证

// 年份

// 特定 特定时间 特定日期

// 子域名

// cms 不同框架 的 字典 [][]string

// windows linux 不同系统下的 字典 [][]string
