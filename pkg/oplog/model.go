package oplog

type LogContentRequest struct {
	OpID   string `json:"opID"`
	StepID string `json:"stepID"`
	Offset int64  `json:"offset"`
	Length int    `json:"length"`
}

type LogContentResponse struct {
	Content      string `json:"content"`
	LogSize      int64  `json:"logSize"`
	DeliverySize int64  `json:"deliverySize"`
}
