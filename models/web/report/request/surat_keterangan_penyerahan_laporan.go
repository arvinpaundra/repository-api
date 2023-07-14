package request

type SuratKeteranganPenyerahanLaporanRequest struct {
	PemustakaId  string `json:"pemustaka_id" form:"pemustaka_id" validate:"required"`
	CollectionId string `json:"collection_id" form:"collection_id" validate:"required"`
	RepositoryId string `json:"repository_id" form:"repository_id" validate:"required"`
}
