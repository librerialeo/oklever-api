package websocket

import (
	"fmt"
	"os"

	"github.com/librerialeo/oklever-api/pkg/database"
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

// UpdateUserInformation update teachers user data
func UpdateUserInformation(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		email, emailOk := data["email"].(string)
		firstname, firstnameOk := data["firstname"].(string)
		lastname, lastnameOk := data["lastname"].(string)
		gender, genderOk := data["gender"].(string)
		phone, phoneOk := data["phone"].(string)
		if emailOk && firstnameOk && lastnameOk && genderOk && phoneOk && s.user.ID != 0 {
			u, err := s.io.service.UpdateUserInformation(s.user.ID, firstname, lastname, email, gender, phone)
			if err != nil {
				s.EmitServerError("UpdateUserInformation: update user information", err)
				return
			}
			if u.Rol.Int == 2 {
				license, licenseOk := data["license"].(string)
				rfc, rfcOk := data["rfc"].(string)
				if licenseOk && rfcOk {
					var u2 *database.DBUser
					if u.Rol.Int == 2 {
						u2, err = s.io.service.UpdateTeacherInformation(s.user.ID, license, rfc)
						if err != nil {
							s.EmitServerError("UpdateUserInformation: update teacher information", err)
							return
						}
					} else {
						u2, err = s.io.service.UpdateAcademyInformation(s.user.ID, license, rfc)
						if err != nil {
							s.EmitServerError("UpdateUserInformation: update academy information", err)
							return
						}
					}
					if u2.License.String == "" || u2.RFC.String == "" {
						s.EmitError("Licencia y/o RFC duplicados, no se actualizaron todos los datos")
						s.Emit("UPDATE_USER_INFORMATION", map[string]interface{}{
							"email":     u.Email.String,
							"firstname": u.Firstname.String,
							"lastname":  u.Lastname.String,
							"gender":    u.Gender.String,
							"phone":     u.Phone.String,
							"license":   u.License.String,
							"rfc":       u.RFC.String,
						})
					} else {
						s.Emit("UPDATE_USER_INFORMATION", map[string]interface{}{
							"email":     u.Email.String,
							"firstname": u.Firstname.String,
							"lastname":  u.Lastname.String,
							"gender":    u.Gender.String,
							"phone":     u.Phone.String,
							"license":   u2.License.String,
							"rfc":       u2.RFC.String,
						})
					}
				} else {
					s.Emit("UPDATE_USER_INFORMATION", map[string]interface{}{
						"email":     u.Email.String,
						"firstname": u.Firstname.String,
						"lastname":  u.Lastname.String,
						"gender":    u.Gender.String,
						"phone":     u.Phone.String,
					})
				}
			}
		}
	}
}
