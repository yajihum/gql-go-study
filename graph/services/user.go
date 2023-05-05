package services

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/yajium/gql-go-study/graph/db"
	"github.com/yajium/gql-go-study/graph/model"
)

// private
type userService struct {
	exec boil.ContextExecutor // クエリの実行に使用されるデータベース接続を提供するインターフェース
}

// private
// db.User型からmodel.User型に変換する
func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

/*
ユーザーネームからmodel.Userオブジェクトを取得する
csx: リクエストの寿命やキャンセル信号などの情報を保持
name: 検索対象のユーザー名
*/
func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す
	user, err := db.Users( // from users
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name), // select id, name
		db.UserWhere.Name.EQ(name),                                  // where name = {引数nameの内容}
	).One(ctx, u.exec) // limit 1 Oneメソッドで1つのレコードを取得
	// 2. エラー処理
	if err != nil {
		return nil, err
	}
	// 3. 戻り値の*model.User型を作る
	return convertUser(user), nil
}
