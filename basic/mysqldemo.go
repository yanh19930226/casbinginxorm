package main

import "fmt"

// channel 练习
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

// import (
// 	"fmt"

// 	"github.com/garyburd/redigo/redis"
// 	_ "github.com/go-sql-driver/mysql"
// )

// type Person struct {
// 	UserId   int    `db:"user_id"`
// 	Username string `db:"username"`
// 	Sex      string `db:"sex"`
// 	Email    string `db:"email"`
// }

// type Place struct {
// 	Country string `db:"country"`
// 	City    string `db:"city"`
// 	TelCode int    `db:"telcode"`
// }

// // var Db *sqlx.DB

// // func init() {
// // 	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
// // 	if err != nil {
// // 		fmt.Println("open mysql failed,", err)
// // 		return
// // 	}
// // 	Db = database
// // 	// defer db.Close() // 注意这行代码要写在上面err判断的下面
// // }

// var pool *redis.Pool //创建redis连接池

// func init() {
// 	pool = &redis.Pool{ //实例化一个连接池
// 		MaxIdle: 16, //最初的连接数量
// 		// MaxActive:1000000,    //最大连接数量
// 		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
// 		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
// 		Dial: func() (redis.Conn, error) { //要连接的redis数据库
// 			return redis.Dial(
// 				"tcp",
// 				"121.43.34.54:6379",
// 				redis.DialDatabase(13),
// 				redis.DialPassword("FOXIT@^^#!66"),
// 			)
// 		},
// 	}
// }

// func main() {

// 	c := pool.Get() //从连接池，取一个链接
// 	defer c.Close() //函数运行结束 ，把连接放回连接池

// 	_, err := c.Do("Set", "abc", 200)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	r, err := redis.Int(c.Do("Get", "abc"))
// 	if err != nil {
// 		fmt.Println("get abc faild :", err)
// 		return
// 	}
// 	fmt.Println(r)
// 	pool.Close() //关闭连接池

// 	// r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
// 	// if err != nil {
// 	// 	fmt.Println("exec failed, ", err)
// 	// 	return
// 	// }
// 	// id, err := r.LastInsertId()
// 	// if err != nil {
// 	// 	fmt.Println("exec failed, ", err)
// 	// 	return
// 	// }

// 	// fmt.Println("insert succ:", id)

// 	// var person []Person
// 	// err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
// 	// if err != nil {
// 	// 	fmt.Println("exec failed, ", err)
// 	// 	return
// 	// }

// 	// fmt.Println("select succ:", person)

// 	// res, err := Db.Exec("update person set username=? where user_id=?", "stu0003", 2)
// 	// if err != nil {
// 	// 	fmt.Println("exec failed, ", err)
// 	// 	return
// 	// }
// 	// row, err := res.RowsAffected()
// 	// if err != nil {
// 	// 	fmt.Println("rows failed, ", err)
// 	// }
// 	// fmt.Println("update succ:", row)

// 	// res, err := Db.Exec("delete from person where user_id=?", 2)
// 	// if err != nil {
// 	// 	fmt.Println("exec failed, ", err)
// 	// 	return
// 	// }

// 	// row, err := res.RowsAffected()
// 	// if err != nil {
// 	// 	fmt.Println("rows failed, ", err)
// 	// }

// 	// fmt.Println("delete succ: ", row)

// 	// conn, err := Db.Begin()
// 	// if err != nil {
// 	// 	fmt.Println("begin failed :", err)
// 	// 	return
// 	// }

// 	// r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
// 	// if err != nil {
// 	// 	fmt.Println("exec failed, ", err)
// 	// 	conn.Rollback()
// 	// 	return
// 	// }
// 	// id, err := r.LastInsertId()
// 	// if err != nil {
// 	// 	fmt.Println("exec failed, ", err)
// 	// 	conn.Rollback()
// 	// 	return
// 	// }
// 	// fmt.Println("insert succ:", id)

// 	// r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
// 	// if err != nil {
// 	// 	fmt.Println("exec failed, ", err)
// 	// 	conn.Rollback()
// 	// 	return
// 	// }
// 	// id, err = r.LastInsertId()
// 	// if err != nil {
// 	// 	fmt.Println("exec failed, ", err)
// 	// 	conn.Rollback()
// 	// 	return
// 	// }
// 	// fmt.Println("insert succ:", id)

// 	// conn.Commit()

// 	// c, err := redis.Dial(
// 	// 	"tcp",
// 	// 	"121.43.34.54:6379",
// 	// 	redis.DialDatabase(13),
// 	// 	redis.DialPassword("FOXIT@^^#!66"),
// 	// )
// 	// if err != nil {
// 	// 	fmt.Println("conn redis failed,", err)
// 	// 	return
// 	// }
// 	// fmt.Println("redis conn success")
// 	// defer c.Close()

// 	// _, err = c.Do("HSet", "books", "abc", 100)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// r, err := redis.Int(c.Do("HGet", "books", "abc"))
// 	// if err != nil {
// 	// 	fmt.Println("get abc failed,", err)
// 	// 	return
// 	// }

// 	// fmt.Println(r)

// 	// _, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// r, err := redis.String(c.Do("lpop", "book_list"))
// 	// if err != nil {
// 	// 	fmt.Println("get abc failed,", err)
// 	// 	return
// 	// }

// 	// fmt.Println(r)
// 	// _, err = c.Do("expire", "abc", 10)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// _, err = c.Do("MSet", "abc", 100, "efg", 300)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
// 	// if err != nil {
// 	// 	fmt.Println("get abc failed,", err)
// 	// 	return
// 	// }

// 	// for _, v := range r {
// 	// 	fmt.Println(v)
// 	// }

// 	// _, err = c.Do("Set", "abc", 100)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// r, err := redis.Int(c.Do("Get", "abc"))
// 	// if err != nil {
// 	// 	fmt.Println("get abc failed,", err)
// 	// 	return
// 	// }

// 	// fmt.Println(r)

// }
