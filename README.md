# fiber-filemw

FileServer middleware for `fiber`

### pkger

```go
package main

import (
	mw "github.com/arsmn/fiber-filemw"
	"github.com/gofiber/fiber"
	"github.com/markbates/pkger"
)

func main() {
	app := fiber.New()

	app.Use(New(Config{
		Prefix: "/assets",
		Root:   pkger.Dir("/assets"),
	}))

	app.Listen(8080)
}
```

### packr

```go
package main

import (
	mw "github.com/arsmn/fiber-filemw"
	"github.com/gofiber/fiber"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	app := fiber.New()
	assetsBox := packr.New("Assets Box", "/assets")

	app.Use(mw.New(mw.Config{
		Prefix: "/assets",
		Root:   assetsBox,
	}))

	app.Listen(8080)
}
```