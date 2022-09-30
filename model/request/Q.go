package request

func init() {}

type Q struct {
	Queue   string `json:"queue"`
	Message string `json:"message"`
}
