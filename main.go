package main

import (
	"encoding/base32"
	"time"

	"github.com/gofiber/fiber"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func main() {
	app := fiber.New()

	app.Get("/", generatePassCode)

	app.Listen(3001)
}

func generatePassCode(c *fiber.Ctx) {
	secret := base32.StdEncoding.EncodeToString([]byte("secret"))
	code, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})

	if err != nil {
		c.Fasthttp.Error(err.Error(), 503)

	} else {
		c.Send(code)
	}
}
