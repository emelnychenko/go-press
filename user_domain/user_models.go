package user_domain

type (
	User struct {
		Id *UserId `json:"id" xml:"id"`
		FirstName string `json:"firstName" xml:"firstName"`
		LastName string `json:"lastName" xml:"lastName"`
		Email string `json:"email" xml:"email"`
		Verified bool `json:"verified" xml:"verified"`
	}

	UserCreate struct {
		FirstName string `json:"firstName" xml:"firstName"`
		LastName string `json:"lastName" xml:"lastName"`
		Email string `json:"email" xml:"email"`
		Password string `json:"password" xml:"password"`
	}

	UserUpdate struct {
		FirstName string `json:"firstName" xml:"firstName"`
		LastName string `json:"lastName" xml:"lastName"`
	}

	UserChangePassword struct {
		OldPassword string `json:"oldPassword" xml:"oldPassword"`
		NewPassword string `json:"newPassword" xml:"newPassword"`
	}

	UserChangeIdentity struct {
		Email string `json:"email" xml:"email"`
	}
)
