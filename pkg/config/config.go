package config

import "github.com/spf13/viper"

type Config struct {
	ChikensCount            int `mapstructure:"CHICKS_COUNT"`
	EggsMinSpawnCount       int `mapstructure:"EGGS_MIN_SPAWN_COUNT"`
	EggsMaxSpawnCount       int `mapstructure:"EGGS_MAX_SPAWN_COUNT"`
	EggsSpawnMinDelay       int `mapstructure:"EGGS_SPAWN_MIN_DELAY"`
	EggsSpawnMaxDelay       int `mapstructure:"EGGS_SPAWN_MAX_DELAY"`
	FarmerCheckMinDelay     int `mapstructure:"MIN_CHECK_DELAY"`
	FarmerCheckMaxDelay     int `mapstructure:"MAX_CHECK_DELAY"`
	FarmerMaxNeededQuantity int `mapstructure:"MAX_NEEDED_QUANTITY"`
	FarmerMinNeededQuantity int `mapstructure:"MIN_NEEDED_QUANTITY"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile("app.env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
