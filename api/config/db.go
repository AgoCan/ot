package config

type Db struct {
	Mysql struct {
		DbName   string
		Password string
		Username string
		Port     string
		Host     string
	}
}

func (d *Db) Dsn() string {
	return Conf.Db.Mysql.Username + ":" + Conf.Db.Mysql.Password + "@tcp(" +
		Conf.Db.Mysql.Host + ":" + Conf.Db.Mysql.Port + ")/" + Conf.Db.Mysql.DbName +
		"?charset=utf8mb4&parseTime=True&loc=Local"
}
