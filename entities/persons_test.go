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
}

func TestMergeForPerson(t *testing.T) {
	var p Person = makePerson("Gerard")
	p.setAttribute("Nom", "Son nom")
	var newAttributes map[string]string = make(map[string]string)
	newAttributes["Prenom"] = "Gerard"
	p.merge(newAttributes)
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
	var impl Entitier = &p // because receiver is a pointer and not the value
	if len(impl.id()) == 0 {
		t.Fail()
	}
}
