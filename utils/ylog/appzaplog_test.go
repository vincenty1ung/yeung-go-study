package ylog

import (
	"encoding/json"
	"fmt"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type payload struct {
	Level *zapcore.Level `json:"level"`
}

func Example() {
	if err := InitAppLog(TestEnv(false)); err != nil {
		fmt.Printf("InitAppLog err:%v", err)
	}
	defer func() {
		if err := Sync(); err != nil {
			fmt.Printf("InitAppLog err:%v", err)
		}
	}()
	Info("hello world", zap.String("author", "lbq"))
}

func TestJsonPayLoad(t *testing.T) {
	level := zapcore.Level(-1)
	pl := payload{
		Level: &level,
	}
	binarypl, err := json.Marshal(pl)
	if err != nil {
		t.Errorf("failed err:%v", err)
	}
	t.Logf("json str:%v", string(binarypl))
}
