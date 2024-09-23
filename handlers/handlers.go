package handlers

import (
	"encoding/base64"
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

func invalidUrlErr() *fiber.Error {
	return fiber.NewError(http.StatusBadRequest, "Invalid url")
}

func ServeImage(c *fiber.Ctx) error {
	urlParam := strings.TrimSpace(c.Query("url", ""))

	if urlParam == "" {
		return invalidUrlErr()
	}

	parsedUrl, err := url.Parse(urlParam)

	if err != nil {
		return invalidUrlErr()
	}

	hostname := parsedUrl.Hostname()

	if !isHostnameAllowed(hostname) {
		return invalidUrlErr()
	}

	res, err := http.Get(parsedUrl.String())

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Server error")
	}

	defer res.Body.Close()

	// go doesn't return err if statusCode is != 200
	if res.StatusCode != http.StatusOK {
		return fiber.NewError(res.StatusCode, res.Status)
	}

	imgExt := path.Ext(urlParam)

	if imgExt == "" {
		imgExt = ".jpeg"
	}

	c.Response().Header.Set("Content-Type", "image/"+imgExt[1:])

	decoder := base64.NewDecoder(base64.StdEncoding, res.Body)

	_, err = io.Copy(c, decoder)

	return err
}
