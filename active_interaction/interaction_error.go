package active_interaction

import (
	"fmt"
	"strings"
)

type InteractionError struct {
	errors map[string][]string
}

func (m InteractionError) Error() string {
	var fields []string = []string{}
	for key, element := range m.errors {
		messages := strings.Join(element, "\",\"")
		fields = append(fields, fmt.Sprintf("\"%s\": [%s]", key, "\""+messages+"\""))
	}
	return strings.Join(fields, ", ")
}

func (m *InteractionError) AddError(field string, message string) {
	if m.errors == nil {
		m.errors = map[string][]string{}
	}
	m.errors[field] = append(m.errors[field], message)
}

func (m InteractionError) ErrorMap() map[string][]string {
	return m.errors
}
