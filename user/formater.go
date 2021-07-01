package user

type UserFormater struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUser(user User) UserFormater {
	formater := UserFormater{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return formater
}
