package serve

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net"
	"net/http"
	"strconv"
)

func HTTP(log *zap.SugaredLogger, host string, port int, handler http.Handler) func(context.Context) error {
	return func(ctx context.Context) error {
		srv := &http.Server{
			Addr:    net.JoinHostPort(host, strconv.Itoa(port)),
			Handler: handler,
		}

		errCh := make(chan error, 1)
		go func() { errCh <- srv.ListenAndServe() }()
		log.Infow("http server is started",
			"host", host,
			"port", port,
		)
		defer log.Infow("http server shutdown")

		var err error
		select {
		case err = <-errCh:
		case <-ctx.Done():
			err = srv.Shutdown(context.Background())
		}
		if err != nil {
			return fmt.Errorf("srv.ListenAndServe: %w", err)
		}

		return nil
	}
}
