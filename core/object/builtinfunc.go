package object

import "github.com/gookit/slog"

func show(args ...Object) {
	for _, arg := range args {
		if str, ok := arg.(*String); ok {
			slog.Trace(str.Value)
		} else {
			slog.Trace(arg.String())
		}
	}
}
