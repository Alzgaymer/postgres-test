package db

import (
	"context"
	"errors"
	"fmt"
	"strings"
	customuser "test-postgres/custom_user"
	"test-postgres/postgressql"

	"github.com/jackc/pgconn"
)

type repository struct {
	client postgressql.Client
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

type Repository interface {
	Create(ctx context.Context, user customuser.CustomUser) error
	FindAll(ctx context.Context) (u []customuser.CustomUser, err error)
	FindOne(ctx context.Context, user customuser.CustomUser) (string, error)
	Update(ctx context.Context, user customuser.CustomUser) (string, error)
	Delete(ctx context.Context, user customuser.CustomUser) (string, error)
}

func (r *repository) Create(ctx context.Context, user customuser.CustomUser) error {

	querry := `insert into custom_user(name,age)
			   values($1)
				returning id
	`
	q := formatQuery(querry)
	if err := r.client.QueryRow(ctx, q, user.Name).Scan(&user.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))

			return newErr
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) (u []customuser.CustomUser, err error) {

	querry := `
	SELECT id, name, age FROM public.custom_user
	`
	q := formatQuery(querry)
	rows, err := r.client.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	users := make([]customuser.CustomUser, 0)

	for rows.Next() {
		var newUser customuser.CustomUser
		err = rows.Scan(&newUser.ID, &newUser.Name, &newUser.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, newUser)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
func (r *repository) FindOne(ctx context.Context, user customuser.CustomUser) (string, error) {
	panic("implement me")
}
func (r *repository) Update(ctx context.Context, user customuser.CustomUser) (string, error) {
	panic("implement me")
}
func (r *repository) Delete(ctx context.Context, user customuser.CustomUser) (string, error) {
	panic("implement me")
}
func NewRepository(client postgressql.Client) Repository {
	return &repository{client: client}
}
