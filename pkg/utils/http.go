package utils

func GetConfigPath(configPath string) string {
	if configPath == "prod" {
		return "./config/config"
	}
	return "./config/config-dev"
}