package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"BACKEND/internal/models"
	"BACKEND/internal/service"
)

const (
	AuthUserKey = "authUser"
)

func GetAuthUser(c *fiber.Ctx) *models.AuthUser {
	user, ok := c.Locals(AuthUserKey).(models.AuthUser)
	if !ok {
		return nil
	}
	return &user
}

func Auth(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			if logger != nil {
				logger.Warn("missing authorization header", zap.String("path", c.Path()))
			}
			return models.SendUnauthorized(c, "Missing authorization header", GetRequestID(c))
		}

		authHeaderLower := strings.ToLower(authHeader)
		if !strings.HasPrefix(authHeaderLower, "bearer ") {
			if logger != nil {
				logger.Warn("invalid authorization header format", zap.String("path", c.Path()))
			}
			return models.SendUnauthorized(c, "Invalid authorization header format. Expected: Bearer <token>", GetRequestID(c))
		}

		parts := strings.Fields(authHeader)
		if len(parts) < 2 {
			if logger != nil {
				logger.Warn("empty token", zap.String("path", c.Path()))
			}
			return models.SendUnauthorized(c, "Token is required", GetRequestID(c))
		}

		if !strings.EqualFold(parts[0], "bearer") {
			if logger != nil {
				logger.Warn("invalid authorization header format", zap.String("path", c.Path()))
			}
			return models.SendUnauthorized(c, "Invalid authorization header format. Expected: Bearer <token>", GetRequestID(c))
		}

		tokenString := strings.Join(parts[1:], " ")
		if tokenString == "" {
			if logger != nil {
				logger.Warn("empty token", zap.String("path", c.Path()))
			}
			return models.SendUnauthorized(c, "Token is required", GetRequestID(c))
		}

		token, err := jwt.ParseWithClaims(tokenString, &service.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			if logger != nil {
				logger.Warn("token validation failed", zap.Error(err), zap.String("path", c.Path()))
			}
			return models.SendError(c, fiber.StatusUnauthorized, "Invalid or expired token", models.ErrCodeInvalidToken, GetRequestID(c))
		}

		claims, ok := token.Claims.(*service.JWTClaims)
		if !ok || !token.Valid {
			if logger != nil {
				logger.Warn("invalid token claims", zap.String("path", c.Path()))
			}
			return models.SendError(c, fiber.StatusUnauthorized, "Invalid token claims", models.ErrCodeInvalidToken, GetRequestID(c))
		}

		authUser := models.AuthUser{
			ID:   claims.UserID,
			Role: claims.Role,
		}
		c.Locals(AuthUserKey, authUser)

		if logger != nil {
			logger.Info("user authenticated",
				zap.Int32("user_id", authUser.ID),
				zap.String("role", authUser.Role),
				zap.String("path", c.Path()),
			)
		}

		return c.Next()
	}
}

func RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
	
		authUser := GetAuthUser(c)
		if authUser == nil {
			if logger != nil {
				logger.Warn("role check failed: no authenticated user in context",
					zap.String("path", c.Path()),
				)
			}
			return models.SendUnauthorized(c, "Unauthorized", GetRequestID(c))
		}

		hasRole := false
		for _, role := range allowedRoles {
			if authUser.Role == role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			if logger != nil {
				logger.Warn("role check failed: insufficient permissions",
					zap.Int32("user_id", authUser.ID),
					zap.String("user_role", authUser.Role),
					zap.Strings("required_roles", allowedRoles),
					zap.String("path", c.Path()),
				)
			}
			return models.SendError(c, fiber.StatusForbidden, "Forbidden: insufficient permissions", models.ErrCodeInsufficientPerms, GetRequestID(c))
		}

		if logger != nil {
			logger.Info("role check passed",
				zap.Int32("user_id", authUser.ID),
				zap.String("role", authUser.Role),
				zap.String("path", c.Path()),
			)
		}

		return c.Next()
	}
}
