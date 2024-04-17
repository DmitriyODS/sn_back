package pdb

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
)

const (
	SqlSelectUser = `
SELECT u.id            as id,
       u.surname       as surname,
       u.name          as name,
       u.patronymic    as patronymic,
       u.gender_id     as gender_id,
       g.name          as gender_name,
       u.email         as email,
       u.phone         as phone,
       u.nick_tg       as nick_tg,
       u.nick_vk       as nick_vk,
       u.date_birthday as date_birthday,
       u.date_create   as date_create
FROM user_data."user" u
         LEFT JOIN user_data.gender g on g.id = u.gender_id
WHERE u.id = $1;
`
	SqlSelectUserList = `
SELECT u.id            as id,
       u.surname       as surname,
       u.name          as name,
       u.patronymic    as patronym,
       u.gender_id     as gender_id,
       g.name          as gender_name,
       u.email         as email,
       u.phone         as phone,
       u.nick_tg       as nick_tg,
       u.nick_vk       as nick_vk,
       u.date_birthday as date_birthday,
       u.date_create   as date_create
FROM user_data."user" u
         LEFT JOIN user_data.gender g on g.id = u.gender_id
`
	SqlInsertUser = `
INSERT INTO user_data."user" (
                	surname,
                    name,
                    patronymic,
                    gender_id,
                    email,
                    phone,
                    nick_tg,
                    nick_vk,
                    date_birthday)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;
`
	SqlUpdateUserByID = `
UPDATE user_data."user"
SET surname = $1,
    name = $2,
    patronymic = $3,
    gender_id = $4,
    email = $5,
    phone = $6,
    nick_tg = $7,
    nick_vk = $8,
    date_birthday = $9
WHERE id = $10;
`
	SqlDeleteUserByID = `
DELETE FROM user_data."user"
WHERE id = $1;
`
)

var fieldMapUsers = FieldsMap{
	"id":            "u.id",
	"surname":       "u.surname",
	"name":          "u.name",
	"patronymic":    "u.patronymic",
	"gender_id":     "u.gender_id",
	"email":         "u.email",
	"phone":         "u.phone",
	"nick_tg":       "u.nick_tg",
	"nick_vk":       "u.nick_vk",
	"date_birthday": "u.date_birthday",
	"date_create":   "u.date_create",
}

func (p *PDB) SelectUser(ctx context.Context, user domains.User) (domains.User, error) {
	var rows pgx.Rows
	var err error

	// здесь мы можем получить пользователя по различным идентификаторам
	if user.ID != 0 {
		rows, err = p.QueryTx(ctx, SqlSelectUser, user.ID)
	}

	if rows == nil {
		return domains.User{}, domains.ErrNoRows
	}
	if err != nil {
		return domains.User{}, err
	}
	defer rows.Close()

	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[domains.User])
	if errors.Is(err, pgx.ErrNoRows) {
		return domains.User{}, domains.ErrNoRows
	}
	if err != nil {
		return domains.User{}, err
	}

	return user, nil
}

func (p *PDB) SelectUserList(ctx context.Context, params global.OptionsList) (domains.UserList, error) {
	querySql := getQueryWithOptions(SqlSelectUserList, &params, fieldMapUsers)
	rows, err := p.QueryTx(ctx, querySql)
	if err != nil {
		return domains.UserList{}, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[domains.User])
	if errors.Is(err, pgx.ErrNoRows) {
		return domains.UserList{}, nil
	}
	if err != nil {
		return domains.UserList{}, err
	}

	return users, nil
}

func (p *PDB) InsertUser(ctx context.Context, user domains.User) (domains.User, error) {
	rows, err := p.QueryTx(ctx, SqlInsertUser, user.InsertPlaceholder()...)
	if err != nil {
		return domains.User{}, err
	}
	defer rows.Close()

	userID, err := pgx.CollectOneRow(rows, pgx.RowTo[uint64])
	if err != nil {
		return domains.User{}, err
	}
	user.ID = userID
	user.Auth.UserID = userID

	return user, nil
}

func (p *PDB) UpdateUser(ctx context.Context, user domains.User) error {
	commandTag, err := p.ExecTx(ctx, SqlUpdateUserByID, user.UpdatePlaceholder()...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return domains.ErrNoRows
	}

	return nil
}

func (p *PDB) DeleteUser(ctx context.Context, user domains.User) error {
	commandTag, err := p.ExecTx(ctx, SqlDeleteUserByID, user.ID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return domains.ErrNoRows
	}

	return nil
}
