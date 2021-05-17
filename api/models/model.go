package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	// 导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	
	"ot/config"
)

// DB db handler
var DB *sqlx.DB

// baseModel 继承使用，减少反复写此段代码
type baseModel struct { 
	ID        int        `db:"id"      json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

func judgeType(v string) string {
	// 判断类型是否是数字，数字的类型和字符串在sql里要区别
	index := strings.Contains(v, "int")
	f := strings.Contains(v, "float")
	var str string
	if index {
		str = "%v"
	} else {
		if f {
			str = "%v"
		} else {
			str = "'%v'"
		}

	}
	return str

}

// ReflectTag 构建  tag=值的字符串， 只能两层结构体
func ReflectTag(s interface{}, arg ...string) (resArray []string, err error) {
	var res string

	st := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		field := st.Field(i)
		// 指定结构体的key直接不添加到切片中，这样的目的为了去掉时间戳的问题
		if arg != nil {
			flag := false
			for k := 0; k < len(arg); k++ {
				if field.Name == arg[k] {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
		}
		// id和时间戳，时间戳是不转化的，自己在代码里面进行构建
		if field.Name == "baseModel" {
			continue
		}

		if !v.Field(i).IsZero() {
			// 如果结构体有下一层，直接开始下一层的判断，构造
			if v.Field(i).Type().Kind() == reflect.Struct {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					if structField.Field(j).Tag.Get("db") == "" {
						continue
					}
					str := "%s=" + judgeType(fmt.Sprintf("%v", v.Field(i).Field(j).Type()))
					res = fmt.Sprintf(str, structField.Field(j).Tag.Get("db"), v.Field(i).Field(j))
					resArray = append(resArray, res)
				}
				continue
			} else {
				str := "%s=" + judgeType(fmt.Sprintf("%v", v.Field(i).Type()))
				res = fmt.Sprintf(str, field.Tag.Get("db"), v.Field(i))
				resArray = append(resArray, res)
			}

		}
	}
	if len(resArray) == 0 {
		return resArray, errors.New("resArray is empty")
	}
	return resArray, nil
}

// InitMysql 初始化数据库
func InitMysql() (err error) { 
	DB, err = sqlx.Open("mysql", config.MysqlConnect)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

// Close 关闭数据库
func Close() {
	err := DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}