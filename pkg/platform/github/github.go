package github

import (
	"os"

	"github.com/gothew/config"
	log "github.com/gothew/l-og"
)

type ConfigCredentials struct {
	Token string `yaml:"token"`
}

type GithubCredentials struct {
	PersonalToken string
}

// GetGithubCredentials return github credentials account.
func GetGithubCredentials() (GithubCredentials, error) {
	tokenEnvar := os.Getenv("GITHUB_PERSONAL_TOKEN")
	xdgEnvar := os.Getenv("XDG_CONFIG_HOME")
	//usr, err := user.Current()

	//if err != nil {
	//  fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}

	if len(tokenEnvar) != 0 {
		return getGithubCredentialsFromToken(tokenEnvar), nil
	}

	if len(xdgEnvar) != 0 {
		cfg, err := getConfig()
		if err != nil {
			log.SetOutput(os.Stderr)
			log.Fatalf("Error get config %v", err)
			os.Exit(1)
		}
		log.Info(cfg.Services.credentials)
	}

	return GithubCredentials{}, nil
}

func getGithubCredentialsFromToken(token string) GithubCredentials {
	return GithubCredentials{
		PersonalToken: token,
	}
}

//func getGithubCredentialsFromFile() (GithubCredentials, error) {
//	cfg, err := getConfig()
//	if err != nil {
//		log.SetOutput(os.Stderr)
//		log.Fatalf("Error get config %v", err)
//		os.Exit(1)
//	}
//
//  return GithubCredentials{
//    PersonalToken: cfg.Services.credentials.token,
//  }
//}

func fnConfigXDG() config.ConfigOptions {
	fnConfig := make(map[string]interface{})
	fnConfig["credentials"] = ConfigCredentials{Token: "xxx"}

	return config.ConfigOptions{
		Services: fnConfig,
	}
}

// TODO: pass for context from cmd file.
func getConfig() (config.ConfigOptions, error) {
	config.SetAppDir("nrs")
	config.SetConfigOptions(fnConfigXDG)
	cfg, err := config.ParseConfig()
	if err != nil {
		return config.ConfigOptions{}, err
	}
	return cfg, nil
}
