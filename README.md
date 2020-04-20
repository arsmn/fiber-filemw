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
	dir := pkger.Dir("/assets")

	app.Use(mw.New(mw.Config{
		Prefix: "/assets",
		Root:  dir,
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

### go.rice

```go
package main

import (
	mw "github.com/arsmn/fiber-filemw"
	"github.com/gofiber/fiber"
	rice "github.com/GeertJohan/go.rice"
)

func main() {
	app := fiber.New()
	assetsBox := rice.MustFindBox("assets")

	app.Use(mw.New(mw.Config{
		Prefix: "/assets",
		Root:   assetsBox.HTTPBox(),
	}))

	app.Listen(8080)
}
```