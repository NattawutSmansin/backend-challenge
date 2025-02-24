package domain

type UseCase interface {
	GetPieFireDire() (map[string]int32, error)
}

type Repository interface {
	FetchBeefData() (string, error)
}
