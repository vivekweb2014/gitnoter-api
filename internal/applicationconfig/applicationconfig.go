package applicationconfig

import (
	"github.com/git-noter/gitnoter-api/internal/auth"
	"github.com/git-noter/gitnoter-api/internal/config"
	"github.com/git-noter/gitnoter-api/internal/github"
	"github.com/git-noter/gitnoter-api/internal/preference"
	"github.com/git-noter/gitnoter-api/internal/user"
	"golang.org/x/oauth2"
	gh "golang.org/x/oauth2/github"
	"gorm.io/gorm"
)

// ApplicationConfig is an application config store used to store and get the application config & dependencies.
type ApplicationConfig struct {
	Config            config.Config
	DB                *gorm.DB
	OAuth2Config      oauth2.Config
	AuthService       auth.Service
	UserService       user.Service
	PreferenceService preference.Service
	GithubService     github.Service
}

// NewApplicationConfig creates and returns an application config store.
func NewApplicationConfig(config config.Config, db *gorm.DB) *ApplicationConfig {
	// create github oauth2 config
	oauth2Config := oauth2.Config{
		RedirectURL:  config.OAuth2.Github.RedirectURL,
		ClientID:     config.OAuth2.Github.ClientID,
		ClientSecret: config.OAuth2.Github.ClientSecret,
		Scopes:       []string{"read:user", "user:email", "repo"},
		Endpoint:     gh.Endpoint,
	}

	authService := auth.NewService(auth.TokenConfig{
		SecretKey: config.App.SecretKey,
		Issuer:    "https://gitnoter.com",
	})
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	preferenceRepo := preference.NewRepository(db)
	preferenceService := preference.NewService(preferenceRepo)

	githubClientBuilder := github.NewClientBuilder(&oauth2Config)
	githubService := github.NewService(githubClientBuilder)

	return &ApplicationConfig{
		Config:            config,
		DB:                db,
		OAuth2Config:      oauth2Config,
		AuthService:       authService,
		UserService:       userService,
		PreferenceService: preferenceService,
		GithubService:     githubService,
	}
}
