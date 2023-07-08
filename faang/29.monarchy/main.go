package main

import (
	"fmt"
)

type Monarchy struct{}

type Monarch struct {
	familyTree map[string]*Monarch
	name       string
	isAlive    bool
	children   []*Monarch
}

func initMonarchy(name string) *Monarch {
	m := new(Monarch)
	m.name = name
	m.isAlive = true
	m.familyTree = make(map[string]*Monarch)
	m.familyTree[name] = m
	return m
}

func (m *Monarch) newMonarch(name string) *Monarch {
	newM := new(Monarch)
	newM.name = name
	newM.isAlive = true
	m.familyTree[name] = newM
	newM.familyTree = m.familyTree
	return newM
}

func (m *Monarch) birth(child, parent string) {
	if p, ok := m.familyTree[parent]; ok {
		newMonarch := m.newMonarch(child)
		p.children = append(m.familyTree[parent].children, newMonarch)
	}
}

func (m *Monarch) death(name string) {
	if dead, ok := m.familyTree[name]; ok {
		dead.isAlive = false
	}
}

func (m *Monarch) getOrderOfSuccession() *[]string {
	order := new([]string)
	dfs(m, order)
	return order
}

func dfs(m *Monarch, order *[]string) {
	if m == nil {
		return
	}

	if m.isAlive {
		*order = append(*order, m.name)
	}

	for _, child := range m.children {
		dfs(child, order)
	}
}

func main() {
	mon := initMonarchy("Jake")

	mon.birth("Catherine", "Jake")
	mon.birth("Tom", "Jake")
	mon.birth("Celine", "Jake")
	mon.birth("Peter", "Celine")
	mon.birth("Jane", "Catherine")
	mon.birth("Farah", "Jane")
	mon.birth("Mark", "Catherine")

	o := mon.getOrderOfSuccession()
	fmt.Println(o)
	mon.death("Jake")
	mon.death("Jane")

	o = mon.getOrderOfSuccession()
	fmt.Println(o)
}
