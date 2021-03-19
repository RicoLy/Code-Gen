package models

const (
	PosReq    = "/param/req/"
	PosRes    = "/param/res/"
	PosVo     = "/param/vo/"
	PosErr    = "/param/erro/"
	PosGlobal = "/param/global/"
)
const (
	PosEndpoint  = "/src/endpoint/"
	PosService   = "/src/service/"
	PosTransport = "/src/transport/"
	PosDao       = "/src/dao/"
)
const (
	PosMain   = "/main/"
	PosClient = "/client/"
)

// proto信息
type Message struct {
	MessageMeta  string        // 元数据
	MessageName  string        // 消息名
	ElementInfos []ElementInfo // proto字段
}

// 结构体信息
type StructInfo struct {
	StructName   string
	ElementInfos []ElementInfo // golang字段
	HasWrite     bool          // 是否写入文件
}

// 字段信息
type ElementInfo struct {
	Name string            // 名称
	Type string            // golang 数据类型
	Tags map[string]string // 标签信息 tag | value  json:"pid" form:"pid"
}

// 方法信息 rpc[\s]*([\w]*)\((\w*)\)\s*\w*\s*\((\w*)\)
type Method struct {
	MethodName string // 方法名
	Param      string // 参数
	Returns    string // 返回值
}

// 服务信息  service\s*(\w*)\s*{
type ProtoService struct {
	ServiceName string
	Methods     []Method
}
