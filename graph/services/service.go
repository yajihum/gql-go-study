package services

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yajium/gql-go-study/graph/model"
)

type Services interface {
	UserService
	RepositoryService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error) // user.goにあるGetUserByNamaeを参照できるようにする
}

type RepositoryService interface {
	GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error)
}

type services struct {
	*userService
	*repositoryService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

// ファクトリー関数（オブジェクト生成のために使用される特殊な関数）
func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
	}
}
