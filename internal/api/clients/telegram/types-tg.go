package telegram

type UpdatesResponse struct {
	ok          bool
	result      []Update
	description string
}

type Update struct {
	UpdateID int    `json:"update_id"`
	Message  string `json:"message"`
}
