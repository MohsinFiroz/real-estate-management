package property

type Service struct {
	data *Data
}

func NewService(data *Data) *Service {
	return &Service{data: data}
}
