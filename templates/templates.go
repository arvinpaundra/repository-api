package templates

import _ "embed"

//go:embed template-forgot-password.html
var ForgotPassword string

//go:embed template-verified-register.html
var VerifiedRegister string
