package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
)

type (
	IJournalEntriesService interface {
		FindAll(request Dto.GetJournalRequest) (res []*Models.JournalEntry, meta Common.Meta, err error)
	}

	JournalEntriesService struct {
		JournalEntriesRepository Repositories.IJournalEntriesRepository
	}
)

func JournalEntriesProvider(journalRepo Repositories.IJournalEntriesRepository) *JournalEntriesService {
	return &JournalEntriesService{JournalEntriesRepository: journalRepo}
}

func (service *JournalEntriesService) FindAll(request Dto.GetJournalRequest) (res []*Models.JournalEntry, meta Common.Meta, err error) {
	res, totalData, err := service.JournalEntriesRepository.FindAll(request)
	if err != nil {
		return nil, Common.Meta{}, err
	}

	meta = Common.PaginateMetadata(nil, totalData, request.Limit, request.Page)

	return res, meta, nil
}
