package middleware

import (
"net/http"
"strings"

"github.com/ViitoJooj/clown-crm/pkg/jwtTokens"
"github.com/gin-gonic/gin"
"github.com/golang-jwt/jwt/v4"
"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
return func(ctx *gin.Context) {
authHeader := ctx.GetHeader("Authorization")
if authHeader == "" {
ctx.JSON(http.StatusUnauthorized, gin.H{
"error": "authorization header required",
})
ctx.Abort()
return
}

parts := strings.Split(authHeader, " ")
if len(parts) != 2 || parts[0] != "Bearer" {
ctx.JSON(http.StatusUnauthorized, gin.H{
"error": "invalid authorization header format",
})
ctx.Abort()
return
}

tokenString := parts[1]
token, err := jwtTokens.ValidateToken(tokenString)
if err != nil {
ctx.JSON(http.StatusUnauthorized, gin.H{
"error": "invalid or expired token",
})
ctx.Abort()
return
}

claims, ok := token.Claims.(jwt.MapClaims)
if !ok || !token.Valid {
ctx.JSON(http.StatusUnauthorized, gin.H{
"error": "invalid token claims",
})
ctx.Abort()
return
}

userID, ok := claims["user_id"].(string)
if !ok {
ctx.JSON(http.StatusUnauthorized, gin.H{
"error": "invalid user_id in token",
})
ctx.Abort()
return
}

userUUID, err := uuid.Parse(userID)
if err != nil {
ctx.JSON(http.StatusUnauthorized, gin.H{
"error": "invalid user_id format",
})
ctx.Abort()
return
}

ctx.Set("user_id", userUUID)
ctx.Next()
}
}

func GetUserIDFromContext(ctx *gin.Context) (uuid.UUID, bool) {
userID, exists := ctx.Get("user_id")
if !exists {
return uuid.Nil, false
}

userUUID, ok := userID.(uuid.UUID)
return userUUID, ok
}
