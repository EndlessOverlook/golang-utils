package ninelaw

type NineLawOpenApiKeys struct {
	typeDescription string // 类型描述
	appKey          string // appKey
	securityKey     string // securityKey
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

// 九品调用管理人接口时Header中需要指定的参数
var NineLawHeaderKeys NineLawOpenApiKeys = NineLawOpenApiKeys{typeDescription: "NineLaw => PoChan-GLR", appKey: "vR2eN0hL3e", securityKey: "fV1tQ3uH3bJ1yP3cQ4tS4zY0gU7vE4pL"}

// 管理人调用九品接口时Feign中需要指定的参数
var NineLawFeignKeys NineLawOpenApiKeys = NineLawOpenApiKeys{typeDescription: "PoChan-GLR => NineLaw", appKey: "6Tzuo9wQ5O", securityKey: "Rx42VF95ZvILq7FtvVh8P1zTJOqWIoOW"}
