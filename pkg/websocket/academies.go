package websocket

import (
	"fmt"
	"time"

	"github.com/librerialeo/oklever-api/pkg/utils"
)

// AcademyRegister register new user
func AcademyRegister(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})

	if ok {
		firstname, firstnameOk := data["firstname"].(string)
		lastname, lastnameOk := data["lastname"].(string)
		email, emailOk := data["email"].(string)
		password, passwordOk := data["password"].(string)
		repassword, repasswordOk := data["repassword"].(string)
		if firstnameOk && lastnameOk && emailOk && passwordOk && repasswordOk && password == repassword {
			hash, err := utils.HashPassword(password)
			if err != nil {
				s.EmitServerError("AcademyRegister: error al encriptar password", err)
			}
			u, err := s.io.service.AddAcademy(firstname, lastname, email, hash)
			if err != nil {
				s.EmitServerError("AcademiesRegister: error al guardar el usuario", err)
			} else {
				token, err := s.io.service.CreateToken(u.ID.Int, u.Rol.Int, u.Status.String, false)
				if err != nil {
					s.EmitServerError("AcademyRegister: error al generar el token", err)
				} else {
					s.JoinRoom("students")
					s.JoinRoom("academies")
					s.SetToken(token)
					s.Emit("ACADEMY_LOGIN", map[string]interface{}{
						"email":     u.Email.String,
						"firstname": u.Firstname.String,
						"lastname":  u.Lastname.String,
					})
				}
			}
		} else {
			s.EmitError("Error en los datos de registro")
		}
	} else {
		s.EmitError("Error en los datos de registro")
	}
}

// AcademyLogin handle the academies login action
func AcademyLogin(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		email, emailOk := data["email"].(string)
		password, passwordOk := data["password"].(string)
		remember, _ := data["remember"].(bool)
		if emailOk && passwordOk {
			user, err := s.io.service.GetAcademyUserByEmail(email)
			if err != nil {
				fmt.Println(err)
				switch err.Error() {
				case "duplicated email":
				case "email not found":
					s.EmitError("Email o contraseña incorrecta")
				default:
					s.EmitServerError("AcademyLogin: get user by email", err)
				}
			} else {
				if utils.CheckPasswordHash(password, user.Password.String) {
					token, err := s.io.service.CreateToken(user.ID.Int, user.Rol.Int, user.Status.String, remember)
					if err != nil {
						s.EmitServerError("AcademyLogin: generate token", err)
					} else {
						s.io.service.UpdateUserLastAction(user.ID.Int, time.Now())
						s.SetToken(token)
						s.Emit("ACADEMY_LOGIN", map[string]interface{}{
							"email":     user.Email.String,
							"firstname": user.Firstname.String,
							"lastname":  user.Lastname.String,
							"gender":    user.Gender.String,
							"phone":     user.Phone.String,
							"image":     user.Image.String,
							"license":   user.License.String,
							"rfc":       user.RFC.String,
							"rol":       user.Rol.Int,
							"status":    user.Status.String,
						})
						s.EmitSuccess("Iniciaste sesión correctamente")
						idsLanguages, idsErr := s.io.service.GetAllUsersLanguages(user.ID.Int)
						if idsErr != nil {
							s.EmitServerError("Users get languages", idsErr)
						} else {
							s.Emit("ADD_USERS_LANGUAGE", idsLanguages)
						}
						academies, aErr := s.io.service.GetUserAcademiesByUserID(user.ID.Int)
						if aErr != nil {
							s.EmitServerError("Users get academies", aErr)
						} else {
							s.Emit("ADD_USERS_ACADEMY", academies)
						}
					}
				} else {
					s.EmitError("Email o contraseña incorrecta")
				}
			}
		}
	}
}
