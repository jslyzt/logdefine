package logDefine

import "encoding/xml"

const (
	T_INT      int8 = iota // int
	T_FLOAT                // float
	T_DOUBLE               // double
	T_STRING               // string
	T_DATETIME             // 时间日期
	T_BOOL                 // 布尔类型
	T_SHORT                // 短整型
	T_LONG                 // 长整型
	T_USERDEF              // 自定义类型
)

const (
	ET_GO   int8 = iota // 导出go
	ET_CPP              // 导出c++
	ET_JAVA             // 导出java
)

const (
	UDT_NONE  int8 = iota // 无
	UDT_LIST              // 列表
	UDT_PLIST             // 指针列表
	UDT_MAP               // 键值对
	UDT_PMAP              // 指针键值对
)

var (
	TIME_FORMATE_UNIX = "2006-01-02T15:04:05+08:00"
)

// XmlLogNode 节点定义
type XmlLogNode struct {
	Xname string `xml:"name,attr"` // 节点名字
	Name  string // 真正名字
	SType string `xml:"type,attr"` // 节点类型 -- xml
	Type  int8   // 节点类型 -- true
	//Defvalue interface{} `xml:"defaultvalue,attr"` // 节点默认值
	Desc   string `xml:"desc,attr"` // 节点说明
	UDType int8   // 扩展类型
}

// XmlLogNodes 节点数组定义
type XmlLogNodes []XmlLogNode

// XmlLogStruct 日志描述定义
type XmlLogStruct struct {
	Name    string      `xml:"name,attr"`    // 名字
	Alias   string      `xml:"alias,attr"`   // 别名
	Version int16       `xml:"version,attr"` // 版本号
	Desc    string      `xml:"desc,attr"`    // 说明
	Nodes   XmlLogNodes `xml:"entry"`        // 节点列表
}

// XmlLogStructs 日志描述数组定义
type XmlLogStructs []XmlLogStruct

// XmlLogStrMap 日志描述map定义
type XmlLogStrMap map[string]*XmlLogStruct

// XmlLogFile 日志文件定义
type XmlLogFile struct {
	file    string        // 日志文件
	XMLName xml.Name      `xml:"logs"`         // 入口节点
	Version int16         `xml:"version,attr"` // 版本号
	Name    string        `xml:"name,attr"`    // 名字
	MName   string        // 大写名字
	Stus    XmlLogStructs `xml:"struct"` // 日志结构数组
	Logs    XmlLogStructs `xml:"log"`    // 日志数组
	StuMp   XmlLogStrMap  // 日志结构map
}
