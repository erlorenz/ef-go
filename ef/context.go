package ef

import (
	"context"
	"log/slog"
)

// Context wraps context.Context with some extra methods.
type Context struct {
	context.Context
}

// Logger retrieves the slog.Logger.
func (c *Context) Logger() *slog.Logger {
	return getLogger(c)
}
