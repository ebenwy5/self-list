package list

// replikasi dari list
// Add()
// Get(i)
// Size()
// Remove(i) index
// Rules:
// - no append

type List struct {
	NodeMap map[int]*Node // new field
	Len     int
	Cap     int
}

// Node double linked list
type Node struct {
	Element string
	// Removed unnecessary field
	// Next    *Node
	// Prev    *Node
}

func NewList() List {
	return List{}
}

func (l *List) Add(input string) {
	n := Node{
		Element: input,
	}

	if l.Len == 0 {
		l.NodeMap = make(map[int]*Node, 2)
		l.NodeMap[0] = &n

		l.Cap = 2
	} else {
		// Amortized O(1)
		if l.Len == l.Cap {
			l.Cap *= 2
			tempMap := make(map[int]*Node, l.Cap)
			for k, v := range l.NodeMap {
				tempMap[k] = v
			}
		}

		l.NodeMap[l.Len] = &n

		// # PREVIOUS CODE, O(n)

		// currentNode := l.Nodes
		// for currentNode.Next != nil {
		// 	currentNode = currentNode.Next
		// }

		// n.Prev = currentNode
		// currentNode.Next = &n
	}

	l.Len++

	return
}

func (l *List) Get(index int) string {
	// O(1)
	val, ok := l.NodeMap[index]
	if !ok {
		return ""
	}

	return val.Element

	// # PREVIOUS CODE, O(n)

	// currentNode := l.Nodes
	// for i := 0; i < index; i++ {
	// 	currentNode = l.Nodes.Next
	// }
	// return currentNode.Element
}

func (l *List) Size() int {
	return l.Len
}

func (l *List) Remove(index int) string {
	var result string
	// handle empty list
	if l.Len == 0 {
		return ""
	}

	currentVal, ok := l.NodeMap[index]
	if !ok {
		return ""
	}

	result = currentVal.Element

	// handle single element
	if l.Len == 1 {
		delete(l.NodeMap, 0)
		l.Len--
		return result
	}

	// O(n)
	for i := index; i < l.Len; i++ {
		l.NodeMap[i] = l.NodeMap[i+1]
	}

	// delete last index
	l.Len--
	delete(l.NodeMap, l.Len)

	return result

	// # PREVIOUS CODE, O(n)

	// for i := 0; i < index; i++ {
	// 	currentNode = l.Nodes.Next
	// }

	// if currentNode.Next != nil {
	// 	currentNode.Next.Prev = currentNode.Prev
	// }
	// if currentNode.Prev != nil {
	// 	currentNode.Prev.Next = currentNode.Next
	// }

	// return currentNode.Element
}
