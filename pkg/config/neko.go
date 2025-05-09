package config

type NekoConf struct {
	Host  string `yaml:"host" json:"host"`
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

func (n *NekoConf) Default() {
	n.Host = "stuoj-neko"
	n.Port = "14515"
	n.Token = ""
}
