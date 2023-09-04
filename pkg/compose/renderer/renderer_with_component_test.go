package renderer_test

import (
	"bytes"
	"context"
	"testing"

	testingx "github.com/octohelm/x/testing"

	. "github.com/octohelm/ebiten-compose/pkg/compose"
	. "github.com/octohelm/ebiten-compose/pkg/compose/modifier"
	"github.com/octohelm/ebiten-compose/pkg/compose/renderer"
)

type SubComp struct {
	Name string
}

func (a SubComp) Build(ctx BuildContext) VNode {
	return Box(DisplayName(a.Name))
}

type App struct {
	Name string
}

func (a App) Build(ctx BuildContext) VNode {
	return Box(DisplayName("AppRoot")).
		Children(H(SubComp{Name: a.Name}), Fragment(ctx.ChildVNodes()...))
}

func TestRenderComponent(t *testing.T) {
	ctx := context.Background()
	buf := bytes.NewBuffer(nil)
	r := renderer.CreateRoot()

	t.Run("Direct render", func(t *testing.T) {
		buf.Reset()

		r.Render(ctx, H(App{Name: "Sub"}).Children(Box(DisplayName("Appended"))))

		renderer.RenderToXML(buf, r.Root.Node())
		testingx.Expect(t, buf.String(), testingx.Equal("<Root><AppRoot><Sub></Sub><Appended></Appended></AppRoot></Root>"))

		t.Run("when prop changed, should render correct value", func(t *testing.T) {
			buf.Reset()

			r.Render(ctx, H(App{Name: "Changed"}).Children(Box(DisplayName("Appended"))))

			renderer.RenderToXML(buf, r.Root.Node())
			testingx.Expect(t, buf.String(), testingx.Equal("<Root><AppRoot><Changed></Changed><Appended></Appended></AppRoot></Root>"))
		})
	})

}
