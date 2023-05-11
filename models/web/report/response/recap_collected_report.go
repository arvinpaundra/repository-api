package response

import "github.com/arvinpaundra/repository-api/models/domain"

type RecapCollectedReportResponse struct {
	StudyProgram   string `json:"study_program"`
	PemustakaCount int64  `json:"pemustaka_count"`
}

func ToRecapCollectedReportResponse(report domain.Report) RecapCollectedReportResponse {
	return RecapCollectedReportResponse{
		StudyProgram:   report.StudyProgram,
		PemustakaCount: report.PemustakaCount,
	}
}

func ToRecapCollectedReportArrayResponse(reports []domain.Report) []RecapCollectedReportResponse {
	var responses []RecapCollectedReportResponse

	for _, report := range reports {
		responses = append(responses, ToRecapCollectedReportResponse(report))
	}

	return responses
}
