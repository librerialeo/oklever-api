package websocket

import (
	"fmt"
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
					s.Emit("TEACHER_LOGIN", u)
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
						s.Emit("TEACHER_LOGIN", user)
						s.EmitSuccess("Iniciaste sesión correctamente")
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
	fmt.Println(a)
}
