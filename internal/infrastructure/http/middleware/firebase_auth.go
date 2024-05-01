package middleware

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FirebaseAuth struct {
	client *auth.Client
}

func NewFirebaseAuth(client *auth.Client) *FirebaseAuth {
	return &FirebaseAuth{
		client: client,
	}
}

func (fa FirebaseAuth) EchoRoleBased(roles ...models.KindRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			authHeader, err := getAuthHeader(c.Request())
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			token, err := fa.client.VerifyIDToken(c.Request().Context(), authHeader)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			if !isAccess(roles, token.Claims["role"].(string)) {
				return echo.NewHTTPError(http.StatusForbidden, "access denied")
			}

			ctx = context.WithValue(ctx, userContextKey, User{
				Uuid:        token.UID,
				Email:       token.Claims["email"].(string),
				Role:        token.Claims["role"].(string),
				DisplayName: token.Claims["name"].(string),
			})

			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func (fa FirebaseAuth) Http(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		authHeader, err := getAuthHeader(request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := fa.client.VerifyIDToken(request.Context(), authHeader)
		if err != nil {
			http.Error(writer, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, userContextKey, User{
			Uuid:        token.UID,
			Email:       token.Claims["email"].(string),
			Role:        token.Claims["role"].(string),
			DisplayName: token.Claims["name"].(string),
		})

		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	})
}
