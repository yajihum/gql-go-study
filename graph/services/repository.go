package services

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/yajium/gql-go-study/graph/db"
	"github.com/yajium/gql-go-study/graph/model"
)

// private
type repoService struct {
	exec boil.ContextExecutor // クエリの実行に使用されるデータベース接続を提供するインターフェース
}

// private
// db.User型からmodel.User型に変換する
func convertRepository(repo *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        repo.ID,
		Name:      repo.Name,
		Owner:     &model.User{ID: repo.Owner},
		CreatedAt: repo.CreatedAt,
	}
}

func (r *repoService) GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error) {
	repo, err := db.FindRepository(ctx, r.exec, id,
		db.RepositoryColumns.ID, db.RepositoryColumns.Name, db.RepositoryColumns.Owner, db.RepositoryColumns.CreatedAt,
	)
	// 2. エラー処理
	if err != nil {
		return nil, err
	}
	// 3. 戻り値の*model.User型を作る
	return convertRepository(repo), nil
}

func (r *repoService) GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,
			db.RepositoryColumns.Name,
			db.RepositoryColumns.Owner,
			db.RepositoryColumns.CreatedAt,
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}
	return convertRepository(repo), nil
}
