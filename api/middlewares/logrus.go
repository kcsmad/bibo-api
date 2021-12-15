package middlewares

import (
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type (
	LogrusConfig struct {
		FieldMap map[string]string
		Logger   logrus.FieldLogger
		Skipper  emw.Skipper
	}
)

var (
	DefaultLogrusConfig = LogrusConfig{
		FieldMap: map[string]string{
			"id":            "@id",
			"remote_ip":     "@remote_ip",
			"uri":           "@uri",
			"host":          "@host",
			"method":        "@method",
			"status":        "@status",
			"latency":       "@latency",
			"latency_human": "@latency_human",
			"bytes_in":      "@bytes_in",
			"bytes_out":     "@bytes_out",
		},
		Logger:  logrus.StandardLogger(),
		Skipper: emw.DefaultSkipper,
	}
)

func Logrus() echo.MiddlewareFunc {
	return LogrusWithConfig(DefaultLogrusConfig)
}

func LogrusWithConfig(config LogrusConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultLogrusConfig.Skipper
	}

	if config.Logger == nil {
		config.Logger = DefaultLogrusConfig.Logger
	}

	if len(config.FieldMap) == 0 {
		config.FieldMap = DefaultLogrusConfig.FieldMap
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			start := time.Now()

			if err = next(c); err != nil {
				c.Error(err)
			}

			stop := time.Now()
			entry := config.Logger

			for k, v := range config.FieldMap {
				if v == "" {
					continue
				}

				switch v {
				case "@id":
					id := req.Header.Get(echo.HeaderXRequestID)

					if id == "" {
						id = res.Header().Get(echo.HeaderXRequestID)
					}

					entry = entry.WithField(k, id)
				case "@remote_ip":
					entry = entry.WithField(k, c.RealIP())
				case "@uri":
					entry = entry.WithField(k, req.RequestURI)
				case "@host":
					entry = entry.WithField(k, req.Host)
				case "@method":
					entry = entry.WithField(k, req.Method)
				case "@path":
					p := req.URL.Path

					if p == "" {
						p = "/"
					}

					entry = entry.WithField(k, p)
				case "@referer":
					entry = entry.WithField(k, req.Referer())
				case "@user_agent":
					entry = entry.WithField(k, req.UserAgent())
				case "@status":
					entry = entry.WithField(k, res.Status)
				case "@latency":
					l := stop.Sub(start)
					entry = entry.WithField(k, strconv.FormatInt(int64(l), 10))
				case "@latency_human":
					entry = entry.WithField(k, stop.Sub(start).String())
				case "@bytes_in":
					cl := req.Header.Get(echo.HeaderContentLength)

					if cl == "" {
						cl = "0"
					}

					entry = entry.WithField(k, cl)
				case "@bytes_out":
					entry = entry.WithField(k, strconv.FormatInt(res.Size, 10))
				default:
					switch {
					case strings.HasPrefix(v, "@header:"):
						entry = entry.WithField(k, c.Request().Header.Get(v[8:]))
					case strings.HasPrefix(v, "@query:"):
						entry = entry.WithField(k, c.QueryParam(v[7:]))
					case strings.HasPrefix(v, "@form:"):
						entry = entry.WithField(k, c.FormValue(v[6:]))
					case strings.HasPrefix(v, "@cookie:"):
						cookie, err := c.Cookie(v[8:])
						if err == nil {
							entry = entry.WithField(k, cookie.Value)
						}
					}
				}
			}

			entry.Print("Handle request")

			return
		}
	}
}
