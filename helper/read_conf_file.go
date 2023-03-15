package helper

import "github.com/spf13/viper"

func ReadConfigFile(dir string, filename string) (string, error) {
	if dir == "" {
		dir = "/config"
	}
	filePath, err := GetProjectPath(dir, filename)
	if err != nil {
		return "", err
	}
	viper.SetConfigFile(filePath)
	readErr := viper.ReadInConfig()
	if err != nil {
		return "", readErr
	}
	return filePath, nil
}
