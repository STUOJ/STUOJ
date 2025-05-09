package config

type EmailConf struct {
	Host    string `yaml:"host" json:"host"`
	Port    string `yaml:"port" json:"port"`
	Email   string `yaml:"email" json:"email"`
	SmtpPwd string `yaml:"smtp_pwd" json:"smtp_pwd"`
}

func (e *EmailConf) Default() {
	e.Host = "smtp.qq.com"
	e.Port = "465"
	e.Email = ""
	e.SmtpPwd = ""
}
