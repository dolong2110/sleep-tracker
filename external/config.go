package external

import (
	"context"
	"github.com/spf13/viper"
	"mindx/pkg/zapx"
)

type Configs struct {
	URLS         URLS   `mapstructure:"URLS"`
	Host         string `mapstructure:"HOST" default:"localhost"`
	Port         string `mapstructure:"PORT" default:"8000"`
	MaxBodyBytes int64  `mapstructure:"MAX_BODY_BYTES" default:"4194304"` // 4MB in Bytes ~ 4 * 1024 * 1024
	Timeout      int64  `mapstructure:"TIMEOUT" default:"5"`
	DS           DS     `mapstructure:"DATA_SOURCES"`
}

type URLS struct {
	APIURL   string `mapstructure:"API_URL" default:"api/"`
	UserURL  string `mapstructure:"USER_URL" default:"user/"`
	AdminURL string `mapstructure:"ADMIN_URL" default:"admin/"`
}

type DS struct {
	PostgreSQL PostgreSQL `mapstructure:"POSTGRE_SQL"`
}

type PostgreSQL struct {
	Host     string `mapstructure:"HOST" default:"postgres-server"`
	Port     string `mapstructure:"PORT" default:"5432"`
	User     string `mapstructure:"USER" default:"postgres"`
	Password string `mapstructure:"PASSWORD" required:"true"`
	DB       string `mapstructure:"DB" default:"postgres"`
	SSL      string `mapstructure:"SSL" default:"disable"`
	Timeout  int64  `mapstructure:"TIMEOUT" default:"10"`
}

// GetConfigs parse configs file from local into defined Configs struct - nested struct
func GetConfigs(path string, name string, fileType string) (*Configs, error) {
	var (
		err    error
		config *Configs
	)

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(fileType)
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		zapx.Error(context.TODO(), "failed to find and load config file.", err)
		return nil, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		zapx.Error(context.TODO(), "failed to parse config file.", err)
		return nil, err
	}

	return config, nil
}
