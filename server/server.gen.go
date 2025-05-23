// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /auth/login)
	LogIn(c *gin.Context)

	// (GET /auth/me)
	GetUserInfo(c *gin.Context)

	// (POST /auth/refresh)
	UpdateTokens(c *gin.Context)

	// (POST /auth/signup)
	SignUp(c *gin.Context)
	// Get user tasks with optional filters
	// (GET /tasks)
	GetUserTasks(c *gin.Context, params GetUserTasksParams)
	// Create a new task
	// (POST /tasks)
	CreateTask(c *gin.Context)

	// (DELETE /tasks/{id})
	DeleteTask(c *gin.Context, id int)
	// Patch task
	// (PATCH /tasks/{id})
	PatchTask(c *gin.Context, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// LogIn operation middleware
func (siw *ServerInterfaceWrapper) LogIn(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.LogIn(c)
}

// GetUserInfo operation middleware
func (siw *ServerInterfaceWrapper) GetUserInfo(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserInfo(c)
}

// UpdateTokens operation middleware
func (siw *ServerInterfaceWrapper) UpdateTokens(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateTokens(c)
}

// SignUp operation middleware
func (siw *ServerInterfaceWrapper) SignUp(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SignUp(c)
}

// GetUserTasks operation middleware
func (siw *ServerInterfaceWrapper) GetUserTasks(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserTasksParams

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", c.Request.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter status: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "deadline" -------------

	err = runtime.BindQueryParameter("form", true, false, "deadline", c.Request.URL.Query(), &params.Deadline)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter deadline: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserTasks(c, params)
}

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateTask(c)
}

// DeleteTask operation middleware
func (siw *ServerInterfaceWrapper) DeleteTask(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteTask(c, id)
}

// PatchTask operation middleware
func (siw *ServerInterfaceWrapper) PatchTask(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchTask(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/auth/login", wrapper.LogIn)
	router.GET(options.BaseURL+"/auth/me", wrapper.GetUserInfo)
	router.POST(options.BaseURL+"/auth/refresh", wrapper.UpdateTokens)
	router.POST(options.BaseURL+"/auth/signup", wrapper.SignUp)
	router.GET(options.BaseURL+"/tasks", wrapper.GetUserTasks)
	router.POST(options.BaseURL+"/tasks", wrapper.CreateTask)
	router.DELETE(options.BaseURL+"/tasks/:id", wrapper.DeleteTask)
	router.PATCH(options.BaseURL+"/tasks/:id", wrapper.PatchTask)
}
