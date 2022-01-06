package main

import (
	"bitmex-subscriber/internal/srv/http"
	"bitmex-subscriber/pkg/serve"
	"context"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	handler := http.New(sugar, nil)

	if err := serve.Start(context.Background(), serve.HTTP(sugar, "localhost", 8000, handler.Init())); err != nil {
		sugar.Fatal("uuups")
	}
}
