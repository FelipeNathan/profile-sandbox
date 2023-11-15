package slack

type Body struct {
	Text   string `json:"text"`
	UserId string `json:"user_id"`
}

type ResponseBlock struct {
	Text         string `json:"text"`
	ResponseType string `json:"response_type"`
	Type         string `json:"type"`
}

type Response struct {
	Blocks []ResponseBlock `json:"blocks"`
}
