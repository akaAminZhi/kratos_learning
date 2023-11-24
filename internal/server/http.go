package server

import (
	"fmt"
	httpn "net/http"
	v1 "shortUrl/api/shortUrl/v1"
	"shortUrl/internal/conf"
	"shortUrl/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func redirecFunc(srv v1.ShortUrlHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetUrlRequest
		if err := ctx.BindQuery(&in); err != nil {
			fmt.Println("heiheihei")

			return err
		}

		if err := ctx.BindVars(&in); err != nil {
			return err
		}

		// fmt.Println(in)
		reply, err := srv.GetUrl(ctx, &in)
		if err != nil {
			fmt.Println(err)

			return err
		}
		fmt.Println(reply)
		httpn.Redirect(ctx.Response(), ctx.Request(), reply.LongUrl, 302)
		return nil
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, short *service.ShortUrlService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),

			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	srv.Route("/").GET("/v1/{shortUrl}", redirecFunc(short))
	v1.RegisterShortUrlHTTPServer(srv, short)
	return srv
}
