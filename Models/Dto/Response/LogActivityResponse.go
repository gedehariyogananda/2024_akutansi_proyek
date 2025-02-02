package Response

import (
	"2024_akutansi_project/Models"
)

type LogActivityResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Device    string `json:"device"`
	CreatedAt string `json:"created_at"`
}

func ToLogActivity(logActivity *Models.LogActivity) LogActivityResponse {
	return LogActivityResponse{
		ID:        logActivity.ID,
		Name:      logActivity.Name,
		Device:    logActivity.Device,
		CreatedAt: logActivity.CreatedAt.Format("02 January 2006 15 : 04"),
	}
}

func ToLogActivitySlice(logActivities []*Models.LogActivity) []LogActivityResponse {
	var logActivityList []LogActivityResponse
	for _, logActivity := range logActivities {
		logActivityList = append(logActivityList, ToLogActivity(logActivity))
	}
	return logActivityList
}
