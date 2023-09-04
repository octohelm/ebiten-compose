package renderer_test

import (
	"bytes"
	"context"
	"testing"

	. "github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/compose/modifier"
	"github.com/octohelm/ebiten-compose/pkg/compose/renderer"
	testingx "github.com/octohelm/x/testing"
)

func TestRender(t *testing.T) {
	ctx := context.Background()
	buf := bytes.NewBuffer(nil)
	r := renderer.CreateRoot()

	t.Run("Direct render", func(t *testing.T) {
		v := Box()

		buf.Reset()
		r.Render(ctx, v)

		renderer.RenderToXML(buf, r.Root.Node())
		testingx.Expect(t, buf.String(), testingx.Equal("<Root><Box></Box></Root>"))

		t.Run("Re-render should replace node", func(t *testing.T) {
			{
				v := Box(modifier.DisplayName("ReplacedBox"))

				buf.Reset()
				r.Render(ctx, v)
				renderer.RenderToXML(buf, r.Root.Node())
				testingx.Expect(t, buf.String(), testingx.Equal("<Root><ReplacedBox></ReplacedBox></Root>"))
			}
		})
	})

	t.Run("Re-render should insert node", func(t *testing.T) {
		{
			v := Box().Children(
				Box(modifier.DisplayName("Box1")),
				Box(modifier.DisplayName("Box2")),
			)

			buf.Reset()
			r.Render(ctx, v)
			renderer.RenderToXML(buf, r.Root.Node())
			testingx.Expect(t, buf.String(), testingx.Equal("<Root><Box><Box1></Box1><Box2></Box2></Box></Root>"))
		}

		{
			v := Box().Children(
				Box(modifier.DisplayName("Box1")),
				Box(modifier.DisplayName("Box3")),
				Box(modifier.DisplayName("Box2")),
			)

			buf.Reset()
			r.Render(ctx, v)
			renderer.RenderToXML(buf, r.Root.Node())
			testingx.Expect(t, buf.String(), testingx.Equal("<Root><Box><Box1></Box1><Box3></Box3><Box2></Box2></Box></Root>"))
		}
	})
}
