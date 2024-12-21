package config

import (
	"template/internal/fiberserver"
	"template/pkg/env"
)

type Config struct {
	ClientTimeout      int // таймаут http клиента
	ServerReadTimeout  int // таймаут на чтение сервера
	ServerWriteTimeout int // таймаут сервера на запись

	TypeDetector string

	ConsumerAddress         string // адрес
	ReceiveEndpoint         string // # api endpoint
	ConsumerReceiveEndpoint string // адрес куда будут перенаправлены сообщения (не реализовано, не указывать в config)

	AddressToSend     string
	FlagSendToAddress bool

	IntervalVisor int

	TimeUpdate int

	Storage  StorageConfig
	SXTPUser SxtpUser

	Listen

	FiberServer fiberserver.Config
}

type Listen struct {
	HostEndpoint string // хост
	PortEndpoint int    // порт
	WithEndpoint bool   // включение api
	//ReceiveEndpoint string // # api endpoint
}

type SxtpUser struct {
	UserName string
	Password string
	HostSXTP string
	PortSXTP string
}

type StorageConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
	//DatabaseLog string
	PostgresUsername string
	PostgresPassword string

	MongoUsername string
	MongoPassword string
	MongoHost     string
	MongoPort     string

	RedisHost string
	RedisPort string
}

func NewConfig() Config {
	sxtpUser := SxtpUser{
		UserName: env.GetEnv("USER_NAME", DefaultUsernameSXTP),
		Password: env.GetEnv("PASSWORD", DefaultPasswordSXTP),
		HostSXTP: env.GetEnv("HOST_SXTP", DefaultHostSXTP),
		PortSXTP: env.GetEnv("PORT_SXTP", DefaultPortSXTP),
	}

	f := false
	storage := StorageConfig{
		PostgresHost:     env.GetEnv("HOST_DB", DefaultHostDB),
		PostgresPort:     env.GetEnv("PORT_DB", DefaultPortDB),
		PostgresDatabase: env.GetEnv("DATABASE", DefaultDBName),
		PostgresUsername: env.GetEnv("USERNAME_DB", DefaultUsernameDB),
		PostgresPassword: env.GetEnv("PASSWORD_DB", DefaultPasswordDB),
		RedisHost:        env.GetEnv("REDIS_HOST", DefaultRedisHost),
		RedisPort:        env.GetEnv("REDIS_PORT", DefaultRedisPort),
		MongoUsername:    env.GetEnv("MONGO_USERNAME", DefaultMongoUsername),
		MongoPassword:    env.GetEnv("MONGO_PASSWORD", DefaultMongoPassword),
		MongoHost:        env.GetEnv("MONGO_HOST", DefaultMongoHost),
		MongoPort:        env.GetEnv("MONGO_PORT", DefaultMongoPost),
	}

	config := Config{
		ClientTimeout:           env.GetEnvAsInt("CLIENT_TIMEOUT", DefaultClientTimeout),
		ServerReadTimeout:       env.GetEnvAsInt("SERVER_READ_TIMEOUT", DefaultServerReadTimeout),
		ServerWriteTimeout:      0,
		ConsumerAddress:         "",
		ReceiveEndpoint:         "",
		ConsumerReceiveEndpoint: "",
		TimeUpdate:              env.GetEnvAsInt("TIME_UPDATE", 800),
		IntervalVisor:           env.GetEnvAsInt("INTERVAL_VISOR", 5),
		Storage:                 storage,
		SXTPUser:                sxtpUser,
		Listen: Listen{
			HostEndpoint: env.GetEnv("HOST_ENDPOINTS", DefaultHost),
			PortEndpoint: env.GetEnvAsInt("PORT_ENDPOINTS", DefaultPort),
		},
		FiberServer: fiberserver.Config{
			Host:                        "0.0.0.0:8888",
			ShowUnknownErrorsInResponse: true,
			AllowOrigins:                "https://testtest.devap24.ru,https://testtestadmin.axarea.ru,https://testtest-admin.devap24.ru,https://aiforymaster.tech,https://admin.devap24.ru",
			AllowHeaders:                "app-version,content-type,ref-cache,authorization,sg,mt-key,refresh,access,shouldRetry,Access-Control-Allow-Origin,x-requested-with",
			ExposeHeaders:               "X-Trace-Id",
			IpHeader:                    "CF-Connecting-IP",
			SecureReqJsonPaths:          []string{"asd"},
			SecureResJsonPaths:          []string{"asd"},
			StreamRequestBody:           &f,
		},
	}
	return config
}
