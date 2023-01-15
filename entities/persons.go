package entities

type Person struct {
	entity AbstractEntity
}

func makePerson(fullName string) Person {
	return Person{makeAbstractEntity(fullName)}
}

func (p *Person) setAttribute(key string, value string) {
	p.entity.setAttribute(key, value)
}

func (p *Person) addAttributes(otherValues map[string]string) {
	p.entity.addAttributes(otherValues)
}

func (p *Person) id() string {
	return p.entity.entityId
}

func (p *Person) name() string {
	return p.entity.entityName
}

func (p *Person) attributesLen() int {
	return p.entity.attributesLen()
}

func (p *Person) getAttribute(key string) (value string, ok bool) {
	return p.entity.getAttribute(key)
}

func (p *Person) addTag(tag string) {
	p.entity.addTag(tag)
}

func (p *Person) allTags() []string {
	return p.entity.allTags()
}

func (p *Person) allAttributesKey() []string {
	return p.entity.allAttributesKey()
}
