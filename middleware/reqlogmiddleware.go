package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	//日志地址
	logFilePath = "./"
	//日志文件名
	logFileName = "system.log"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func getRequestBody(ctx *gin.Context) interface{} {
	switch ctx.Request.Method {
	case http.MethodGet:
		return ctx.Request.URL.Query()
	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		fallthrough
	case http.MethodPatch:
		var bodyBytes []byte
		bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			return nil
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		return string(bodyBytes)
	}
	return nil
}

// Write 读取响应数据
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ReqLog() gin.HandlerFunc {
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	logger := logrus.New()
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置输出
	logger.Out = src
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()

		// 初始化bodyLogWriter
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// 获取请求信息
		requestBody := getRequestBody(c)

		c.Next()

		// 响应内容
		responseBody := blw.body.String()

		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"client_ip":    clientIP,
			"header":       c.Request.Header,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
			"req_body":     requestBody,
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"end_time":     endTime,
			"res_body":     responseBody,
		}).Info()
	}
}
