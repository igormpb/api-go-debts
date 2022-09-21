package utils

type TypeError struct {
	Message string
}

func (m *TypeError) Error(msg string) string {
	m.Message = msg
	return m.Message
}
