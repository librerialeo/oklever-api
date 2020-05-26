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
							"image":     user.Image.String,
							"license":   user.License.String,
							"rfc":       user.RFC.String,
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

// ValidateTeacherProfile validates teacher acception first step
func ValidateTeacherProfile(s *Socket, a *Action) {
	if s.userID == 0 {
		return
	}
	points := 0
	user, err := s.io.service.GetTeacherByUserID(s.userID)
	if err != nil {
		s.EmitServerError("ValidateTeacherProfile: GetTeacherByUserID", err)
		return
	} else if user.TeachingMonths.Int >= 12 && user.TeachingMonths.Int < 36 {
		points++
	} else if user.TeachingMonths.Int >= 36 && user.TeachingMonths.Int < 72 {
		points += 3
	} else if user.TeachingMonths.Int >= 72 {
		points += 5
	}
	academies, err := s.io.service.GetUserAcademiesByUserID(s.userID)
	if err != nil {
		s.EmitServerError("ValidateTeacherProfile: GetUserAcademiesByUserID", err)
		return
	}
	var aux int32 = 0
	for _, academy := range *academies {
		if academy.DegreeID.Int > aux {
			aux = academy.DegreeID.Int
		}
	}
	if aux == 4 || aux == 5 {
		points++
	} else if aux == 2 || aux == 3 {
		points += 3
	} else if aux == 1 {
		points += 5
	}
	signatures, err := s.io.service.GetUserTeachingSignatures(s.userID)
	if err != nil {
		s.EmitServerError("ValidateTeacherProfile: GetUserTeachingSignatures", err)
		return
	}
	aux = 0
	for _, signature := range *signatures {
		if signature.DegreeID.Int > aux {
			aux = signature.DegreeID.Int
		}
	}
	if aux == 4 || aux == 5 {
		points++
	} else if aux == 2 || aux == 3 {
		points += 3
	} else if aux == 1 {
		points += 5
	}
	investments, err := s.io.service.GetUserInvestments(s.userID)
	if err != nil {
		s.EmitServerError("ValidateTeacherProfile: GetUserInvestments", err)
		return
	} else if len(*investments) >= 3 && len(*investments) <= 4 {
		points++
	} else if len(*investments) >= 5 && len(*investments) <= 6 {
		points += 3
	} else if len(*investments) >= 7 {
		points += 5
	}
	managements, err := s.io.service.GetUserManagements(s.userID)
	if err != nil {
		s.EmitServerError("ValidateTeacherProfile: GetUserManagements", err)
		return
	}
	aux = 0
	for _, m := range *managements {
		aux += int32(m.Months.Int)
	}
	if aux >= 24 && aux < 36 {
		points++
	} else if aux >= 36 && aux < 60 {
		points += 3
	} else if aux >= 60 {
		points += 5
	}
	expertises, err := s.io.service.GetUserExpertises(s.userID)
	if err != nil {
		s.EmitServerError("ValidateTeacherProfile: GetUserExpertises", err)
		return
	}
	aux = 0
	for _, m := range *expertises {
		aux += int32(m.Months.Int)
	}
	if aux >= 36 && aux < 48 {
		points++
	} else if aux >= 48 && aux < 72 {
		points += 3
	} else if aux >= 72 {
		points += 5
	}
	if points >= 18 {
		err := s.io.service.SetUserStatus(s.userID, "passant")
		if err != nil {
			s.EmitServerError("ValidateTeacherProfile: SetUserStatus", err)
			return
		}
		s.Emit("SET_USER_STATUS", "passant")
		s.Redirect("/")
		s.EmitSuccess("Tu perfil cumple con los requisitos básicos para ser docente")
	} else {
		s.EmitError("Lo sentimos tu perfil no cumple con los requisitos básicos para ser docente")
	}
}
