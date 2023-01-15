package entities

import "testing"

func TestGeneratePersonWithMandatoryInformation(t *testing.T) {
	var p Person = makePerson("Gerard")
	if len(p.id()) == 0 {
		t.Fail()
	}
	if p.name() != "Gerard" {
		t.Fail()
	}
	if p.attributesLen() != 0 {
		t.Fail()
	}
}

func TestAttributesForPerson(t *testing.T) {
	var p Person = makePerson("Gerard")
	if _, found := p.getAttribute("popo"); found {
		t.Fail()
	}
	p.setAttribute("Nom", "Son nom")
	if p.attributesLen() != 1 {
		t.Fail()
	}
	if value, found := p.getAttribute("Nom"); !found || value != "Son nom" {
		t.Fail()
	}
	p.setAttribute("Prenom", "Son prenom")
	var values []string = p.allAttributesKey()
	if len(values) != 2 || values[0] != "Nom" || values[1] != "Prenom" {
		t.Fail()
	}
}

func TestMergeForPerson(t *testing.T) {
	var p Person = makePerson("Gerard")
	p.setAttribute("Nom", "Son nom")
	var newAttributes map[string]string = make(map[string]string)
	newAttributes["Prenom"] = "Gerard"
	p.addAttributes(newAttributes)
	if p.attributesLen() != 2 {
		t.Fail()
	}
	if value, found := p.getAttribute("Nom"); !found || value != "Son nom" {
		t.Fail()
	}
	if value, found := p.getAttribute("Prenom"); !found || value != "Gerard" {
		t.Fail()
	}
}

func TestTags(t *testing.T) {
	var p Person = makePerson("Gerard")
	if len(p.allTags()) != 0 {
		t.Fail()
	}
	p.addTag("Personne")
	var values []string = p.allTags()
	if len(values) != 1 {
		t.Fail()
	}
	if values[0] != "Personne" {
		t.Fail()
	}
}

func TestComplianceWithInterface(t *testing.T) {
	var p Person = makePerson("Person")
	var impl IEntity = &p // because receiver is a pointer and not the value
	if len(impl.id()) == 0 {
		t.Fail()
	}
	p.setAttribute("Name", "A name")
	p.setAttribute("Age", "an age")
	if value, found := p.getAttribute("Name"); !found || value != "A name" {
		t.Fail()
	}
	p.addTag("Nice person")
	if tags := p.allTags(); len(tags) != 1 || tags[0] != "Nice person" {
		t.Fail()
	}
}
