package ninelaw

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang-utils/encryptionutils"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 生成九品验签参数
func GenerateNinelawOpenAPISecurityParameters() {
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

	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	random := generateRandom(6)

	for _, n := range NineLawOpenApiKeysSlice {
		parameterSlice := make([]string, 4)
		parameterSlice = append(parameterSlice, n.AppKey())
		parameterSlice = append(parameterSlice, n.SecurityKey())
		parameterSlice = append(parameterSlice, timestamp)
		parameterSlice = append(parameterSlice, random)
		sort.Strings(parameterSlice)
		parameters := strings.Join(parameterSlice, "")
		encryptedParameters := encryptionutils.EncryptWithSha256(parameters)

		sugar.Infof("Ninelaw OpenAPI Security Parameters: Environment: [%s], Description: [%s]", n.Environment(), n.TypeDescription())
		sugar.Infow("[01]", zap.String("appKey", n.AppKey()))
		sugar.Infow("[02]", zap.String("securityKey", n.SecurityKey()))
		sugar.Infow("[03]", zap.String("timestamp", timestamp))
		sugar.Infow("[04]", zap.String("random", random))
		sugar.Infow("[05]", zap.String("sign", encryptedParameters))
		sugar.Info("======================================================================")
	}
}

// 随机生成"英文字母、数字形式"的字符串
func generateRandom(length int) string {
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
