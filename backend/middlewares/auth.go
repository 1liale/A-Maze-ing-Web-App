package middlewares

import (
	"context"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CustomClaims struct {
	Permissions []string `json:"permissions"`
}

// Validate implements validator.CustomClaims. (for our purpose authentication is enough, no need to validate permissions as it is implied with auth)
func (*CustomClaims) Validate(context.Context) error {
	return nil
}

// CheckJWT is a middleware that validates JWT access tokens coming from the client
func CheckJWT(audience, domain string) gin.HandlerFunc {
	issuerURL, err := url.Parse("https://" + domain + "/")
	if err != nil {
		logrus.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	// Set up the validator.
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
	)
	if err != nil {
		logrus.Fatalf("Error setting up the validator: %v", err)
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {
			logrus.Printf("Encountered error while validating JWT: %v", err)
		}),
	)

	return func(ctx *gin.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			ctx.Request = r
			ctx.Next()
		}

		middleware.CheckJWT(handler).ServeHTTP(ctx.Writer, ctx.Request)

		if encounteredError {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT Access Token"})
		}
	}
}
