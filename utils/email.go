package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/smtp"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	EmailHost        string
	EmailPort        string
	FromEmail        string
	FromEmailSmtpPwd string
)

var (
	// cache for storing verification codes
	// 缓存中的验证代码将在创建后5分钟内有效，且每隔10分钟进行一次���理。
	verificationCodeCache = cache.New(5*time.Minute, 10*time.Minute)
)

// SendVerificationCode sends a verification code to the user's email
func SendVerificationCode(to string) error {
	// Check if a code was sent within the last 60 seconds
	if timestamp, found := verificationCodeCache.Get(to + "_timestamp"); found {
		if time.Since(timestamp.(time.Time)) < 60*time.Second {
			return fmt.Errorf("请在60秒后再试")
		}
	}

	code := generateVerificationCode()

	err := sendVerificationCode(to, code)
	if err != nil {
		return err
	}

	// store the verification code and timestamp in the cache for later verification
	verificationCodeCache.Set(to, code, cache.DefaultExpiration)
	verificationCodeCache.Set(to+"_timestamp", time.Now(), cache.DefaultExpiration)

	return nil
}

// sendVerificationCode 发送验证代码到指定的邮箱。
// 参数 to: 邮件接收人的邮箱地址。
// 参数 code: 需要发送的验证代码。
// 返回值 error: 发送过程中遇到的任何错误。
// sendVerificationCode 发送验证代码到指定的邮箱。
func sendVerificationCode(to string, code string) error {
	server := fmt.Sprintf("%s:%s", EmailHost, EmailPort)

	header := make(map[string]string)

	header["From"] = "STUOJ" + "<" + FromEmail + ">"
	header["To"] = to
	header["Subject"] = "STUOJ 邮箱验证"
	header["Content-Type"] = "text/html;chartset=UTF-8"

	body := fmt.Sprintf("您的验证码是: %s", code)

	message := ""

	for k, v := range header {
		message += fmt.Sprintf("%s:%s\r\n", k, v)
	}

	message += "\r\n" + body

	auth := smtp.PlainAuth(
		"",
		FromEmail,
		FromEmailSmtpPwd,
		EmailHost,
	)

	err := SendMailUsingTLS(
		server,
		auth,
		FromEmail,
		[]string{to},
		[]byte(message),
	)
	return err
}

// 随机生成一个6位数的验证码。
func generateVerificationCode() string {
	// 使用 crypto/rand 生成一个 0 到 999999 之间的随机数
	max := big.NewInt(999999)
	num, err := rand.Int(rand.Reader, max.Add(max, big.NewInt(1)))
	if err != nil {
		// 处理随机数生成失败的情况
		fmt.Println("Error generating random number:", err)
		return "000000" // 返回默认值或采取其他措施
	}

	// 将大整数转换为字符串并格式化为 6 位数
	code := fmt.Sprintf("%06d", num.Int64())
	return code
}

// VerifyVerificationCode verifies the verification code sent to the user
func VerifyVerificationCode(email string, code string) bool {
	// retrieve the verification code from the cache
	cachedCode, found := verificationCodeCache.Get(email)
	// 如果没有找到验证码或者验证码过期，返回false
	if !found {
		return false
	}

	// compare the cached code with the provided code
	if cachedCode != code {
		return false
	}

	// 验证成功后，从缓存中删除验证码
	verificationCodeCache.Delete(email)
	// verificationCodeCache.Delete(email + "_timestamp")

	return true
}

// return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Panicln("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

// 参考net/smtp的func SendMail()
// 使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
// len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = c.Quit()
	if err != nil && !strings.Contains(err.Error(), "250") {
		return err
	}
	return nil
}
