package service

type Service struct {
	Struct *StructService
}

func New(structStorage StructStorage, hasher Hasher) *Service {
	structService := NewStructService(structStorage, hasher)
	return &Service{
		Struct: structService,
	}
}
