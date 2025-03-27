package config

const (
	DEFAULT_CONFIG_FILE_NAME string = "config"

	database         = "database"
	databaseName     = database + ".name"
	databaseHost     = database + ".host"
	databasePort     = database + ".port"
	databaseUsername = database + ".username"
	databasePassword = database + ".password"

	jwt             = "jwt"
	jwtSecretKey    = jwt + ".secretKey"
	jwtHourLifeSpan = jwt + ".hourLifeSpan"
	jwtBearerPrefix = jwt + ".bearerPrefix"


	redis         = "redis"
	redisHost     = redis + ".host"
	redisPort     = redis + ".port"
	redisPassword = redis + ".password"
	redisDB       = redis + ".db"
)
