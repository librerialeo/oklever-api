package websocket

import (
	"fmt"

	"github.com/librerialeo/oklever-api/pkg/database"
	"github.com/librerialeo/oklever-api/pkg/utils"
)

// TeacherRegister register new user
func TeacherRegister(s *Socket, a *Action) {
	fmt.Println("action", a)
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
				s.EmitServerError("Teacher Register hash generation error", err)
			}
			fmt.Println(firstname, lastname, email, hash)
			rows, err := s.io.service.AddTeacher(firstname, lastname, email, hash)
			if err != nil {
				s.EmitServerError("TeachersRegister: error al guardar el usuario", err)
			} else {
				defer rows.Close()
				var u database.DBUser
				if rows.Next() {
					err = rows.Scan(&u.ID, &u.Email, &u.Firstname, &u.Lastname, &u.Rol, &u.Created)
					if err != nil {
						s.EmitServerError("u", err)
					} else {
						fmt.Println("u", u)
					}
				}
				fmt.Println("rows", rows)
				s.EmitMessage("Cuenta creada correctamente", "success")
				s.Emit("LOGIN", map[string]interface{}{
					"id":        u.ID.Get(),
					"email":     u.Email.Get(),
					"firstname": u.Firstname.Get(),
					"lastname":  u.Lastname.Get(),
					"rol":       u.Rol.Get(),
					"created":   u.Created.Get(),
				})
			}
		} else {
			s.EmitError("Error en los datos de registro")
		}
	} else {
		s.EmitError("Error en los datos de registro")
	}
}
