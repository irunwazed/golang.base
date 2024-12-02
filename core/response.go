package core

import "fmt"

func ResponseSuccess(c *Context, message string) error {
	c.SetContentType("application/json; charset=utf-8")
	c.SetStatusCode(200)
	c.WriteString(fmt.Sprintf(`{"status": true, "message": "%v"}`, message))
	return nil
}

func ResponseError(c *Context, message string) error {
	c.SetContentType("application/json; charset=utf-8")
	c.SetStatusCode(500)
	c.WriteString(fmt.Sprintf(`{"status": false, error: "ERROR", "message": "%v"}`, message))
	return nil
}

func ResponseBadRequest(c *Context, message string) error {
	c.SetContentType("application/json; charset=utf-8")
	c.SetStatusCode(402)
	c.WriteString(fmt.Sprintf(`{"status": false, error: "Bad Request", "message": "%v"}`, message))
	return nil
}

func ResponseUnAuthor(c *Context, message string) error {
	c.SetContentType("application/json; charset=utf-8")
	c.SetStatusCode(401)
	c.WriteString(fmt.Sprintf(`{"status": false, error: "Un Authorization", "message": "%v"}`, message))
	return nil
}

func ResponseNotFound(c *Context, message string) error {
	c.SetContentType("application/json; charset=utf-8")
	c.SetStatusCode(404)
	c.WriteString(fmt.Sprintf(`{"status": false, error: "Not Found", "message": "%v"}`, message))
	return nil
}

func ResponseText(c *Context, statusCode int, text string) error {
	c.SetContentType("text/plain; charset=utf-8")
	c.SetStatusCode(statusCode)
	c.WriteString(text)
	return nil
}
