package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/DanielKirkwood/youchooseserver/ent"
	"github.com/DanielKirkwood/youchooseserver/ent/user"
	"github.com/DanielKirkwood/youchooseserver/internal/domain/auth"
	"github.com/DanielKirkwood/youchooseserver/internal/util/email"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

type Auth interface {
	Create(ctx context.Context, req *auth.LoginRequest) error
	Verify(ctx context.Context, req *auth.VerifyRequest) (string, error)
}

type authRepository struct {
	ent       *ent.Client
	email     *email.Client
	tokenAuth *jwtauth.JWTAuth
}

func New(ent *ent.Client, email *email.Client, secret string) *authRepository {
	return &authRepository{
		ent:       ent,
		email:     email,
		tokenAuth: jwtauth.New("HS256", []byte(secret), nil),
	}
}

func (r *authRepository) Create(ctx context.Context, req *auth.LoginRequest) error {
	var (
		u   *ent.User
		err error
	)

	u, err = r.ent.User.Query().Where(user.Email(req.Email)).Only(ctx)
	if err != nil && err != err.(*ent.NotFoundError) {
		return err
	}

	if u == nil {
		u, err = r.ent.User.
			Create().
			SetEmail(req.Email).
			Save(ctx)

		if err != nil {
			return err
		}
	}

	// generate uuid as OTP and store it alongside user details in db with current time
	otp, err := generateUUID()
	if err != nil {
		return err
	}

	u, err = u.Update().SetOtp(otp).SetOtpExpiresAt(time.Now().Add(time.Minute * 15)).Save(ctx)
	if err != nil {
		return err
	}

	link := "/auth/passwordless/verify_redirect" +
		"?otp=" + otp + "&userid=" + strconv.Itoa(u.ID)

	msg := "You (or someone who knows your email address) wants " +
		"to sign in to You Choose.\n\n" +
		"Your PIN is " + otp + " - or use the following link: " +
		link + "\n\n" +
		"(If you did not request or were not expecting this email, " +
		"you can safely ignore it.)"

	if err := r.email.SendMail(r.email.CreateMessage(u.Email, msg)); err != nil {
		return err
	}

	return nil
}

func (r *authRepository) Verify(ctx context.Context, req *auth.VerifyRequest) (string, error) {
	u, err := r.ent.User.Get(ctx, req.UserID)
	if err != nil {
		return "", err
	}

	if time.Now().After(*u.OtpExpiresAt) {
		return "", fmt.Errorf("time expired. Please request a new password")
	}

	_, tokenString, err := r.tokenAuth.Encode(map[string]interface{}{
		"userid": u.ID,
		"email":  u.Email,
	})
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// GenerateUUID returns a new v4 UUID.
func generateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
