package utils

import (
	"github.com/dmxmss/tasks/entities"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"time"
)

func ToGetTaskDto(task *entities.Task) entities.GetTaskDto {
	return entities.GetTaskDto{
			ID: task.ID,
			CreatedAt: task.CreatedAt,
			Name: task.Name,
			Description: task.Description,
			Status: task.Status,
			Deadline: task.Deadline,
			Tags: task.Tags,
			Weather: task.Weather,
			UserID: task.UserID,
		}
}

func FromCreateTaskDto(createTask entities.CreateTaskDto) entities.Task {
	return entities.Task{
		Name: createTask.Name,
		Description: createTask.Description,
		Deadline: createTask.Deadline,
		Tags: createTask.Tags,
	}
}

func WriteTokenToCookies(c *gin.Context, accessToken, refreshToken string, accessExpire, refreshExpire int) {
	c.SetCookie(
    "access_token",
    accessToken,
		accessExpire,
		"/", 
		"", 
		true, 
		true,
	)

	c.SetCookie(
    "refresh_token",
    refreshToken,
		refreshExpire,
		"/auth/refresh", 
		"", 
		true, 
		true,
	)
}

func GetClaims(userId int, city string, expirationTime int) entities.Claims {
	return entities.Claims{
		UserID: userId,
		City: city,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationTime)*time.Second)),
		},
	}
}
