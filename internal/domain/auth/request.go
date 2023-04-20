package auth

type LoginRequest struct {
	Email string `json:"email"`
}

type VerifyRequest struct {
	UserID int    `json:"userid"`
	OTP    string `json:"otp"`
}
