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
func GenerateNinelawOpenAPISecurityParameters(appKey string, securityKey string, timestamp string, random string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	parameterSlice := make([]string, 4)
	parameterSlice = append(parameterSlice, appKey)
	parameterSlice = append(parameterSlice, securityKey)
	parameterSlice = append(parameterSlice, timestamp)
	parameterSlice = append(parameterSlice, random)
	sort.Strings(parameterSlice)
	parameters := strings.Join(parameterSlice, "")
	encryptedParameters := encryptionutils.EncryptWithSha256(parameters)
	sugar.Info("Ninelaw OpenAPI Security Parameters:")
	sugar.Info("========================================")
	sugar.Infow("[appKey & securityKey]", zap.String("appKey", appKey), zap.String("securityKey", securityKey))
	sugar.Info("========================================")
	sugar.Infow("[timestamp]", zap.String("timestamp", timestamp))
	sugar.Infow("[random]", zap.String("random", random))
	sugar.Infow("[sign]", zap.String("sign", encryptedParameters))
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
