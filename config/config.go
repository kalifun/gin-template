package config

type ConfigInfo struct {
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	System System `json:"system" yaml:"system"`
	Log    Log    `json:"log" yaml:"log"`
}

type Mysql struct {
	UserName     string `json:"user_name" yaml:"username"`
	PassWord     string `json:"pass_word" yaml:"password"`
	Path         string `json:"path" yaml:"path"`
	Port         string `json:"port" yaml:"port"`
	DbName       string `json:"db_name" yaml:"db-name"`
	Config       string `json:"config" yaml:"config"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `json:"max_open_conns" yaml:"max-open-conns"`
}

type System struct {
	Port string `json:"port" yaml:"port"`
}

type Log struct {
	Prefix  string `json:"prefix" yaml:"prefix"`
	LogFile bool   `json:"log_file" yaml:"log-file"`
	Stdout  string `json:"stdout" yaml:"stdout"`
	File    string `json:"file" yaml:"file"`
}
