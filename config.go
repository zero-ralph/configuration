package config

import "github.com/spf13/viper"

type ConfigManager struct {
	configInstance *viper.Viper
}

func NewConfigManager() (configManager *ConfigManager) {
	configManager = &ConfigManager{
		configInstance: viper.New(),
	}
	configManager.configInstance.SetConfigName(DEFAULT_CONFIG_FILE_NAME)
	return
}

func (configManager *ConfigManager) ReadConfigFile(configFile string) (err error) {
	configManager.configInstance.SetConfigFile(configFile)
	if err = configManager.configInstance.ReadInConfig(); err != nil {
		panic(err)
	}
	return
}

func (configManager *ConfigManager) GetDatabaseHost() string {
	return configManager.configInstance.GetString(databaseHost)
}

func (configManager *ConfigManager) GetDatabasePort() int {
	return configManager.configInstance.GetInt(databasePort)
}

func (configManager *ConfigManager) GetDatabaseName() string {
	return configManager.configInstance.GetString(databaseName)
}

func (configManager *ConfigManager) GetDatabaseUsername() string {
	return configManager.configInstance.GetString(databaseUsername)
}

func (configManager *ConfigManager) GetDatabasePassword() string {
	return configManager.configInstance.GetString(databasePassword)
}
