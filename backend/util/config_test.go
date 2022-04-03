package util

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		wantConfig Config
		wantErr    bool
	}{
		{
			name: "get correct config info",
			args: args{
				path: "../",
			},
			wantConfig: Config{
				DB_DRIVER:      "postgres",
				DB_SOURCE:      "postgres://root:secret@localhost:5432/products?sslmode=disable",
				SERVER_ADDRESS: "localhost:8000",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfig, err := LoadConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("LoadConfig() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}
