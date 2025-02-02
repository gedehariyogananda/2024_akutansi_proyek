package Services

import (
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
)

type (
	ILogActivityService interface {
		FindAll(userID string) (res []Response.LogActivityResponse, err error)
	}

	LogActivityService struct {
		logActivityRepository Repositories.ILogActivityRepository
	}
)

func LogActivityServiceProvider(logActivityRepository Repositories.ILogActivityRepository) *LogActivityService {
	return &LogActivityService{logActivityRepository: logActivityRepository}
}

func (l *LogActivityService) FindAll(userID string) (res []Response.LogActivityResponse, err error) {
	logActivities, err := l.logActivityRepository.FindByUserId(userID)

	if err != nil {
		return nil, err
	}

	res = Response.ToLogActivitySlice(logActivities)

	return res, nil
}
