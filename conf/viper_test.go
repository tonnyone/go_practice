package conf

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestViper(t *testing.T) {
	viper := viper.GetViper()
	viper.AddConfigPath(".")
	viper.SetConfigName("a.json")
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		t.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}
	host := viper.GetString("host.port")
	t.Log(host)
}
