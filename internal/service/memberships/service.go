package memberships

import (
	"context"

	"github.com/xdenistwn/xg-forum.git/internal/configs"
	"github.com/xdenistwn/xg-forum.git/internal/model/memberships"
)

type membershipsRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	cfg            *configs.Config
	membershipRepo membershipsRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipsRepository) *service {
	return &service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
