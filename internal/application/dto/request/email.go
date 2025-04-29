package request

type SendEmailVerificationCode struct {
	Email string `json:"email"`
}
