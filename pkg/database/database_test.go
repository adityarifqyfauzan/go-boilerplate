package database

import (
	"testing"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/spf13/viper"
)

func TestInitDB(t *testing.T) {
	conf := config.New(viper.New())

	type args struct {
		conf *config.AppConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-db-connection",
			args: args{
				conf: conf,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitDB(tt.args.conf)
		})
	}
}
