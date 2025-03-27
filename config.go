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

func (configManager *ConfigManager) GetJWTSecretKey() string {
	return configManager.configInstance.GetString(jwtSecretKey)
}

func (configManager *ConfigManager) GetJWTHourLifeSpan() int {
	return configManager.configInstance.GetInt(jwtHourLifeSpan)
}

func (configManager *ConfigManager) GetRedisHost() string {
	return configManager.configInstance.GetString(redisHost)
}

func (configManager *ConfigManager) GetRedisPort() int {
	return configManager.configInstance.GetInt(redisPort)
}

func (configManager *ConfigManager) GetRedisPassword() string {
	return configManager.configInstance.GetString(redisPassword)
}

func (configManager *ConfigManager) GetRedisDatabase() int {
	return configManager.configInstance.GetInt(redisDB)
}
