package mailing

import (
	"fmt"
	"strconv"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/templates"
	"gopkg.in/gomail.v2"
)

type Mailing struct {
	host        string
	port        string
	senderName  string
	email       string
	password    string
	frontendURL string
}

type User struct {
	Fullname       string
	Email          string
	IdentityNumber string
	Departement    string
	StudyProgram   string
	Role           string
	YearGen        string
}

type Repository struct {
	Title         string
	Authors       string
	Collection    string
	Category      string
	Departement   string
	DateValidated string
}

// NewMailing returns a new Mailing struct
func NewMailing(host, port, email, password, senderName, frontendURL string) *Mailing {
	return &Mailing{
		host:        host,
		port:        port,
		email:       email,
		password:    password,
		senderName:  senderName,
		frontendURL: frontendURL,
	}
}

func (mc *Mailing) send(email, subject, message string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", mc.senderName)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", message)

	port, _ := strconv.Atoi(mc.port)

	dialer := gomail.NewDialer(mc.host, port, mc.email, mc.password)

	if err := dialer.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}

func (mc *Mailing) SendForgotPasswordMail(recipient, subject, token string) error {
	data := map[string]interface{}{
		"token":    token,
		"email":    recipient,
		"base_url": mc.frontendURL,
	}

	result, err := helper.RenderHTMLToString(templates.ForgotPassword, nil, data)

	if err != nil {
		return err
	}

	if err := mc.send(recipient, subject, result); err != nil {
		return err
	}

	return nil
}

func (mc *Mailing) SendVerifiedRegisterMail(subject string, user User) error {
	var message string

	if user.Role == "Mahasiswa" {
		message = fmt.Sprintf("Hai, %s<br><br>Akun kamu <b>berhasil terverifikasi</b> dan sekarang menjadi bagian dari REKSI PNC. Kami ucapkan selamat datang di dunia pengetahuan tak terbatas!<br><br>Berikut adalah data yang kamu berikan :<br>Nama Lengkap : <b>%s</b><br>NIM : <b>%s</b><br>Jurusan : <b>%s</b><br>Program Studi : <b>%s</b><br>Alamat Email : <b>%s</b><br>Role : <b>%s</b><br>Tahun Angkatan : <b>%s</b><br><br>Salam hangat,<br>REKSI PNC", user.Fullname, user.Fullname, user.IdentityNumber, user.Departement, user.StudyProgram, user.Email, user.Role, user.YearGen)
	} else {
		message = fmt.Sprintf("Hai, %s<br><br>Akun kamu <b>berhasil terverifikasi</b> dan sekarang menjadi bagian dari REKSI PNC. Kami ucapkan selamat datang di dunia pengetahuan tak terbatas!<br><br>Berikut adalah data yang kamu berikan :<br>Nama Lengkap : <b>%s</b><br>NIM : <b>%s</b><br>Jurusan : <b>%s</b><br>Program Studi : <b>%s</b><br>Alamat Email : <b>%s</b><br>Role : <b>%s</b><br><br>Salam hangat,<br>REKSI PNC", user.Fullname, user.Fullname, user.IdentityNumber, user.Departement, user.StudyProgram, user.Email, user.Role)
	}

	if err := mc.send(user.Email, subject, message); err != nil {
		return err
	}

	return nil
}

func (mc *Mailing) SendDeniedRegisterMail(subject, reasons string, user User) error {
	message := fmt.Sprintf("Hai, %s<br><br>Dengan menyesal, kami ingin memberitahukan bahwa registrasi Anda telah <b>ditolak</b>. Setelah melakukan peninjauan, kami menemukan ketidaksesuaian dengan informasi yang Anda berikan.<br><br>Berikut adalah alasan penolakan registrasi Anda :<br>%s<br><br>Jika terdapat kekeliruan silahkan segera menghubungi pihak perpustakaan.<br><br>Salam hangat,<br>REKSI PNC", user.Fullname, reasons)

	if err := mc.send(user.Email, subject, message); err != nil {
		return err
	}

	return nil
}

func (mc *Mailing) SendVerifiedRepositoryMail(subject string, user User, repository Repository) error {
	message := fmt.Sprintf("Hai, %s<br><br>Horee, karya tulis ilmiah yang Anda unggah telah <b>disetujui</b>. Kami sangat menghargai kontribusi dan semangat Anda dalam berbagi pengetahuan.<br><br>Berikut adalah detail karya tulis ilmiah yang telah Anda unggah :<br>Judul : <b>%s</b><br>Penulis : <b>%s</b><br>Koleksi : <b>%s</b><br>Kategori : <b>%s</b><br>Jurusan : <b>%s</b><br>Tanggal Divalidasi : <b>%s</b><br><br>Salam hangat,<br>REKSI PNC", user.Fullname, repository.Title, repository.Authors, repository.Collection, repository.Category, repository.Departement, repository.DateValidated)

	if err := mc.send(user.Email, subject, message); err != nil {
		return err
	}

	return nil
}

func (mc *Mailing) SendDeniedRepositoryMail(subject, reasons string, user User, repository Repository) error {
	message := fmt.Sprintf("Hai, %s<br><br>Email ini ditujukan untuk memberitahukan bahwa unggahan karya tulis ilmiah Anda dengan judul <b>%s</b> telah <b>ditolak</b>. Setelah melakukan peninjauan, kami menemukan ketidaksesuaian pada karya tulis ilmiah yang Anda unggah.<br><br>Berikut adalah beberapa alasan penolakan unggahan karya tulis ilmiah Anda :<br>%s<br><br>Jika terdapat kekeliruan silahkan segera menghubungi pihak perpustakaan.<br><br>Salam hangat,<br>REKSI PNC<br>", user.Fullname, repository.Title, reasons)

	if err := mc.send(user.Email, subject, message); err != nil {
		return err
	}

	return nil
}
