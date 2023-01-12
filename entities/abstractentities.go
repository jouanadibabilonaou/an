package entities

import (
	"sort"
)

type IEntity interface {
	setAttribute(key string, value string)
	merge(otherValues map[string]string)
	id() string
	name() string
	attributesLen() int
	getAttribute(key string) (value string, ok bool)
}

type AbstractEntity struct {
	entityId         string
	entityName       string
	entityAttributes map[string]string
	entityTags       map[string]bool
}

func makeAbstractEntity(entityName string) AbstractEntity {
	return AbstractEntity{entityId: generateId(), entityName: entityName,
		entityAttributes: make(map[string]string), entityTags: make(map[string]bool)}
}

func (p *AbstractEntity) setAttribute(key string, value string) {
	if p.entityAttributes == nil {
		p.entityAttributes = make(map[string]string)
	}
	p.entityAttributes[key] = value
}

func (p *AbstractEntity) merge(otherValues map[string]string) {
	if otherValues != nil {
		if p.entityAttributes == nil {
			p.entityAttributes = make(map[string]string)
		}
		for key, value := range otherValues {
			p.entityAttributes[key] = value
		}
	}
}

func (p *AbstractEntity) attributesLen() int {
	return len(p.entityAttributes)
}

func (p *AbstractEntity) getAttribute(key string) (value string, ok bool) {
	value, ok = p.entityAttributes[key]
	return value, ok
}

func (p *AbstractEntity) addTag(tag string) {
	p.entityTags[tag] = true
}

func (p *AbstractEntity) allTags() []string {
	keys := make([]string, len(p.entityTags))
	var index int = 0
	for key := range p.entityTags {
		keys[index] = key
		index++
	}
	sort.Strings(keys)
	return keys
}
