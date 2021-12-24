package config

type Mysql struct {
	Username string `yaml:"username"` // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	URL      string `yaml:"url"`      // 数据库ip地址
	Port     int    `yaml:"port"`     // 数据库端口号
	DBName   string `yaml:"dbName"`   // 数据库名称
}
