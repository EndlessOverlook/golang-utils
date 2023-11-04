package ninelaw

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang-utils/encryptionutils"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

// 生成九品签名
func GenerateNinelawOpenAPISecurityParameters(keyType byte, timestamp string, random string) {
	// logger, _ := zap.NewProduction()
	// defer logger.Sync()
	// sugar := logger.Sugar()

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.DebugLevel)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// writeSyncer := zapcore.AddSync(&lumberjack.Logger{
	// 	// 日志名称
	// 	Filename: "ninelaw-sign.log",
	// 	// 日志大小限制，单位MB
	// 	MaxSize: 100,
	// 	// 历史日志文件保留天数
	// 	MaxAge: 30,
	// 	// 最大保留历史日志数量
	// 	MaxBackups: 10,
	// 	// 本地时区
	// 	LocalTime: true,
	// 	// 历史日志文件压缩标识
	// 	Compress: false,
	// })

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), atomicLevel),
		// zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writeSyncer, atomicLevel),
	)
	logger := zap.New(core)
	sugar := logger.Sugar()
	defer logger.Sync()

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
		typeDescription = ThunisoftNineLawFeignKeys.TypeDescription()
		appKey = ThunisoftNineLawFeignKeys.AppKey()
		securityKey = ThunisoftNineLawFeignKeys.SecurityKey()
	case 3:
		typeDescription = SaasNineLawFeignKeys.TypeDescription()
		appKey = SaasNineLawFeignKeys.AppKey()
		securityKey = SaasNineLawFeignKeys.SecurityKey()
	}
	parameterSlice = append(parameterSlice, appKey)
	parameterSlice = append(parameterSlice, securityKey)
	parameterSlice = append(parameterSlice, timestamp)
	parameterSlice = append(parameterSlice, random)
	sort.Strings(parameterSlice)
	parameters := strings.Join(parameterSlice, "")
	encryptedParameters := encryptionutils.EncryptWithSha256(parameters)

	sugar.Infof("Ninelaw OpenAPI Security Parameters: [%s]", typeDescription)
	sugar.Infow("[01]", zap.String("appKey", appKey), zap.String("securityKey", securityKey))
	sugar.Infow("[02]", zap.String("timestamp", timestamp))
	sugar.Infow("[03]", zap.String("random", random))
	sugar.Infow("[04]", zap.String("sign", encryptedParameters))
	sugar.Info("======================================================================")
}

// 随机生成"英文字母、数字形式"的字符串
func GenerateRandom(length int) string {
	randomBase := "abcdefghijklmnopqrstuvwxyz0123456789"
	var builder bytes.Buffer
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		random := rand.Intn(len(randomBase))
		char := string([]rune(randomBase)[random])
		// fmt.Printf("随机数: %d, 对应字符: %s\n", random, char)
		builder.WriteString(char)
	}
	return builder.String()
}