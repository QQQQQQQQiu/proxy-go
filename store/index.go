package store

var PORT = "801"
var SECRET = "123"

// 创建一个映射表，用于存储变量名和其对应的变量引用
var vars = map[string]*string{
	"port":   &PORT,
	"secret": &SECRET,
}

func PortCtl(actionType string, value string) string {
	return VarCtl("port", actionType, value).(string)
}
func SecretCtl(actionType string, value string) string {
	return VarCtl("secret", actionType, value).(string)
}

// VarCtl 是一个通用的变量控制函数
func VarCtl(varName string, actionType string, value string) any {
	varRef, ok := vars[varName]
	if !ok {
		return nil // 如果变量名不存在，则返回nil
	}
	switch actionType {
		case "get":
			return *varRef
		case "set":
			*varRef = value
			fallthrough
		default:
			return *varRef
	}
}