package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/secure"
	"github.com/rs/cors"
)

type Config struct {
	Port          string
	SecureOptions secure.Options
	CorsOptions   cors.Options
	Audience      string
	Domain        string
}

// Enforces SSL headers
func SecureOptions() secure.Options {
	return secure.Options{
		AllowedHosts:          []string{"example.com", "ssl.example.com"},
		SSLRedirect:           true,
		SSLHost:               "ssl.example.com",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
	}
}

// Specify options that enable cross origin resource sharing
func CorsOptions(clientOriginUrl string) cors.Options {
	return cors.Options{
		AllowedOrigins: []string{clientOriginUrl},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:         86400,
	}
}

func InitConfig() *Config {
	port := ":" + os.Getenv("API_PORT")
	clientOriginUrl := os.Getenv("CLIENT_ORIGIN_URL")
	audience := os.Getenv("AUTH0_AUDIENCE")
	domain := os.Getenv("AUTH0_DOMAIN")

	config := Config{
		Port:          port,
		SecureOptions: SecureOptions(), // don't use in development
		CorsOptions:   CorsOptions(clientOriginUrl),
		Audience:      audience,
		Domain:        domain,
	}

	fmt.Printf("%+v\n\n", config)

	return &config
}
