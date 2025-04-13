package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

const configPath string = "gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}
func getConfigPath() (string, error){
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := path.Join(homeDir, configPath)
	return path, nil

}
func ReadConfig() (Config, error) {

	path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil{
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(userName string) error{

	c.CurrentUserName = userName
	if err := c.write(); err != nil {
		return err
	}
	return nil
}

func (c *Config) write() error{
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error marshaling data")
	}
	if err := os.WriteFile(path, data, 0777); err != nil{
		return fmt.Errorf("Error writing to file")
	}
	return nil
}
