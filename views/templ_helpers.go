package views

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func CSS(classes map[string]bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		for class, ok := range classes {
			if ok {
				_, err := io.WriteString(w, class+" ")
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
