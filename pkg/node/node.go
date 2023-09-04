package node

type Node interface {
	String() string

	ParentNode() Node
	FirstChild() Node
	LastChild() Node

	PreviousSibling() Node
	NextSibling() Node

	SetParentNode(n Node)
	SetFirstChild(n Node)
	SetLastChild(n Node)

	SetPreviousSibling(p Node)
	SetNextSibling(p Node)

	InsertBefore(newChild, referenceNode Node) Node
	RemoveChild(n Node) Node
	AppendChild(n Node)
}
