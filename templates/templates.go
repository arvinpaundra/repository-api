package templates

import _ "embed"

//go:embed template-forgot-password.html
var ForgotPassword string

//go:embed template-id-card.html
var IDCard string

//go:embed template-surat-keterangan-penyerahan-laporan.html
var SuratKeteranganPenyerahanLaporan string

//go:embed template-rekap-penyerahan-laporan.html
var RekapPenyerahanLaporan string
