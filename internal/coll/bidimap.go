package coll

import (
	"fmt"
	"strings"
)

type Bidimap[K, V comparable] struct {
	forward  map[K]V
	backward map[V]K
}

func NewBidimap[K, V comparable]() *Bidimap[K, V] {
	return &Bidimap[K, V]{
		forward:  make(map[K]V),
		backward: make(map[V]K),
	}
}

func (b *Bidimap[K, V]) Put(k K, v V) {
	if oldV, ok := b.forward[k]; ok {
		delete(b.backward, oldV)
	}
	if oldK, ok := b.backward[v]; ok {
		delete(b.forward, oldK)
	}
	b.forward[k] = v
	b.backward[v] = k
}

func (b *Bidimap[K, V]) Remove(k K) {
	if _, exists := b.forward[k]; !exists {
		return
	}
	delete(b.backward, b.forward[k])
	delete(b.forward, k)
}

func (b *Bidimap[K, V]) ByValue(v V) (K, bool) {
	k, ok := b.backward[v]
	return k, ok
}

func (b *Bidimap[K, V]) ByKey(k K) (V, bool) {
	v, ok := b.forward[k]
	return v, ok
}

func (b *Bidimap[K, V]) HasKey(k K) bool {
	_, ok := b.forward[k]
	return ok
}

func (b *Bidimap[K, V]) HasValue(v V) bool {
	_, ok := b.backward[v]
	return ok
}

func (b *Bidimap[K, V]) String(short bool) string {
	indent := ""
	var sb strings.Builder
	if !short {
		indent = "  "
		sb.WriteString("Bidimap{\n")
	}
	for k, v := range b.forward {
		sb.WriteString(fmt.Sprintf("%s%v: %v,\n", indent, k, v))
	}
	if !short {
		sb.WriteString(indent + "}")
	}
	return sb.String()
}
