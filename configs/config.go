package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURI               string `mapstructure:"mongo_uri" env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
	DBName                 string `mapstructure:"db_name" env:"DB_NAME" envDefault:"employee_db"`
	ConnectionTimeout      int    `mapstructure:"connection_timeout" env:"CONNECTION_TIMEOUT" envDefault:"10"`
	MaxPoolSize            int    `mapstructure:"max_pool_size" env:"MAX_POOL_SIZE" envDefault:"10"`
	MinPoolSize            int    `mapstructure:"min_pool_size" env:"MIN_POOL_SIZE" envDefault:"1"`
	MaxConnIdleTime        int    `mapstructure:"max_conn_idle_time" env:"MAX_CONN_IDLE_TIME" envDefault:"10"`
	SocketTimeout          int    `mapstructure:"socket_timeout" env:"SOCKET_TIMEOUT" envDefault:"10"`
	ServerSelectionTimeout int    `mapstructure:"server_selection_timeout" env:"SERVER_SELECTION_TIMEOUT" envDefault:"10"`
	Port                   string `mapstructure:"port" env:"PORT" envDefault:"8080"`
}

var AppConfig *Config

func LoadConfig() {
	// Load the configuration from environment variables or default values
	// This is a placeholder for the actual implementation
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	log.Printf("Loaded config: %+v", AppConfig)

}
