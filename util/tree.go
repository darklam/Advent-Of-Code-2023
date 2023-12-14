package util

import "github.com/samber/mo"

type TreeNode[T any] struct {
	left  *TreeNode[T]
	right *TreeNode[T]
	value T
}

type Tree[T any] struct {
	root        *TreeNode[T]
	compareFunc func(a T, b T) int
}

func NewTree[T any](compareFunc func(T, T) int) *Tree[T] {
	return &Tree[T]{
		compareFunc: compareFunc,
	}
}

func (t *Tree[T]) Add(value T) {
	if t.root == nil {
		t.root = &TreeNode[T]{
			value: value,
		}
		return
	}
	current := t.root
	for {
		if t.compareFunc(value, current.value) < 0 {
			if current.left != nil {
				current = current.left
			} else {
				current.left = &TreeNode[T]{
					value: value,
				}
				break
			}
		} else {
			if current.right != nil {
				current = current.right
			} else {
				current.right = &TreeNode[T]{
					value: value,
				}
				break
			}
		}
	}
}

func (t *Tree[T]) Find(value T) mo.Option[T] {
	current := t.root
	for current != nil {
		comparison := t.compareFunc(value, current.value)
		if comparison == 0 {
			return mo.Some(current.value)
		} else if comparison < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}

	return mo.None[T]()
}

func (t *Tree[T]) FindFunc(compareFunc func(T) int) mo.Option[T] {
	current := t.root
	for current != nil {
		comparison := compareFunc(current.value)
		if comparison == 0 {
			return mo.Some(current.value)
		} else if comparison < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}

	return mo.None[T]()
}
