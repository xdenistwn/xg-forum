package memberships

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/xdenistwn/xg-forum.git/internal/model/memberships"
	"github.com/xdenistwn/xg-forum.git/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	// check user
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")

	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("Email not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return "", errors.New("Invalid email or password.")
	}

	// generate token jwt
	jwt, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		return "", err
	}

	return jwt, nil
}
