package websocket

// InitActions adds all handlers to io router handler
func (io *IO) InitActions() {
	// dev test actions
	io.AddActionHandler("TEST", func(s *Socket, a *Action) {
		s.Emit("TEST_RES", map[string]string{"message": "Mensaje de prueba"})
	})

	// public actions
	io.AddActionHandler("USER_REGISTER", UserRegister)
	io.AddActionHandler("GET_LANGUAGES", GetAllLanguages)
	io.AddActionHandler("GET_DEGREES", GetAllDegrees)
	io.AddActionHandler("GET_QUIZ", GetQuizByID)

	// user actions
	io.AddActionHandler("LOGOUT", Logout)

	io.AddActionHandler("ADD_USERS_LANGUAGE", UsersAddLanguage)
	io.AddActionHandler("DELETE_USERS_LANGUAGE", UsersDeleteLanguage)
	io.AddActionHandler("GET_USERS_LANGUAGES", UsersGetLanguages)

	io.AddActionHandler("GET_USERS_ACADEMY", UsersGetAcademy)
	io.AddActionHandler("ADD_USERS_ACADEMY", UsersAddAcademy)
	io.AddActionHandler("UPDATE_USERS_ACADEMY", UsersUpdateAcademy)
	io.AddActionHandler("DELETE_USERS_ACADEMY", UsersDeleteAcademy)

	io.AddActionHandler("GET_USER_SIGNATURES", GetUserTeachingSignatures)
	io.AddActionHandler("ADD_USER_SIGNATURE", AddUserTeachingSignature)
	io.AddActionHandler("DELETE_USER_SIGNATURE", DeleteUserTeachingSignature)
	io.AddActionHandler("UPDATE_USER_SIGNATURE", UpdateUserTeachingSignature)

	io.AddActionHandler("GET_USER_INSTITUTIONS", GetUserTeachingInstitutions)
	io.AddActionHandler("ADD_USER_INSTITUTION", AddUserTeachingInstitution)
	io.AddActionHandler("DELETE_USER_INSTITUTION", DeleteUserTeachingInstitution)
	io.AddActionHandler("UPDATE_USER_INSTITUTION", UpdateUserTeachingInstitution)

	io.AddActionHandler("GET_USER_EXPERIENCE", GetUserExperience)
	io.AddActionHandler("SET_USER_EXPERIENCE", SetUserExperience)

	io.AddActionHandler("GET_USER_INVESTMENTS", GetUserInvestments)
	io.AddActionHandler("ADD_USER_INVESTMENT", AddUserInvestment)
	io.AddActionHandler("DELETE_USER_INVESTMENT", DeleteUserInvestment)
	io.AddActionHandler("UPDATE_USER_INVESTMENT", UpdateUserInvestment)

	io.AddActionHandler("GET_USER_MANAGEMENTS", GetUserManagements)
	io.AddActionHandler("ADD_USER_MANAGEMENT", AddUserManagement)
	io.AddActionHandler("DELETE_USER_MANAGEMENT", DeleteUserManagement)
	io.AddActionHandler("UPDATE_USER_MANAGEMENT", UpdateUserManagement)

	io.AddActionHandler("GET_USER_EXPERTISES", GetUserExpertises)
	io.AddActionHandler("ADD_USER_EXPERTISE", AddUserExpertise)
	io.AddActionHandler("DELETE_USER_EXPERTISE", DeleteUserExpertise)
	io.AddActionHandler("UPDATE_USER_EXPERTISE", UpdateUserExpertise)

	io.AddActionHandler("GET_USER_BIOGRAPHY", GetUserBiography)
	io.AddActionHandler("SET_USER_BIOGRAPHY", SetUserBiography)

	io.AddActionHandler("SET_USER_IMAGE", SetUserImage)

	// student actions

	// teacher actions
	io.AddActionHandler("TEACHER_REGISTER", TeacherRegister)
	io.AddActionHandler("TEACHER_LOGIN", TeacherLogin)

	io.AddActionHandler("UPDATE_USER_INFORMATION", UpdateUserInformation)
	io.AddActionHandler("SET_TEACHER_TEST_CLASS", AddTeachersTestClass)
	io.AddActionHandler("VALIDATE_TEACHER_PROFILE", ValidateTeacherProfile)

	// academy actions
	io.AddActionHandler("ACADEMY_LOGIN", AcademyLogin)
	io.AddActionHandler("ACADEMY_REGISTER", AcademyRegister)
}
