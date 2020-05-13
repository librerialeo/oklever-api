package websocket

import "fmt"

//GetAllDegrees get all degrees
func GetAllDegrees(s *Socket, a *Action) {
	degrees, err := s.io.service.GetAllDegrees()
	if err != nil {
		s.EmitServerError("Get degrees", err)
	} else {
		fmt.Println(degrees)
		s.Emit("ADD_DEGREES", degrees)
	}
}
