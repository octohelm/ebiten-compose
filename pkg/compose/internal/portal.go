package internal

var _ Component = Portal{}

type Portal struct{}

func (Portal) Build(BuildContext) VNode {
	return nil
}

func (Portal) IsRoot() bool {
	return true
}
