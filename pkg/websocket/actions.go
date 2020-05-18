package websocket

// InitActions adds all handlers to io router handler
func (io *IO) InitActions() {
	// dev test actions
	io.AddActionHandler("TEST", func(s *Socket, a *Action) {
		s.Emit("TEST_RES", map[string]string{"message": "Mensaje de prueba"})
	}, []string{})

	// public actions
	io.AddActionHandler("USER_REGISTER", UserRegister, []string{})
	io.AddActionHandler("GET_LANGUAGES", GetAllLanguages, []string{})
	io.AddActionHandler("GET_DEGREES", GetAllDegrees, []string{})

	// user actions
	io.AddActionHandler("LOGOUT", Logout, []string{})

	io.AddActionHandler("ADD_USERS_LANGUAGE", UsersAddLanguage, []string{})
	io.AddActionHandler("DELETE_USERS_LANGUAGE", UsersDeleteLanguage, []string{})
	io.AddActionHandler("GET_USERS_LANGUAGES", UsersGetLanguages, []string{})

	io.AddActionHandler("GET_USERS_ACADEMY", UsersGetAcademy, []string{})
	io.AddActionHandler("ADD_USERS_ACADEMY", UsersAddAcademy, []string{})
	io.AddActionHandler("UPDATE_USERS_ACADEMY", UsersUpdateAcademy, []string{})
	io.AddActionHandler("DELETE_USERS_ACADEMY", UsersDeleteAcademy, []string{})

	io.AddActionHandler("GET_USER_SIGNATURES", GetUserTeachingSignatures, []string{})
	io.AddActionHandler("ADD_USER_SIGNATURE", AddUserTeachingSignature, []string{})
	io.AddActionHandler("DELETE_USER_SIGNATURE", DeleteUserTeachingSignature, []string{})
	io.AddActionHandler("UPDATE_USER_SIGNATURE", UpdateUserTeachingSignature, []string{})

	io.AddActionHandler("GET_USER_INSTITUTIONS", GetUserTeachingInstitutions, []string{})
	io.AddActionHandler("ADD_USER_INSTITUTION", AddUserTeachingInstitution, []string{})
	io.AddActionHandler("DELETE_USER_INSTITUTION", DeleteUserTeachingInstitution, []string{})
	io.AddActionHandler("UPDATE_USER_INSTITUTION", UpdateUserTeachingInstitution, []string{})

	io.AddActionHandler("GET_USER_EXPERIENCE", GetUserExperience, []string{})
	io.AddActionHandler("SET_USER_EXPERIENCE", SetUserExperience, []string{})

	io.AddActionHandler("GET_USER_INVESTMENTS", GetUserInvestments, []string{})
	io.AddActionHandler("ADD_USER_INVESTMENT", AddUserInvestment, []string{})
	io.AddActionHandler("DELETE_USER_INVESTMENT", DeleteUserInvestment, []string{})
	io.AddActionHandler("UPDATE_USER_INVESTMENT", UpdateUserInvestment, []string{})

	io.AddActionHandler("GET_USER_MANAGEMENTS", GetUserManagements, []string{})
	io.AddActionHandler("ADD_USER_MANAGEMENT", AddUserManagement, []string{})
	io.AddActionHandler("DELETE_USER_MANAGEMENT", DeleteUserManagement, []string{})
	io.AddActionHandler("UPDATE_USER_MANAGEMENT", UpdateUserManagement, []string{})

	io.AddActionHandler("GET_USER_EXPERTISES", GetUserExpertises, []string{})
	io.AddActionHandler("ADD_USER_EXPERTISE", AddUserExpertise, []string{})
	io.AddActionHandler("DELETE_USER_EXPERTISE", DeleteUserExpertise, []string{})
	io.AddActionHandler("UPDATE_USER_EXPERTISE", UpdateUserExpertise, []string{})

	io.AddActionHandler("GET_USER_BIOGRAPHY", GetUserBiography, []string{})
	io.AddActionHandler("SET_USER_BIOGRAPHY", SetUserBiography, []string{})

	io.AddActionHandler("SET_USER_IMAGE", SetUserImage, []string{})

	// student actions

	// teacher actions
	io.AddActionHandler("TEACHER_REGISTER", TeacherRegister, []string{})
	io.AddActionHandler("TEACHER_LOGIN", TeacherLogin, []string{})
	io.AddActionHandler("UPDATE_TEACHER_INFORMATION", UpdateTeacherInformation, []string{})

	// academy actions
}
