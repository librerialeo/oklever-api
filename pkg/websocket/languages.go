package websocket

//GetAllLanguages get all languages
func GetAllLanguages(s *Socket, a *Action) {
	languages, err := s.io.service.GetAllLanguages()
	if err != nil {
		s.EmitServerError("Get languages", err)
	} else {
		s.Emit("ADD_LANGUAGES", languages)
	}
}
