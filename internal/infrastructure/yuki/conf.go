package yuki

type YukiConf struct {
	Host  string `yaml:"host" json:"host"`
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

var (
	config YukiConf
	preUrl string
)
