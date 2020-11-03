package models

import "encoding/json"

type CreateOrDeleteReaction struct {
	// Emoji the emoji short-code to add / remove. E.g. `:thumbsup::skin-tone-4:`.
	Emoji string `json:"emoji,omitempty"`
}

func (m *CreateOrDeleteReaction) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateOrDeleteReaction) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
