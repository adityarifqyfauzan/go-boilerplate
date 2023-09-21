package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestAppConfig_GetString(t *testing.T) {
	type fields struct {
		conf *viper.Viper
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "test-config-get-string",
			fields: fields{
				conf: viper.New(),
			},
			args: args{
				key: "app.name",
			},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.fields.conf)
			assert.Equal(t, "go-boilerplate", c.GetString("app.name"))
		})
	}
}
