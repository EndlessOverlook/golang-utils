package ninelaw

var NineLawOpenApiKeysSlice []NineLawOpenApiKeys = make([]NineLawOpenApiKeys, 0)

func init() {
	// NineLawHeaderKeys, 九品调用管理人接口时Header中需要指定的参数
	var nineLawHeaderKeys NineLawOpenApiKeys = NineLawOpenApiKeys{keyType: 1, typeDescription: "NineLaw => PoChan-GLR", appKey: "vR2eN0hL3e", securityKey: "fV1tQ3uH3bJ1yP3cQ4tS4zY0gU7vE4pL"}
	NineLawOpenApiKeysSlice = append(NineLawOpenApiKeysSlice, nineLawHeaderKeys)
	// 测试环境-管理人调用九品接口时Feign中需要指定的参数
	var thunisoftNineLawFeignKeys NineLawOpenApiKeys = NineLawOpenApiKeys{keyType: 2, typeDescription: "<Thunisoft> PoChan-GLR => NineLaw", appKey: "6Tzuo9wQ5O", securityKey: "Rx42VF95ZvILq7FtvVh8P1zTJOqWIoOW"}
	NineLawOpenApiKeysSlice = append(NineLawOpenApiKeysSlice, thunisoftNineLawFeignKeys)
	// SAAS环境-管理人调用九品接口时Feign中需要指定的参数
	var saasNineLawFeignKeys NineLawOpenApiKeys = NineLawOpenApiKeys{keyType: 3, typeDescription: "<SAAS> PoChan-GLR => NineLaw", appKey: "Lrrv3A5AHE", securityKey: "iTReOhTn9L3vlsjcugIGEMSiQldeCrdy"}
	NineLawOpenApiKeysSlice = append(NineLawOpenApiKeysSlice, saasNineLawFeignKeys)
}

type NineLawOpenApiKeys struct {
	keyType         int    // 类型
	typeDescription string // 类型描述
	appKey          string // appKey
	securityKey     string // securityKey
}

func (k *NineLawOpenApiKeys) KeyType() int {
	return k.keyType
}

func (k *NineLawOpenApiKeys) SetKeyType(keyType int) {
	k.keyType = keyType
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
