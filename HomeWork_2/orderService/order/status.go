package order

const (
	initialized = "Инициализирован"
	completed   = "Выполнен"
	deleted     = "Удален"
)

func isEditable(status string) bool {
	edit := false
	if status == initialized {
		edit = true
	}

	return edit
}
