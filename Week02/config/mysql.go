package config

type Mysql struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DBName string `yaml:"db-name"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	MaxIdleConn string `yaml:"max-idle-conn"`
	MaxOpenConn string `yaml:"max-open-conn"`
}
