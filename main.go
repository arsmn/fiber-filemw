package filemw

import (
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
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

			contentType, err := detectContentType(file, stat)
			if err != nil {
				cfg.ErrorHandler(c, err)
				return
			}

			c.Fasthttp.SetContentType(contentType)
			c.Fasthttp.SetBodyStream(file, int(stat.Size()))
			return
		}
		c.Next()
	}
}

func detectContentType(f http.File, fileInfo os.FileInfo) (string, error) {
	ext := getFileExtension(fileInfo.Name())
	contentType := mime.TypeByExtension(ext)
	if len(contentType) == 0 {
		data, err := readFileHeader(f)
		if err != nil {
			return "", err
		}
		contentType = http.DetectContentType(data)
	}

	return contentType, nil
}

func getFileExtension(path string) string {
	n := strings.LastIndexByte(path, '.')
	if n < 0 {
		return ""
	}

	return path[n:]
}

func readFileHeader(f http.File) ([]byte, error) {
	r := io.Reader(f)
	lr := &io.LimitedReader{
		R: r,
		N: 512,
	}
	data, err := ioutil.ReadAll(lr)
	if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}

	return data, err
}
