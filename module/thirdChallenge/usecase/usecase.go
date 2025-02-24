package usecase

import (
	domain "backend-challenge/module/thirdChallenge/domain"
	"strings"
)

type UseCase struct {
	repo domain.Repository
}

func NewUseCase(repo domain.Repository) domain.UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (usecase *UseCase) GetPieFireDire() (map[string]int32, error) {
	fetchBeefData, err := usecase.repo.FetchBeefData()
	if err != nil {
		return nil, err
	}

	// ประเภทเนื้อทั้งหทด
	meatType := []string{
		"bacon", "beef", "pork", "sausage", "chicken", "buffalo", "lamb", "t-bone", "fatback", "pastrami", "meatloaf", "jowl",
		"enim", "bresaola", "ribs", "picanha", "chuck", "flank", "plank", "round", "tongue", "drumstick", "shank", "short loin",
		"tenderloin", "sirloin", "plate", "brisket", "strip steak",
	}

	fetchBeefData = strings.ToLower(fetchBeefData)

	dataMeat := make(map[string]int32)
	for _, meat := range meatType {
		dataMeat[meat] = int32(strings.Count(fetchBeefData, meat))
	}

	return dataMeat, nil
}
