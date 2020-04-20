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