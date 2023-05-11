package request

type SuratKeteranganPenyerahanLaporanRequest struct {
	PemustakaId  string `json:"pemustaka_id" form:"pemustaka_id" validate:"required"`
	CollectionId string `json:"collection_id" form:"collection_id" validate:"required"`
	Title        string `json:"title" form:"title" validate:"required"`
}
