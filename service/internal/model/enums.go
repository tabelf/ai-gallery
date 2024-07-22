package model

import "reflect"

// 用户状态
const (
	DISABLED = 0 // 不可用
	ACTIVE   = 1 // 可用
)

// 用户权限码
const (
	ADMIN = "ADMIN" // 管理用户
	USER  = "USER"  // 普通用户
)

type Enum struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewEnum(code, name string) *Enum {
	return &Enum{Code: code, Name: name}
}

func Values(obj any) []*Enum {
	var enums []*Enum
	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.IsNil() {
			enums = append(enums, field.Interface().(*Enum))
		}
	}
	return enums
}

// UploadServiceEnum 上传配置
type UploadServiceEnum struct {
	Local  *Enum
	Tx     *Enum
	Aliyun *Enum
}

var UploadEnums = &UploadServiceEnum{
	Local:  NewEnum("local", "本地存储"),
	Tx:     NewEnum("cos", "腾讯云"),
	Aliyun: NewEnum("oss", "阿里云"),
}

func (t *UploadServiceEnum) Values() []*Enum {
	return Values(t)
}

// SettingEnum 系统配置
type SettingEnum struct {
	StoreName    *Enum
	StoreAddress *Enum
	SecureID     *Enum
	SecureKey    *Enum
	StoreBucket  *Enum
	InitPwd      *Enum
}

var SettingEnums = &SettingEnum{
	StoreName:    NewEnum("STORE_NAME", "存储名称"),
	StoreAddress: NewEnum("STORE_ADDRESS", "存储地址"),
	SecureID:     NewEnum("SECURE_ID", "存储ID"),
	SecureKey:    NewEnum("SECURE_KEY", "存储KEY"),
	StoreBucket:  NewEnum("STORE_BUCKET", "存储桶"),
	InitPwd:      NewEnum("INIT_PWD", "初始化密码"),
}

func (t *SettingEnum) Values() []*Enum {
	return Values(t)
}

// TaskStatusEnum 任务状态.
type TaskStatusEnum struct {
	Progress  *Enum
	Complete  *Enum
	Exception *Enum
}

var TaskStatusEnums = &TaskStatusEnum{
	Progress:  NewEnum("PROGRESS", "处理中"),
	Complete:  NewEnum("COMPLETE", "已完成"),
	Exception: NewEnum("EXCEPTION", "异常"),
}

func (t *TaskStatusEnum) Values() []*Enum {
	return Values(t)
}
