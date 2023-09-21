package exception

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_notFoundError(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	type args struct {
		ctx *gin.Context
		err any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-not-found-error-[1]",
			args: args{
				ctx: c,
				err: NotFoundException{Error: "Not Found"},
			},
			want: true,
		},
		{
			name: "test-not-found-error-[2]",
			args: args{
				ctx: c,
				err: CredentialException{Error: "Credential Error"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := notFoundError(tt.args.ctx, tt.args.err); got != tt.want {
				t.Errorf("notFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_credentialError(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	type args struct {
		ctx *gin.Context
		err any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-creds-error-[1]",
			args: args{
				ctx: c,
				err: NotFoundException{Error: "Not Found"},
			},
			want: false,
		},
		{
			name: "test-creds-error-[2]",
			args: args{
				ctx: c,
				err: CredentialException{Error: "Credential Error"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := credentialError(tt.args.ctx, tt.args.err); got != tt.want {
				t.Errorf("credentialError() = %v, want %v", got, tt.want)
			}
		})
	}
}
