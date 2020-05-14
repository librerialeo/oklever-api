package websocket

import (
	"time"

	"github.com/librerialeo/oklever-api/pkg/utils"
)

// TeacherRegister register new user
func TeacherRegister(s *Socket, a *Action) {
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
				s.EmitServerError("TeacherRegister: error al encriptar password", err)
			}
			u, err := s.io.service.AddTeacher(firstname, lastname, email, hash)
			if err != nil {
				s.EmitServerError("TeachersRegister: error al guardar el usuario", err)
			} else {
				token, err := s.io.service.CreateToken(u.ID.Int, u.Rol.Int, false)
				if err != nil {
					s.EmitServerError("TeacherRegister: error al generar el token", err)
				} else {
					s.JoinRoom("students")
					s.JoinRoom("teachers")
					s.SetToken(token)
					s.Emit("TEACHER_LOGIN", map[string]interface{}{
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

// TeacherLogin handle the teachers login action
func TeacherLogin(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		email, emailOk := data["email"].(string)
		password, passwordOk := data["password"].(string)
		remember, _ := data["remember"].(bool)
		if emailOk && passwordOk {
			user, err := s.io.service.GetTeacherUserByEmail(email)
			if err != nil {
				switch err.Error() {
				case "duplicated email":
				case "email not found":
					s.EmitError("Email o contraseña incorrecta")
				default:
					s.EmitServerError("TeacherLogin: get user by email", err)
				}
			} else {
				if utils.CheckPasswordHash(password, user.Password.String) {
					token, err := s.io.service.CreateToken(user.ID.Int, user.Rol.Int, remember)
					if err != nil {
						s.EmitServerError("TeacherLogin: generate token", err)
					} else {
						s.io.service.UpdateUserLastAction(user.ID.Int, time.Now())
						s.SetToken(token)
						s.Emit("TEACHER_LOGIN", map[string]interface{}{
							"email":     user.Email.String,
							"firstname": user.Firstname.String,
							"lastname":  user.Lastname.String,
							"gender":    user.Gender.String,
							"phone":     user.Phone.String,
							"license":   user.License.String,
							"rfc":       user.RFC.String,
						})
						s.EmitSuccess("Iniciaste sesión correctamente")
						idsLanguages, idsErr := s.io.service.GetAllUsersLanguages(user.ID.Int)
						if idsErr != nil {
							s.EmitServerError("Users get languages", idsErr)
						} else {
							s.Emit("ADD_USERS_LANGUAGE", idsLanguages)
						}
						academies, aErr := s.io.service.GetAllUsersAcademy(user.ID.Int)
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

// UpdateTeacherInformation update teachers user data
func UpdateTeacherInformation(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		email, emailOk := data["email"].(string)
		firstname, firstnameOk := data["firstname"].(string)
		lastname, lastnameOk := data["lastname"].(string)
		gender, genderOk := data["gender"].(string)
		phone, phoneOk := data["phone"].(string)
		license, licenseOk := data["license"].(string)
		rfc, rfcOk := data["rfc"].(string)
		if emailOk && firstnameOk && lastnameOk && genderOk && phoneOk && licenseOk && rfcOk && s.userID != 0 {
			if err := s.io.service.UpdateUserInformation(s.userID, firstname, lastname, email, gender, phone); err != nil {
				s.EmitServerError("UpdateTeacherInformation: update user information", err)
				return
			}
			if err := s.io.service.UpdateTeacherInformation(s.userID, license, rfc); err != nil {
				s.EmitServerError("UpdateTeacherInformation: update teacher information", err)
				return
			}
			s.Emit("UPDATE_TEACHER_INFORMATION", map[string]interface{}{
				"email":     email,
				"firstname": firstname,
				"lastname":  lastname,
				"gender":    gender,
				"phone":     phone,
				"license":   license,
				"rfc":       rfc,
			})
		}
	}
}
