package biz

import (
	"context"

	v1 "cube-core/api/realworld/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// GreeterRepo is a Greater repo.
type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

// GreeterUsecase is a Greeter usecase.
type SocialUseCase struct {
	ar  ArticleRepo
	cr  CommentRepo
	tr  TagRepo
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewSocialUsecase(ar ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *SocialUseCase {
	return &SocialUseCase{ar: ar, cr: cr, tr: tr, log: log.NewHelper(logger)}
}

func (uc *SocialUseCase) CreateArticle(ctx context.Context) error {
	return nil
}
