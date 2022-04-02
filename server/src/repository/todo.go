package repository

import (
	"database/sql"
	"errors"
	"todo/src/domain"
	"todo/src/util"

	"github.com/jmoiron/sqlx"
)

type TodoRepository interface {
	GetAllTodo() util.Result[[]domain.Todo]
	GetTodo(int) util.Result[domain.Todo]
	AddTodo(string) util.Result[domain.Todo]
	DeleteTodo(int) error
}

type TodoRepositoryImpl struct {
	db *sqlx.DB
}

func NewTodoRepositoryImpl(db *sqlx.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		db: db,
	}
}

func (r *TodoRepositoryImpl) GetAllTodo() util.Result[[]domain.Todo] {
	query := `SELECT id,content FROM todos`
	rows := util.NewResult(r.db.Query(query))
	todos := util.FlatMapResult(rows, func(rows *sql.Rows) util.Result[[]domain.Todo] {
		var res []domain.Todo
		for rows.Next() {
			todo := domain.Todo{}
			err := rows.Scan(&todo.Id, &todo.Content)
			if err != nil {
				return util.NewResultFailed[[]domain.Todo](err)
			}
			res = append(res, todo)
		}
		return util.NewResultSuccess(res)
	})
	return todos
}

func (r *TodoRepositoryImpl) GetTodo(id int) util.Result[domain.Todo] {
	query := `SELECT id,content FROM todos WHERE id = ?`
	todo := domain.Todo{}
	err := r.db.QueryRow(query, id).Scan(&todo.Id, &todo.Content)
	res := util.NewResult(todo, err)
	res = util.MapResultError(res, func(e error) error {
		return errors.New("Not Found Todo")
	})
	return res
}

func (r *TodoRepositoryImpl) AddTodo(content string) util.Result[domain.Todo] {
	query := `INSERT INTO todos (content) VALUES (?)`
	result := util.NewResult(r.db.Exec(query, content))
	id := util.FlatMapResult(result, func(result sql.Result) util.Result[int64] {
		return util.NewResult(result.LastInsertId())
	})
	todo := util.MapResult(id, func(id int64) domain.Todo {
		return domain.Todo{
			Id:      int(id),
			Content: content,
		}
	})
	todo = util.MapResultError(todo, func(_ error) error {
		return errors.New("Internal Server Error")
	})
	return todo
}

func (r *TodoRepositoryImpl) DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = ?`
	result := util.NewResult(r.db.Exec(query, id))
	return result.Error
}
