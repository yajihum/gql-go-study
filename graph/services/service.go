package services

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yajium/gql-go-study/graph/model"
)

type Services interface {
	UserService
	RepositoryService
	IssueService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error) // user.goにあるGetUserByNamaeを参照できるようにする
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

type RepositoryService interface {
	GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error)
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByID(ctx context.Context, id string) (*model.Issue, error)
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
	ListIssueInRepository(ctx context.Context, repoID string, after *string, before *string, first *int, last *int) (*model.IssueConnection, error)
}

type services struct {
	*userService
	*repoService
	*issueService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

// ファクトリー関数（オブジェクト生成のために使用される特殊な関数）
func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:  &userService{exec: exec},
		repoService:  &repoService{exec: exec},
		issueService: &issueService{exec: exec},
	}
}
