package auth

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestLog(t *testing.T) {

	app := fiber.New()

	app.Post("/login", Login)

	body, err := json.Marshal(map[string]string{
		"pass": "password",
	})

	if err != nil {
		t.Errorf("error :%v", err.Error())
	}

	resp, err := app.Test(httptest.NewRequest("POST", "/login", bytes.NewBuffer(body)))

	if err != nil {
		t.Errorf("error :%v", err.Error())
	}

	if resp.StatusCode != 200 {
		t.Errorf("error status :%v", resp.Status)

	}

}
