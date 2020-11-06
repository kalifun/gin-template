package config

type ConfigInfo struct {
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	System System `json:"system" yaml:"system"`
	Log    Log    `json:"logs" yaml:"logs"`
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
	Port uint64 `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

type Log struct {
	Prefix  string `json:"prefix" yaml:"prefix"`
	LogFile bool   `json:"log_file" yaml:"logs-file"`
	Path    string `json:"path" yaml:"path"`
	Stdout  string `json:"stdout" yaml:"stdout"`
	File    string `json:"file" yaml:"file"`
}
