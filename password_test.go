package goreadonce

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	passwordStr := "Qwerty1234"

	p := NewPassword(passwordStr)
	actual, err := p.Value()

	assert.NoError(t, err)
	assert.Equal(t, passwordStr, actual)
}

func TestPasswordCanBeReadOnlyOnce(t *testing.T) {
	passwordStr := "Qwerty1234"

	p := NewPassword(passwordStr)

	// 1st call
	s, err := p.Value()
	assert.NoError(t, err)
	assert.NotEmpty(t, s)

	// 2nd call
	s, err = p.Value()
	assert.Error(t, err)
	assert.Empty(t, s)
}

func TestPasswordString(t *testing.T) {
	passwordStr := "Qwerty1234"
	p := NewPassword(passwordStr)

	printed := fmt.Sprint(p)
	assert.Equal(t, printed, "{credential}")

	printedDetail := fmt.Sprintf("%+v", p)
	assert.Equal(t, printedDetail, "{credential}")

	printedSharp := fmt.Sprintf("%#v", p)
	assert.Equal(t, printedSharp, "{credential}")
}

func TestPasswordMarshalJSON(t *testing.T) {
	passwordStr := "Qwerty1234"
	p := NewPassword(passwordStr)

	jsonB, err := json.Marshal(p)
	assert.NoError(t, err)
	assert.Equal(t, string(jsonB), "{}")
}

func TestPasswordMarshalXML(t *testing.T) {
	passwordStr := "Qwerty1234"
	p := NewPassword(passwordStr)

	xmlB, err := xml.Marshal(p)
	assert.NoError(t, err)
	assert.Equal(t, string(xmlB), "<Password></Password>")
}
