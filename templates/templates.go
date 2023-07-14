package templates

import _ "embed"

//go:embed template-forgot-password.html
var ForgotPassword string

//go:embed template-verified-register.html
var VerifiedRegister string

//go:embed template-denied-register.html
var DeniedRegister string

//go:embed template-verified-repository.html
var VerifiedRepository string

//go:embed template-denied-repository.html
var DeniedRepository string

//go:embed template-id-card.html
var IDCard string

//go:embed template-surat-keterangan-penyerahan-laporan.html
var SuratKeteranganPenyerahanLaporan string

//go:embed template-rekap-penyerahan-laporan.html
var RekapPenyerahanLaporan string
