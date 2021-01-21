package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

const (
	output_dir = "./logs/"
	out_path   = "foot.log"
	err_path   = "foot.err"
)

func init() {
	_, err := os.Stat(output_dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(output_dir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			}
		}
	}

	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	conf := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "caller",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.999999999"))
		},
		EncodeDuration: zapcore.NanosDurationEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(conf)
	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoHook_1 := os.Stdout
	infoHook_2 := getWriter(out_path)
	errorHook := getWriter(err_path)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoHook_1), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoHook_2), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorHook), warnLevel),
	)

	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	Log = logger.Sugar()
	defer logger.Sync()
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		output_dir+filename+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Minute*3),
		// rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithRotationTime(time.Minute),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func main() {
	defer func() {
		if recover() != nil {
			i := 0
			funcName, file, line, ok := runtime.Caller(i)
			for ok {
				Log.Panic("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
				i++
				funcName, file, line, ok = runtime.Caller(i)
			}
		}
	}()
	for {
		time.Sleep(1 * time.Second)
		Log.Info("日志1")
		Log.Info("日志2")
		Log.Error("错误的日志")
		Log.Info("日志3")
		panic(111)
	}
	return
}
