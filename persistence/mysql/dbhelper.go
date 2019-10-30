package main

/**
 * github.com/go-sql-driver/mysql 需要下载，点击golang开发工具会自动下载
 */
import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

/***
 *数据库常量
 */
const (
	USERNAME = "root"          //用户名
	PASSWORD = "root"          //密码
	NETWORK  = "tcp"           //网络协议
	SERVER   = "192.168.1.231" //数据库IP
	DATABASE = "golang"        //数据库名称
)

func DataBaseConfig() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = USERNAME
	cfg.Passwd = PASSWORD
	cfg.Net = NETWORK
	cfg.Addr = SERVER
	cfg.DBName = DATABASE
	dsn := cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

var pool = DataBaseConfig()

/***
 * InsertUser 插入用户
 */
func InsertUser(userName string, passWord string) (int64, error) {
	res, err := pool.Exec("insert into `t_user` (`user_name` ,`pass_word`) values (?,?)", userName, passWord)
	if err != nil {
		return 0, err
	}
	lastInsertID,errx := res.LastInsertId()  //插入数据的主键id
	fmt.Println("LastInsertID:",lastInsertID)
	return lastInsertID,errx
}

// DeleteUser 删除用户
func DeleteUser(id int64) error {
	_, err := pool.Exec("delete from `t_user` where `id` = ?", id)
	if err != nil {
		return err
	}
	pool.Close()
	return nil
}

/**
 * UpdateUser 更新用户
 */
func UpdateUser(id int64, userName string) error {
	_, err := pool.Exec("update `t_user` set `user_name` = ? where `id` = ?", userName, id)
	if err != nil {
		return err
	}

	return nil
}

/***
 * 根据Id去查询一条数据
 */
func queryOne(id int) *User {
	user := new(User)
	row := pool.QueryRow("select * from t_user where id=?", id)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.ID, &user.UserName, &user.PassWord); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}
	return user
}

/***
 * 查询所有用户信息
 */
func SelectAllUser() ([]User) {
	//执行查询语句
	rows, err := pool.Query("SELECT * from t_user")
	if err != nil {
		fmt.Println("查询出错了，错误信息为:", err)
	}
	var users []User
	//循环读取结果
	for rows.Next() {
		var user User
		//将每一行的结果都赋值到一个user对象中
		err := rows.Scan(&user.ID, &user.UserName, &user.PassWord)
		if err != nil {
			fmt.Println("查询数据库行出错，错误信息为:", err)
		}
		//将user追加到users的这个数组中
		users = append(users, user)
	}
	return users
}

/**
 * 数据结构模型
 */
type User struct {
	ID int64 `db:"id"`
	//由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	UserName string `db:"user_name"`
	PassWord string `db:"pass_word"`
}

func main() {
	//1、添加
	i, e := InsertUser("小梁", "liangjl")
	log.Fatal("返回的信息", e)
	log.Fatal("主键Id", i)

	//2、删除
	//DeleteUser(1)

	//3、修改
	//UpdateUser(2,"liangjl")

	//4、根据id查询
	//user := queryOne(1)
	//fmt.Println("用户名称：", user.UserName)

	//5、查询所有用户信息
	users := SelectAllUser()

	//方法一：循环打印用户信息
	for i := 0; i < len(users); i++ {
		fmt.Println("主键Id", users[i].ID, ",用户名:", users[i].UserName, ",密码:", users[i].PassWord)
	}

	fmt.Println("=============================")

	//方法二：循环打印用户信息
	for _, userVo := range users {
		fmt.Println("主键Id", userVo.ID, ",用户名:", userVo.UserName, ",密码:", userVo.PassWord)
	}

}
