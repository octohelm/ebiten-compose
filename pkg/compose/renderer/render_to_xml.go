package renderer

import (
	"fmt"
	"io"

	"github.com/octohelm/ebiten-compose/pkg/compose/internal"
)

func RenderToXML(w io.Writer, n internal.Element) {
	_, _ = fmt.Fprintf(w, "<%s", n.String())

	_, _ = io.WriteString(w, ">")

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		RenderToXML(w, c.(internal.Element))
	}

	_, _ = fmt.Fprintf(w, "</%s>", n.String())
}
