package ninelaw

var OpenApiKeysSlice []OpenApiKeys = make([]OpenApiKeys, 0)

func init() {
	// 测试环境-九品调用管理人接口时Header中需要指定的参数
	var thunisoftNineLawHeaderKeys OpenApiKeys = OpenApiKeys{environment: "Thunisoft", typeDescription: "NineLaw => PoChan-GLR", appKey: "vR2eN0hL3e", securityKey: "fV1tQ3uH3bJ1yP3cQ4tS4zY0gU7vE4pL"}
	OpenApiKeysSlice = append(OpenApiKeysSlice, thunisoftNineLawHeaderKeys)
	// 测试环境-管理人调用九品接口时Feign客户端中需要指定的参数
	var thunisoftNineLawFeignKeys OpenApiKeys = OpenApiKeys{environment: "Thunisoft", typeDescription: "PoChan-GLR => NineLaw", appKey: "6Tzuo9wQ5O", securityKey: "Rx42VF95ZvILq7FtvVh8P1zTJOqWIoOW"}
	OpenApiKeysSlice = append(OpenApiKeysSlice, thunisoftNineLawFeignKeys)
	// 阿里云环境-九品调用管理人接口时Header中需要指定的参数
	var aliyunNineLawHeaderKeys OpenApiKeys = OpenApiKeys{environment: "Aliyun", typeDescription: "NineLaw => PoChan-GLR", appKey: "vR2eN0hL3e", securityKey: "fV1tQ3uH3bJ1yP3cQ4tS4zY0gU7vE4pL"}
	OpenApiKeysSlice = append(OpenApiKeysSlice, aliyunNineLawHeaderKeys)
	// 阿里云环境-管理人调用九品接口时Feign中需要指定的参数
	var aliyunNineLawFeignKeys OpenApiKeys = OpenApiKeys{environment: "Aliyun", typeDescription: "PoChan-GLR => NineLaw", appKey: "Lrrv3A5AHE", securityKey: "iTReOhTn9L3vlsjcugIGEMSiQldeCrdy"}
	OpenApiKeysSlice = append(OpenApiKeysSlice, aliyunNineLawFeignKeys)
}

type OpenApiKeys struct {
	environment     string // 环境标识
	typeDescription string // 类型描述
	appKey          string // appKey
	securityKey     string // securityKey
}

func (k *OpenApiKeys) Environment() string {
	return k.environment
}

func (k *OpenApiKeys) SetEnvironment(environment string) {
	k.environment = environment
}

func (k *OpenApiKeys) TypeDescription() string {
	return k.typeDescription
}

func (k *OpenApiKeys) SetTypeDescription(typeDescription string) {
	k.typeDescription = typeDescription
}

func (k *OpenApiKeys) AppKey() string {
	return k.appKey
}

func (k *OpenApiKeys) SetAppKey(appKey string) {
	k.appKey = appKey
}

func (k *OpenApiKeys) SecurityKey() string {
	return k.securityKey
}

func (k *OpenApiKeys) SetSecurityKey(securityKey string) {
	k.securityKey = securityKey
}
