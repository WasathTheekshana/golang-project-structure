package interfaces

type IInfoMessage struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Status  ResponseStatus `json:"status"`
}
