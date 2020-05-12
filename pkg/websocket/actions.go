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

	// student actions

	// teacher actions
	io.AddActionHandler("TEACHER_REGISTER", TeacherRegister, []string{})
	io.AddActionHandler("TEACHER_LOGIN", TeacherLogin, []string{})
	io.AddActionHandler("UPDATE_TEACHER_INFORMATION", UpdateTeacherInformation, []string{})

	// academy actions
}
