package Models

import "fmt"

type EmailMessage struct {
	To      string
	Subject string
	Body    string
}

func ToSendOTPMessage(to, name, otp string) *EmailMessage {
	body := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Send OTP</title>
</head>
<body>
    <h1>Halo %s,</h1>
    <p>Berikut adalah Kode OTP (One-Time Password) Anda untuk verifikasi:</p>
    <h2 style="color: blue;">%s</h2>
    <p>Kode ini hanya berlaku selama <b>60 detik</b> dan hanya dapat digunakan sekali. Jangan bagikan kode ini kepada siapapun, termasuk kepada pihak yang mengaku dari DuitAja.</p>
    <p>Jika Anda tidak meminta kode ini, harap abaikan email ini.</p>
    <p>Terima kasih,</p>
    <p><b>Tim DuitAja</b></p>
</body>
</html>`, name, otp)

	return &EmailMessage{
		To:      to,
		Subject: "Kode OTP Anda - Jangab Bagikan",
		Body:    body,
	}
}

func ToSendEmailVerificationMessage(to, name, token string) *EmailMessage {
	body := fmt.Sprintf(`<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Send Verification Account</title>
		</head>
		<body>
			<h1>Halo %s </h1>
			<p>Berikut Adalah Kode OTP (One-Time Padssword) Anda untuk verifikasi</p>
	
			<p>Terima kasih telah mendaftar di DuitAja! Untuk mengaktifkan akun Anda, <b><a href="">silakan verifikasi email Anda %s.</a></b></p>
	
			<p>Jika anda tidak merasa mendaftar, harap abaikan email ini</p>
	
			<p>Terima Kasih,</p>
			<p>Tim DuitAja</p>
		</body>
		</html>`, name, token)

	return &EmailMessage{
		To:      to,
		Subject: "Verifikasi Email Anda - Selangkah Lagi Menuju DuitAja",
		Body:    body,
	}
}
