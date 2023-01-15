package entities

type Location struct {
	entity AbstractEntity
}

func makeLocation(locationName string) Location {
	return Location{makeAbstractEntity(locationName)}
}

func (p *Location) setAttribute(key string, value string) {
	p.entity.setAttribute(key, value)
}

func (p *Location) addAttributes(otherValues map[string]string) {
	p.entity.addAttributes(otherValues)
}

func (p *Location) id() string {
	return p.entity.entityId
}

func (p *Location) name() string {
	return p.entity.entityName
}

func (p *Location) attributesLen() int {
	return p.entity.attributesLen()
}

func (p *Location) getAttribute(key string) (value string, ok bool) {
	return p.entity.getAttribute(key)
}

func (p *Location) addTag(tag string) {
	p.entity.addTag(tag)
}

func (p *Location) allTags() []string {
	return p.entity.allTags()
}

func (p *Location) allAttributesKey() []string {
	return p.entity.allAttributesKey()
}
