package basic

import (
	"ai-gallery/pkg/jwt"
	"context"

	"ai-gallery/ent"
	"ai-gallery/service/internal/dao/account"
	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/model"

	"golang.org/x/crypto/bcrypt"
)

func GetAuthorIDsByRole(ctx context.Context, authorIDs []int) ([]int, error) {
	jwtInfo, err := ContextElement(ctx)
	if err != nil {
		return nil, err
	}
	// 普通用户只能看自己的数据.
	if jwtInfo.Role == model.USER {
		return []int{jwtInfo.ID}, nil
	}
	return authorIDs, nil
}

func ContextElement(ctx context.Context) (*jwt.Claims, error) {
	if info, ok := ctx.Value(model.UserToken).(*jwt.Claims); ok {
		return info, nil
	}
	return nil, errors.ErrAccountTokenNotFound(ctx)
}

func ValidateAccount(ctx context.Context, username, password string) (*ent.Account, error) {
	user, err := account.GetByUsername(ctx, username)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrAccountNotFound(ctx)
		}
		return nil, err
	}
	if user.Status == model.DISABLED {
		return nil, errors.ErrAccountDisabled(ctx)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+model.PwdSalt)); err != nil {
		return nil, errors.ErrAccountPwd(ctx)
	}
	return user, nil
}

func CheckAccountStatus(ctx context.Context, id int) error {
	acc, err := account.GetByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.ErrAccountNotFound(ctx)
		}
		return err
	}
	if acc.Status == model.DISABLED {
		return errors.ErrAccountDisabled(ctx)
	}
	return nil
}

func CheckPermission(ctx context.Context, claims *jwt.Claims) error {
	if claims.Role != model.ADMIN {
		return errors.ErrAccountPermission(ctx)
	}
	return nil
}
