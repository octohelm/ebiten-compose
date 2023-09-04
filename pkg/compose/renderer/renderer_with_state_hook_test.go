package renderer_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/octohelm/ebiten-compose/pkg/compose/renderer"
	testingx "github.com/octohelm/x/testing"

	. "github.com/octohelm/ebiten-compose/pkg/compose"
	. "github.com/octohelm/ebiten-compose/pkg/compose/modifier"
)

type AppWithStateHook struct {
	Value       string
	UpdateValue **func(v string)
}

func (a AppWithStateHook) Build(ctx BuildContext) VNode {
	state := UseState(ctx, a.Value)

	updateValue := func(v string) {
		state.Update(v)
	}

	*a.UpdateValue = &updateValue

	return Box(DisplayName(state.Value())).Children(ctx.ChildVNodes()...)
}

func TestRenderWithStateHook(t *testing.T) {
	ctx := context.Background()
	buf := bytes.NewBuffer(nil)
	r := renderer.CreateRoot()

	t.Run("should re render when stage changed", func(t *testing.T) {
		var updateValue *func(v string)

		r.Render(ctx, H(AppWithStateHook{
			Value:       "StateHookInited",
			UpdateValue: &updateValue,
		}).Children(Box(DisplayName("Appended"))))

		buf.Reset()
		renderer.RenderToXML(buf, r.Root.Node())
		testingx.Expect(t, buf.String(), testingx.Equal("<Root><StateHookInited><Appended></Appended></StateHookInited></Root>"))

		r.Act(func() {
			(*updateValue)("StateHookUpdated")
		})

		buf.Reset()
		renderer.RenderToXML(buf, r.Root.Node())
		testingx.Expect(t, buf.String(), testingx.Equal("<Root><StateHookUpdated><Appended></Appended></StateHookUpdated></Root>"))
	})
}
