package entities

import "testing"

func TestGenerateIdShouldBeNonEmpty(t *testing.T) {
	var testId string = generateId()
	if len(testId) == 0 {
		t.Fail()
	}
}
