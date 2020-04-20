# fiber-filemw

FileServer middleware for `fiber`

### pkger

```go
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
func main() {
	app := fiber.New()
	assetsBox := packr.New("Assets Box", "/assets")

	app.Use(New(Config{
		Prefix: "/assets",
		Root:   assetsBox,
	}))

	app.Listen(8080)
}
```