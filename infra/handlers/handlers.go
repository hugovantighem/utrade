package handlers

import (
	"context"
	"net/http"

	"github.com/hugovantighem/utrade/api"
	"github.com/hugovantighem/utrade/app"
	"github.com/hugovantighem/utrade/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Api struct {
	repo domain.Repository
}

func NewApi(repo domain.Repository) (api.ServerInterface, error) {
	return Api{repo: repo}, nil
}

func (x Api) Ping(ctx echo.Context) error {
	result := api.Pong{Value: "pong"}
	return ctx.JSON(http.StatusOK, result)
}

func (x Api) SaveTodo(ctx echo.Context) error {
	// parse request
	cmd := &api.SaveTodoCommand{}
	if err := ctx.Bind(cmd); err != nil {
		return err
	}

	// retrieve user ID from request
	userID, err := extratUserID(ctx)
	if err != nil {
		logrus.Error(err)

		errStr := err.Error()
		resp := api.Response{
			Error: &errStr,
		}

		return ctx.JSON(http.StatusOK, resp)
	}

	// call usecase
	err = app.SaveTodo(context.Background(), x.repo, app.SaveTodoCmd{
		Msg:    cmd.Msg,
		UserID: userID,
	})

	// response
	if err != nil {
		logrus.Errorf("save todo error: %v", err)

		errStr := err.Error()
		resp := api.Response{
			Error: &errStr,
		}
		return ctx.JSON(http.StatusOK, resp)
	}

	result := api.Response_Result{}
	result.FromSaveTodoResult(api.SaveTodoResult{Result: "ok"})
	resp := api.Response{
		Result: &result,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (x Api) GetTodo(ctx echo.Context, userid string) error {

	// call usecase
	item, err := app.GetTodo(context.Background(), x.repo, app.GetTodoQuery{
		UserID: userid,
	})

	// response
	if err != nil {
		logrus.Errorf("get todo error: %v", err)

		errStr := err.Error()
		resp := api.Response{
			Error: &errStr,
		}
		return ctx.JSON(http.StatusOK, resp)
	}

	result := api.Response_Result{}
	result.FromTodo(api.Todo{Msg: item.Message(), Status: string(item.Status()), UserID: item.UserID()})
	resp := api.Response{
		Result: &result,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (x Api) ListTodos(ctx echo.Context) error {
	// call usecase
	items, err := app.ListCreated(context.Background(), x.repo)

	// response
	if err != nil {
		logrus.Errorf("list todos error: %v", err)

		errStr := err.Error()
		resp := api.Response{
			Error: &errStr,
		}
		return ctx.JSON(http.StatusOK, resp)
	}
	todos := make([]api.Todos_Item, 0, len(items))

	for _, item := range items {
		todo := api.Todos_Item{}
		todo.FromTodo(api.Todo{
			Msg:    item.Message(),
			Status: string(item.Status()),
			UserID: item.UserID(),
		})

		todos = append(todos, todo)
	}

	result := api.Response_Result{}
	result.FromTodos(todos)
	resp := api.Response{
		Result: &result,
	}

	return ctx.JSON(http.StatusOK, resp)
}

// extratUserID extract the user id from the incoming request.
func extratUserID(ctx echo.Context) (string, error) {
	result := ctx.Request().Header["User-Id"][0]
	if len(result) == 0 {
		return "", domain.InvalidArgumentError{ArgName: "header \"User-Id\"", Reason: "should not be empty"}
	}

	return result, nil
}
