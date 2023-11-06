package ninelaw

var NineLawOpenApiKeysSlice []NineLawOpenApiKeys = make([]NineLawOpenApiKeys, 0)

func init() {
	// 测试环境-九品调用管理人接口时Header中需要指定的参数
	var thunisoftNineLawHeaderKeys NineLawOpenApiKeys = NineLawOpenApiKeys{environment: "Thunisoft", typeDescription: "NineLaw => PoChan-GLR", appKey: "vR2eN0hL3e", securityKey: "fV1tQ3uH3bJ1yP3cQ4tS4zY0gU7vE4pL"}
	NineLawOpenApiKeysSlice = append(NineLawOpenApiKeysSlice, thunisoftNineLawHeaderKeys)
	// 测试环境-管理人调用九品接口时Feign客户端中需要指定的参数
	var thunisoftNineLawFeignKeys NineLawOpenApiKeys = NineLawOpenApiKeys{environment: "Thunisoft", typeDescription: "PoChan-GLR => NineLaw", appKey: "6Tzuo9wQ5O", securityKey: "Rx42VF95ZvILq7FtvVh8P1zTJOqWIoOW"}
	NineLawOpenApiKeysSlice = append(NineLawOpenApiKeysSlice, thunisoftNineLawFeignKeys)
	// 阿里云环境-九品调用管理人接口时Header中需要指定的参数
	var aliyunNineLawHeaderKeys NineLawOpenApiKeys = NineLawOpenApiKeys{environment: "Aliyun", typeDescription: "NineLaw => PoChan-GLR", appKey: "vR2eN0hL3e", securityKey: "fV1tQ3uH3bJ1yP3cQ4tS4zY0gU7vE4pL"}
	NineLawOpenApiKeysSlice = append(NineLawOpenApiKeysSlice, aliyunNineLawHeaderKeys)
	// 阿里云环境-管理人调用九品接口时Feign中需要指定的参数
	var aliyunNineLawFeignKeys NineLawOpenApiKeys = NineLawOpenApiKeys{environment: "Aliyun", typeDescription: "PoChan-GLR => NineLaw", appKey: "Lrrv3A5AHE", securityKey: "iTReOhTn9L3vlsjcugIGEMSiQldeCrdy"}
	NineLawOpenApiKeysSlice = append(NineLawOpenApiKeysSlice, aliyunNineLawFeignKeys)
}

type NineLawOpenApiKeys struct {
	environment     string // 环境标识
	typeDescription string // 类型描述
	appKey          string // appKey
	securityKey     string // securityKey
}

func (k *NineLawOpenApiKeys) Environment() string {
	return k.environment
}

func (k *NineLawOpenApiKeys) SetEnvironment(environment string) {
	k.environment = environment
}

func (k *NineLawOpenApiKeys) TypeDescription() string {
	return k.typeDescription
}

func (k *NineLawOpenApiKeys) SetTypeDescription(typeDescription string) {
	k.typeDescription = typeDescription
}

func (k *NineLawOpenApiKeys) AppKey() string {
	return k.appKey
}

func (k *NineLawOpenApiKeys) SetAppKey(appKey string) {
	k.appKey = appKey
}

func (k *NineLawOpenApiKeys) SecurityKey() string {
	return k.securityKey
}

func (k *NineLawOpenApiKeys) SetSecurityKey(securityKey string) {
	k.securityKey = securityKey
}
