package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// contextKey is the Gin context key under which claims are stored.
const contextKey = "claims"

// RequireToken returns middleware that rejects requests without a valid
// "Authorization: Bearer <token>" header. On success it stores the parsed
// claims in the request context for downstream handlers.
func (m *TokenManager) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token, ok := strings.CutPrefix(header, "Bearer ")
		if !ok || strings.TrimSpace(token) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}

		claims, err := m.Parse(strings.TrimSpace(token))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidToken.Error()})
			return
		}

		c.Set(contextKey, claims)
		c.Next()
	}
}

// ClaimsFrom returns the authenticated claims stored by RequireToken.
func ClaimsFrom(c *gin.Context) (*Claims, bool) {
	v, ok := c.Get(contextKey)
	if !ok {
		return nil, false
	}
	claims, ok := v.(*Claims)
	return claims, ok
}
