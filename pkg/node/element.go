package node

import (
	"fmt"
)

var _ Node = &Element{}

type Element struct {
	Name                                                    string
	parent, firstChild, lastChild, prevSibling, nextSibling Node
}

func (e *Element) SetDisplayName(name string) bool {
	e.Name = name
	return e.Name != name
}

func (e *Element) String() string {
	return e.Name
}

func (e *Element) SetParentNode(n Node) {
	e.parent = n
}

func (e *Element) SetFirstChild(n Node) {
	e.firstChild = n
}

func (e *Element) SetLastChild(n Node) {
	e.lastChild = n
}

func (e *Element) SetPreviousSibling(n Node) {
	e.prevSibling = n
}

func (e *Element) SetNextSibling(n Node) {
	e.nextSibling = n
}

func (e *Element) ParentNode() Node {
	if e.parent == nil {
		return nil
	}
	return e.parent
}

func (e *Element) FirstChild() Node {
	if e.firstChild == nil {
		return nil
	}
	return e.firstChild
}

func (e *Element) LastChild() Node {
	if e.lastChild == nil {
		return nil
	}
	return e.lastChild
}

func (e *Element) PreviousSibling() Node {
	if e.prevSibling == nil {
		return nil
	}
	return e.prevSibling
}

func (e *Element) NextSibling() Node {
	if e.nextSibling == nil {
		return nil
	}
	return e.nextSibling
}

func (e *Element) InsertBefore(newChildNode, oldChildNode Node) Node {
	if oldChildNode == nil {
		e.AppendChild(newChildNode)
		return nil
	}

	if newChildNode.ParentNode() != nil || newChildNode.ParentNode() != nil || newChildNode.NextSibling() != nil {
		panic("insertBefore called for an attached child Element")
	}

	var prev, next Node

	if oldChildNode != nil {
		prev, next = oldChildNode.PreviousSibling(), oldChildNode.NextSibling()
	} else {
		prev = e.lastChild
	}
	if prev != nil {
		prev.SetNextSibling(newChildNode)
	} else {
		e.firstChild = newChildNode
	}
	if next != nil {
		next.SetPreviousSibling(newChildNode)
	} else {
		e.lastChild = newChildNode
	}

	newChildNode.SetParentNode(e)
	newChildNode.SetPreviousSibling(prev)
	newChildNode.SetNextSibling(next)

	return oldChildNode
}

func (e *Element) AppendChild(c Node) {
	if c == nil {
		return
	}

	if c.ParentNode() != nil || c.PreviousSibling() != nil || c.NextSibling() != nil {
		panic(fmt.Sprintf("appendChild called for an attached child Element: %v", c))
	}

	last := e.lastChild
	if last != nil {
		last.SetNextSibling(c)
	} else {
		e.firstChild = c
	}

	e.lastChild = c

	c.SetParentNode(e)
	c.SetPreviousSibling(last)
}

func (e *Element) RemoveChild(c Node) Node {
	if c.ParentNode() != e {
		panic("removeChild called for a non-child Element")
	}

	if e.firstChild == c {
		e.firstChild = c.NextSibling()
	}
	if c.NextSibling() != nil {
		c.NextSibling().SetPreviousSibling(c.PreviousSibling())
	}

	if e.lastChild == c {
		e.lastChild = c.PreviousSibling()
	}

	if c.PreviousSibling() != nil {
		c.PreviousSibling().SetNextSibling(c.NextSibling())
	}

	c.SetParentNode(nil)
	c.SetPreviousSibling(nil)
	c.SetNextSibling(nil)

	return c
}
