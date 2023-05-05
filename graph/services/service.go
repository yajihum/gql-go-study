package services

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yajium/gql-go-study/graph/model"
)

type Services interface {
	UserService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type services struct {
	*userService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error) // user.goにあるGetUserByNamaeを参照できるようにする
}

// ファクトリー関数（オブジェクト生成のために使用される特殊な関数）
func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
	}
}
