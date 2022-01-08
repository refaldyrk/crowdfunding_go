package user

type UserVerificationFormatter struct {
	Verif   int    `json:"verif"`
	Message string `json:"message"`
}

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Code       string `json:"code"`
	ImageUrl   string `json:"image_url"`
}

func FormatVerif(user User) UserVerificationFormatter {
	format := UserVerificationFormatter{
		Verif:   user.Verif,
		Message: "Verifikasi Berhasil",
	}
	return format
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		Code:       user.Code,
		ImageUrl:   user.AvatarFileName,
	}

	return formatter
}
