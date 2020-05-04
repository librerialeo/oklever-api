package websocket

// InitActions adds all handlers to io router handler
func (io *IO) InitActions() {
	// dev test actions
	io.AddActionHandler("TEST", func(s *Socket, a *Action) {
		s.Emit("TEST_RES", map[string]string{"message": "Mensaje de prueba"})
	}, []string{})

	// public actions
	io.AddActionHandler("USER_REGISTER", UserRegister, []string{})

	// user actions
	io.AddActionHandler("LOGOUT", Logout, []string{})

	// student actions

	// teacher actions
	io.AddActionHandler("TEACHER_REGISTER", TeacherRegister, []string{})
	io.AddActionHandler("TEACHER_LOGIN", TeacherLogin, []string{})

	// academy actions
}