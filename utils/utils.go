package utils

//公共方法
import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 简化response代码
type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  GetMsg(errCode),
		"data": data,
	})

	return
}

// 输入验证的错误处理
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Println(err.Key, err.Message)
	}

	return
}

// 结构体转映射
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// md5加密
func Md5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

// 返回SHA256加密
func SHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

// 生成随机字符串(大写字母)
func RandomString(len int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < len; {
		if string(RandomInt(65, 90)) != temp {
			temp = string(RandomInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

// 生成随机数字
func RandomInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
