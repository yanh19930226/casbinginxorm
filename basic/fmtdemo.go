package main

// import (
// 	"fmt"
// 	"reflect"
// 	"time"
// )

// func main() {
// 	// fmt.Print("在终端打印该信息。")
// 	// name := "枯藤"
// 	// fmt.Printf("我是：%s\n", name)
// 	// fmt.Println("在终端打印单独一行显示")

// 	//
// 	// // 向标准输出写入内容,os.Stdout往控制台输出标准内容
// 	// fmt.Fprintln(os.Stdout, "向标准输出写入内容")
// 	// fileObj, err := os.OpenFile("./fmtbasic.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	// if err != nil {
// 	// 	fmt.Println("打开文件出错，err:", err)
// 	// 	return
// 	// }
// 	// name := "枯藤"
// 	// // 向打开的文件句柄中写入内容
// 	// fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

// 	//Sprint系列函数会把传入的数据生成并返回一个字符串
// 	// s1 := fmt.Sprint("枯藤")
// 	// name := "枯藤"
// 	// age := 18
// 	// s2 := fmt.Sprintf("name:%s,age:%d", name, age)
// 	// s3 := fmt.Sprintln("枯藤")
// 	// fmt.Println(s1, s2, s3)

// 	// err := fmt.Errorf("这是一个错误")

// 	// fmt.Println(err)

// 	// fmt.Printf("%v\n", 100)
// 	// fmt.Printf("%v\n", false)
// 	// o := struct{ name string }{"枯藤"}
// 	// fmt.Printf("%v\n", o)
// 	// fmt.Printf("%#v\n", o)
// 	// fmt.Printf("%T\n", o)
// 	// fmt.Printf("100%%\n")

// 	// n := 65
// 	// fmt.Printf("%b\n", n)
// 	// fmt.Printf("%c\n", n)
// 	// fmt.Printf("%d\n", n)
// 	// fmt.Printf("%o\n", n)
// 	// fmt.Printf("%x\n", n)
// 	// fmt.Printf("%X\n", n)

// 	// s := "枯藤"
// 	// fmt.Printf("%s\n", s)
// 	// fmt.Printf("%q\n", s)
// 	// fmt.Printf("%x\n", s)
// 	// fmt.Printf("%X\n", s)

// 	// n := 88.88
// 	// fmt.Printf("%f\n", n)
// 	// fmt.Printf("%9f\n", n)
// 	// fmt.Printf("%.2f\n", n)
// 	// fmt.Printf("%9.2f\n", n)
// 	// fmt.Printf("%9.f\n", n)

// 	// var (
// 	// 	name    string
// 	// 	age     int
// 	// 	married bool
// 	// )
// 	// fmt.Scan(&name, &age, &married)
// 	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

// 	// timeDemo()

// 	// timestampDemo()

// 	// now := time.Now()
// 	// later := now.Add(time.Hour) // 当前时间加1小时后的时间
// 	// fmt.Println(later)

// 	// // tickDemo()

// 	// formatDemo()

// 	// now := time.Now()
// 	// fmt.Println(now)
// 	// // 加载时区
// 	// loc, err := time.LoadLocation("Asia/Shanghai")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// // 按照指定时区和指定格式解析字符串时间
// 	// timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(timeObj)
// 	// fmt.Println(timeObj.Sub(now))

// 	// log.Println("这是一条很普通的日志。")
// 	// v := "很普通的"
// 	// log.Printf("这是一条%s日志。\n", v)
// 	// log.Fatalln("这是一条会触发fatal的日志。")
// 	// log.Panicln("这是一条会触发panic的日志。")

// 	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
// 	// log.Println("这是一条很普通的日志。")

// 	// log.SetPrefix("[pprof]")
// 	// log.Println("这是一条很普通的日志。")

// 	// logFile, err := os.OpenFile("./logtest.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	// if err != nil {
// 	// 	fmt.Println("open log file failed, err:", err)
// 	// 	return
// 	// }
// 	// log.SetOutput(logFile)
// 	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
// 	// log.Println("这是一条很普通的日志。")
// 	// log.SetPrefix("[小王子]")
// 	// log.Println("这是一条很普通的日志。")

// 	// logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
// 	// logger.Println("这是自定义的logger记录的日志。")

// 	// Atoi()函数用于将字符串类型的整数转换为int类型，函数签名如下
// 	// s1 := "100"
// 	// i1, err := strconv.Atoi(s1)
// 	// if err != nil {
// 	// 	fmt.Println("can't convert to int")
// 	// } else {
// 	// 	fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
// 	// }

// 	// i2 := 200
// 	// s2 := strconv.Itoa(i2)
// 	// fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"

// 	// b, err := strconv.ParseBool("true")
// 	// f, err := strconv.ParseFloat("3.1415", 64)
// 	// i, err := strconv.ParseInt("-2", 10, 64)
// 	// u, err := strconv.ParseUint("2", 10, 64)

// 	// resp, err := http.Get("https://www.baidu.com/")
// 	// if err != nil {
// 	// 	fmt.Println("get failed, err:", err)
// 	// 	return
// 	// }
// 	// defer resp.Body.Close()
// 	// body, err := ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	fmt.Println("read from resp.Body failed,err:", err)
// 	// 	return
// 	// }
// 	// fmt.Print(string(body))

// 	// apiUrl := "http://127.0.0.1:9090/get"
// 	// // URL param
// 	// data := url.Values{}
// 	// data.Set("name", "枯藤")
// 	// data.Set("age", "18")
// 	// u, err := url.ParseRequestURI(apiUrl)
// 	// if err != nil {
// 	// 	fmt.Printf("parse url requestUrl failed,err:%v\n", err)
// 	// }
// 	// u.RawQuery = data.Encode() // URL encode
// 	// fmt.Println(u.String())
// 	// resp, err := http.Get(u.String())
// 	// if err != nil {
// 	// 	fmt.Println("post failed, err:%v\n", err)
// 	// 	return
// 	// }
// 	// defer resp.Body.Close()
// 	// b, err := ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	fmt.Println("get resp failed,err:%v\n", err)
// 	// 	return
// 	// }
// 	// fmt.Println(string(b))

// 	// p := Person{"5lmh.com", "女"}
// 	// // 编码json
// 	// b, err := json.Marshal(p)
// 	// if err != nil {
// 	// 	fmt.Println("json err ", err)
// 	// }
// 	// fmt.Println(string(b))

// 	// // 格式化输出
// 	// b, err = json.MarshalIndent(p, "", "     ")
// 	// if err != nil {
// 	// 	fmt.Println("json err ", err)
// 	// }
// 	// fmt.Println(string(b))

// 	// // 假数据
// 	// b := []byte(`{"age":"18","name":"5lmh.com","marry":false}`)
// 	// var p Person
// 	// err := json.Unmarshal(b, &p)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Println(p)

// 	// 假数据
// 	// int和float64都当float64
// 	// b := []byte(`{"age":1.3,"name":"5lmh.com","marry":false}`)

// 	// // 声明接口
// 	// var i interface{}
// 	// err := json.Unmarshal(b, &i)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// // 自动转到map
// 	// fmt.Println(i)
// 	// // 可以判断类型
// 	// m := i.(map[string]interface{})
// 	// for k, v := range m {
// 	// 	switch vv := v.(type) {
// 	// 	case float64:
// 	// 		fmt.Println(k, "是float64类型", vv)
// 	// 	case string:
// 	// 		fmt.Println(k, "是string类型", vv)
// 	// 	default:
// 	// 		fmt.Println("其他")
// 	// 	}
// 	// }

// 	// // 只读方式打开当前目录下的main.go文件
// 	// file, err := os.Open("./main.go")
// 	// if err != nil {
// 	// 	fmt.Println("open file failed!, err:", err)
// 	// 	return
// 	// }

// 	// fmt.Println("open file success")

// 	// // 关闭文件
// 	// file.Close()

// 	// // 新建文件
// 	// file, err := os.Create("./test.txt")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// defer file.Close()
// 	// for i := 0; i < 5; i++ {
// 	// 	file.WriteString("ab\n")
// 	// 	file.Write([]byte("cd\n"))
// 	// }

// 	// // 打开文件
// 	// file, err := os.Open("./test.txt")
// 	// if err != nil {
// 	// 	fmt.Println("open file err :", err)
// 	// 	return
// 	// }
// 	// defer file.Close()
// 	// // 定义接收文件读取的字节数组
// 	// var buf [128]byte
// 	// var content []byte
// 	// for {
// 	// 	n, err := file.Read(buf[:])
// 	// 	if err == io.EOF {
// 	// 		// 读取结束
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		fmt.Println("read file err ", err)
// 	// 		return
// 	// 	}
// 	// 	content = append(content, buf[:n]...)
// 	// }
// 	// fmt.Println(string(content))

// 	// // 打开源文件
// 	// srcFile, err := os.Open("./test.txt")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// // 创建新文件
// 	// dstFile, err2 := os.Create("./abc2.txt")
// 	// if err2 != nil {
// 	// 	fmt.Println(err2)
// 	// 	return
// 	// }
// 	// // 缓冲读取
// 	// buf := make([]byte, 1024)
// 	// for {
// 	// 	// 从源文件读数据
// 	// 	n, err := srcFile.Read(buf)
// 	// 	if err == io.EOF {
// 	// 		fmt.Println("读取完毕")
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 		break
// 	// 	}

// 	// 	fmt.Println(n)
// 	// 	//写出去
// 	// 	dstFile.Write(buf[:n])
// 	// }
// 	// srcFile.Close()
// 	// dstFile.Close()

// 	// region Main

// 	// var x float64 = 3.4
// 	// // reflect_type(x)

// 	// // reflect_value(x)

// 	// // 反射认为下面是指针类型，不是float类型
// 	// reflect_set_value(&x)
// 	// fmt.Println("main:", x)

// 	// endregion

// 	// u := User{1, "zs", 20}
// 	// Poni(u)

// 	// m := Boy{User{1, "zs", 20}, "bj"}
// 	// t := reflect.TypeOf(m)
// 	// fmt.Println(t)
// 	// // Anonymous：匿名
// 	// fmt.Printf("%#v\n", t.Field(0))
// 	// // 值信息
// 	// fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))

// 	// u := User{1, "5lmh.com", 20}
// 	// SetValue(&u)
// 	// fmt.Println(u)

// 	// u := User{1, "5lmh.com", 20}
// 	// v := reflect.ValueOf(u)
// 	// // 获取方法
// 	// m := v.MethodByName("Hello")
// 	// // 构建一些参数
// 	// args := []reflect.Value{reflect.ValueOf("6666")}
// 	// // 没参数的情况下：var args2 []reflect.Value
// 	// // 调用方法，需要传入方法的参数
// 	// m.Call(args)

// 	var s Student
// 	v := reflect.ValueOf(&s)
// 	// 类型
// 	t := v.Type()
// 	// 获取字段
// 	f := t.Elem().Field(0)
// 	fmt.Println(f.Tag.Get("json"))
// 	fmt.Println(f.Tag.Get("db"))

// }

// type Student struct {
// 	Name string `json:"name1" db:"name2"`
// }

// func (u User) Hellos(name string) {
// 	fmt.Println("Hello：", name)
// }

// // 修改结构体值
// func SetValue(o interface{}) {
// 	v := reflect.ValueOf(o)
// 	// 获取指针指向的元素
// 	v = v.Elem()
// 	// 取字段
// 	f := v.FieldByName("Name")
// 	if f.Kind() == reflect.String {
// 		f.SetString("kuteng")
// 	}
// }

// // 反射修改值
// func reflect_set_value(a interface{}) {
// 	v := reflect.ValueOf(a)
// 	k := v.Kind()
// 	switch k {
// 	case reflect.Float64:
// 		// 反射修改值
// 		v.SetFloat(6.9)
// 		fmt.Println("a is ", v.Float())
// 	case reflect.Ptr:
// 		// Elem()获取地址指向的值
// 		v.Elem().SetFloat(7.9)
// 		fmt.Println("case:", v.Elem().Float())
// 		// 地址
// 		fmt.Println(v.Pointer())
// 	}
// }

// func reflect_value(a interface{}) {
// 	v := reflect.ValueOf(a)
// 	fmt.Println(v)
// 	k := v.Kind()
// 	fmt.Println(k)
// 	switch k {
// 	case reflect.Float64:
// 		fmt.Println("a是：", v.Float())
// 	}
// }

// func reflect_type(a interface{}) {
// 	t := reflect.TypeOf(a)
// 	fmt.Println("类型是：", t)
// 	// kind()可以获取具体类型
// 	k := t.Kind()
// 	fmt.Println(k)
// 	switch k {
// 	case reflect.Float64:
// 		fmt.Printf("a is float64\n")
// 	case reflect.String:
// 		fmt.Println("string")
// 	}
// }

// // 传入interface{}
// func Poni(o interface{}) {
// 	t := reflect.TypeOf(o)
// 	fmt.Println("类型：", t)
// 	fmt.Println("字符串类型：", t.Name())
// 	// 获取值
// 	v := reflect.ValueOf(o)
// 	fmt.Println(v)
// 	// 可以获取所有属性
// 	// 获取结构体字段个数：t.NumField()
// 	for i := 0; i < t.NumField(); i++ {
// 		// 取每个字段
// 		f := t.Field(i)
// 		fmt.Printf("%s : %v", f.Name, f.Type)
// 		// 获取字段的值信息
// 		// Interface()：获取字段对应的值
// 		val := v.Field(i).Interface()
// 		fmt.Println("val :", val)
// 	}
// 	fmt.Println("=================方法====================")
// 	for i := 0; i < t.NumMethod(); i++ {
// 		m := t.Method(i)
// 		fmt.Println(m.Name)
// 		fmt.Println(m.Type)
// 	}

// }

// // 匿名字段
// type Boy struct {
// 	User
// 	Addr string
// }

// // 绑方法
// func (u User) Hello() {
// 	fmt.Println("Hello")
// }

// // 定义结构体
// type User struct {
// 	Id   int
// 	Name string
// 	Age  int
// }

// type Person struct {
// 	Age       int    `json:"age,string"`
// 	Name      string `json:"name"`
// 	Niubility bool   `json:"niubility"`
// }

// // type Person struct {
// // 	Name  string
// // 	Hobby string
// // }

// func formatDemo() {
// 	now := time.Now()
// 	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
// 	// 24小时制
// 	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
// 	// 12小时制
// 	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
// 	fmt.Println(now.Format("2006/01/02 15:04"))
// 	fmt.Println(now.Format("15:04 2006/01/02"))
// 	fmt.Println(now.Format("2006/01/02"))
// }

// func tickDemo() {
// 	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
// 	for i := range ticker {
// 		fmt.Println(i) //每秒都会执行的任务
// 	}
// }

// func timeDemo() {
// 	now := time.Now() //获取当前时间
// 	fmt.Printf("current time:%v\n", now)

// 	year := now.Year()     //年
// 	month := now.Month()   //月
// 	day := now.Day()       //日
// 	hour := now.Hour()     //小时
// 	minute := now.Minute() //分钟
// 	second := now.Second() //秒
// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
// }

// // 时间戳
// func timestampDemo() {
// 	now := time.Now()            //获取当前时间
// 	timestamp1 := now.Unix()     //时间戳
// 	timestamp2 := now.UnixNano() //纳秒时间戳
// 	fmt.Printf("current timestamp1:%v\n", timestamp1)
// 	fmt.Printf("current timestamp2:%v\n", timestamp2)
// }

// // time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位
