package Services

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
)

type (
	IWaitingListService interface {
		InsertWaitingList(request *Dto.MakeWaitingListRequest) (err error)
	}

	WaitingListService struct {
		WaitingListRepo Repositories.IWaitingListRepository
	}
)

func WaitingListServiceProvider(WaitingListRepo Repositories.IWaitingListRepository) *WaitingListService {
	return &WaitingListService{WaitingListRepo: WaitingListRepo}
}

func (s *WaitingListService) InsertWaitingList(request *Dto.MakeWaitingListRequest) (err error) {
	if err := s.WaitingListRepo.InsertWaitingList(request); err != nil {
		return err
	}

	return nil
}
