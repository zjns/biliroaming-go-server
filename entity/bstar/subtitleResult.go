package bstar

type Subtitles struct {
	ID        int64  `json:"id"`
	Key       string `json:"key"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	IsMachine bool   `json:"is_machine"`
}

type SubtitleResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		SuggestKey string      `json:"suggest_key"`
		Subtitles  []Subtitles `json:"subtitles"`
	} `json:"data"`
}
