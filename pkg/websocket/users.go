package websocket

import (
	"fmt"
	"os"

	"github.com/librerialeo/oklever-api/pkg/utils"
	"github.com/savsgio/atreugo/v11"
)

// UserRegister register new user
func UserRegister(socket *Socket, action *Action) {
}

// Logout desubscribe socket from rooms and send logout action
func Logout(s *Socket, a *Action) {
	for _, r := range s.rooms {
		s.LeaveRoom(r)
	}
	s.SetToken("")
	s.Emit("LOGOUT")
}

// GetUserBiography send user biography
func GetUserBiography(s *Socket, a *Action) {
	biography, err := s.io.service.GetUserBiography(s.user.ID)
	if err != nil {
		s.EmitServerError("GetUserBiography", err)
	} else {
		s.Emit(a.Type, atreugo.JSON{
			"biography": biography,
		})
	}
}

// SetUserBiography send user biography
func SetUserBiography(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		biography, ok := data["biography"]
		if ok {
			DBbiography, err := s.io.service.SetUserBiography(s.user.ID, biography.(string))
			if err != nil {
				s.EmitServerError("GetUserBiography", err)
			} else {
				s.Emit(a.Type, atreugo.JSON{
					"biography": DBbiography,
				})
			}
		}
	}
}

// SetUserImage set user image
func SetUserImage(s *Socket, a *Action) {
	data, ok := a.Data.(string)
	if !ok || s.user.ID == 0 {
		return
	}
	image, err := utils.DataToImage(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := fmt.Sprintf("%d-image.png", s.user.ID)
	err = utils.SaveImageToFile(os.Getenv("USERS_DIR"), filename, image, "image/png")
	if err != nil {
		fmt.Println(err)
		return
	}
	filename, err = s.io.service.SetUserImage(s.user.ID, filename)
	if err != nil {
		return
	}
	s.Emit(a.Type, filename)
}
