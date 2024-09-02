package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var allowedHosts = [2]string{"gist.github.com", "gist.githubusercontent.com"}

func isHostnameAllowed(hostname string) bool {
	for i := 0; i < len(allowedHosts); i++ {
		if allowedHosts[i] == hostname {
			return true
		}
	}

	return false
}

// TODO: Handle image resizing

func ServeImage(c *fiber.Ctx) error {
	urlParam := strings.TrimSpace(c.Query("url", ""))

	if urlParam == "" {
		return fiber.NewError(http.StatusBadRequest, "Invalid url")
	}

	parsedUrl, err := url.Parse(urlParam)

	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid url")
	}

	hostname := parsedUrl.Hostname()

	if !isHostnameAllowed(hostname) {
		return fiber.NewError(http.StatusBadRequest, "Invalid url")
	}

	res, err := http.Get(parsedUrl.String())

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Server error")

	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Println("Error closing body", err)
		}
	}()

	imgExt := (path.Ext(urlParam))

	if imgExt == "" {
		imgExt = ".jpeg"
	}

	c.Response().Header.Set("Content-Type", "image/"+imgExt[1:])

	decoder := base64.NewDecoder(base64.StdEncoding, res.Body)

	_, err = io.Copy(c, decoder)

	return err
}
