package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const configPath string = "gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func ReadConfig() Config {

	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading file")
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil{
		fmt.Println("Error parsing JSON")
	}

	return config
}

func (c *Config) SetUser(userName string) error{

	c.CurrentUserName = userName
	if err := c.write(); err != nil {
		return err
	}
	return nil
}

func (c *Config) write() error{
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error marshaling data")
	}
	if err := os.WriteFile(configPath, data, 0777); err != nil{
		return fmt.Errorf("Error writing to file")
	}
	return nil
}
