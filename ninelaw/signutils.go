package ninelaw

import (
	"bytes"
	"go.uber.org/zap"
	"golang-utils/encryptionutils"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 生成九品签名
func GenerateNinelawOpenAPISecurityParameters(keyType byte, timestamp string, random string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	parameterSlice := make([]string, 4)
	var typeDescription string
	var appKey string
	var securityKey string
	switch keyType {
	case 1:
		typeDescription = NineLawHeaderKeys.TypeDescription()
		appKey = NineLawHeaderKeys.AppKey()
		securityKey = NineLawHeaderKeys.SecurityKey()
	case 2:
		typeDescription = NineLawFeignKeys.TypeDescription()
		appKey = NineLawFeignKeys.AppKey()
		securityKey = NineLawFeignKeys.SecurityKey()
	}
	parameterSlice = append(parameterSlice, appKey)
	parameterSlice = append(parameterSlice, securityKey)
	parameterSlice = append(parameterSlice, timestamp)
	parameterSlice = append(parameterSlice, random)
	sort.Strings(parameterSlice)
	parameters := strings.Join(parameterSlice, "")
	encryptedParameters := encryptionutils.EncryptWithSha256(parameters)

	sugar.Infof("Ninelaw OpenAPI Security Parameters:[%s]", typeDescription)
	sugar.Infow("[01]", zap.String("appKey", appKey), zap.String("securityKey", securityKey))
	sugar.Infow("[02]", zap.String("timestamp", timestamp))
	sugar.Infow("[03]", zap.String("random", random))
	sugar.Infow("[04]", zap.String("sign", encryptedParameters))
	sugar.Info("==============================")
}

// 随机生成"英文字母、数字形式"的字符串
func GenerateRandom(length int) string {
	randomBase := "abcdefghijklmnopqrstuvwxyz0123456789"
	var builder bytes.Buffer
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		random := rand.Intn(len(randomBase))
		char := string([]rune(randomBase)[random])
		//fmt.Printf("随机数: %d, 对应字符: %s\n", random, char)
		builder.WriteString(char)
	}
	return builder.String()
}
