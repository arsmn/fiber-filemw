package filemw

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber"
)

type Config struct {
	Prefix       string
	Root         http.FileSystem
	ErrorHandler func(*fiber.Ctx, error)
}

func New(config ...Config) func(*fiber.Ctx) {
	var cfg Config

	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.Root == nil {
		log.Fatal("Fiber: FileServer middleware requires root")
	}

	if cfg.Prefix == "" {
		cfg.Prefix = "/"
	}

	if !strings.HasPrefix(cfg.Prefix, "/") {
		cfg.Prefix = "/" + cfg.Prefix
	}

	if cfg.ErrorHandler == nil {
		cfg.ErrorHandler = func(c *fiber.Ctx, err error) {
			c.Status(fiber.StatusNotFound)
			c.SendString("File not found")
		}
	}

	return func(c *fiber.Ctx) {
		p := c.Path()

		if !strings.HasPrefix(p, cfg.Prefix) {
			c.Next()
			return
		}

		p = strings.TrimPrefix(p, cfg.Prefix)
		if c.Method() == fiber.MethodGet || c.Method() == fiber.MethodHead {
			file, err := cfg.Root.Open(filepath.Clean(p))
			if err != nil {
				cfg.ErrorHandler(c, err)
				return
			}

			stat, err := file.Stat()
			if err != nil {
				cfg.ErrorHandler(c, err)
				return
			}

			c.Fasthttp.SetBodyStream(file, int(stat.Size()))
			return
		}
		c.Next()
	}
}
