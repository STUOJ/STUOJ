package config

type YukiConf struct {
	Host  string `yaml:"host" json:"host"`
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

func (y *YukiConf) Default() {
	y.Host = "stuoj-yuki"
	y.Port = "7415"
	y.Token = ""
}
