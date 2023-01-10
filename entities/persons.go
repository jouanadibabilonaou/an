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

func (p *Person) merge(otherValues map[string]string) {
	p.entity.merge(otherValues)
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
