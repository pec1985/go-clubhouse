package models

type CreateOrDeleteReaction struct {
	// Emoji the emoji short-code to add / remove. E.g. `:thumbsup::skin-tone-4:`.
	Emoji string `json:"emoji,omitempty"`
}

func (m *CreateOrDeleteReaction) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *CreateOrDeleteReaction) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
