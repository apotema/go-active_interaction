package active_interaction

import (
	"fmt"
	"regexp"
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

func (m *InteractionError) Append(other *InteractionError, prefix string) {
	for k, v := range other.errors {
		prefixedKey := fmt.Sprintf("%s.%s", prefix, k)
		for _, msg := range v {
			m.AddError(prefixedKey, msg)
		}
	}
}

func (m *InteractionError) AddValidatorError(errors error) InteractionError {
	fieldRegex := regexp.MustCompile(`[^']*\.([^']*)'`)
	descriptionRegex := regexp.MustCompile(`Error:(.*)`)

	for _, element := range strings.Split(errors.Error(), "\n") {
		var field string
		var description string

		matches := fieldRegex.FindStringSubmatch(element)
		if len(matches) > 1 {
			field = matches[1]
		}

		descriptions := descriptionRegex.FindStringSubmatch(element)
		if len(descriptions) > 1 {
			description = descriptions[1]
		}

		m.AddError(field, description)
	}

	return *m
}

func (m InteractionError) ErrorMap() map[string][]string {
	return m.errors
}

func (m InteractionError) HasError() bool {
	return len(m.errors) > 0
}
