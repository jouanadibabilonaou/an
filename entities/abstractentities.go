package entities

type Entitier interface {
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
}

func makeAbstractEntity(entityName string) AbstractEntity {
	return AbstractEntity{entityId: generateId(), entityName: entityName, entityAttributes: make(map[string]string)}
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
