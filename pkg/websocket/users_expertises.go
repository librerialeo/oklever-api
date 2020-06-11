package websocket

import "github.com/savsgio/atreugo/v11"

// GetUserExpertise Get user expertise
func GetUserExpertise(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		expertiseID, ok := data["id"].(float64)
		if ok {
			expertise, err := s.io.service.GetUserExpertise(int32(expertiseID))
			if err != nil {
				s.EmitServerError("GetUserExpertise", err)
			} else {
				s.Emit(a.Type, expertise.ParseForUser())
			}
		}
	}
}

// GetUserExpertises Get user expertise
func GetUserExpertises(s *Socket, a *Action) {
	if s.user.ID != 0 {
		expertises, err := s.io.service.GetUserExpertises(s.user.ID)
		if err != nil {
			s.EmitServerError("GetUserExpertises", err)
		} else {
			parsedExpertises := []atreugo.JSON{}
			for _, expertise := range *expertises {
				parsedExpertises = append(parsedExpertises, expertise.ParseForUser())
			}
			s.Emit(a.Type, parsedExpertises)
		}
	}
}

// AddUserExpertise Add user expertise
func AddUserExpertise(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.user.ID != 0 {
		name, nOk := data["name"].(string)
		months, mOk := data["months"].(float64)
		if nOk && mOk {
			expertise, err := s.io.service.AddUserExpertise(s.user.ID, name, int16(months))
			if err != nil {
				s.EmitServerError("AddUserExpertise", err)
			} else {
				s.Emit(a.Type, expertise.ParseForUser())
			}
		}
	}
}

// UpdateUserExpertise Update user expertise
func UpdateUserExpertise(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok || s.user.ID == 0 {
		return
	}
	expertiseID, ok := data["id"].(float64)
	name, nOk := data["name"].(string)
	months, mOk := data["months"].(float64)
	if !ok || !mOk || !nOk {
		return
	}
	expertise, err := s.io.service.GetUserExpertise(int32(expertiseID))
	if err != nil || expertise == nil {
		s.EmitServerError("UpdateUserExpertise: expertise not found", err)
		return
	}
	if expertise.UserID.Int != s.user.ID {
		return
	}
	expertise, err = s.io.service.UpdateUserExpertise(int32(expertiseID), name, int16(months))
	if err != nil {
		s.EmitServerError("UpdateUserExpertise: update user error", err)
		return
	}
	s.Emit(a.Type, expertise.ParseForUser())
}

// DeleteUserExpertise Delete user expertise
func DeleteUserExpertise(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok {
		return
	}
	expertiseID, ok := data["id"].(float64)
	if !ok {
		return
	}
	expertise, err := s.io.service.GetUserExpertise(int32(expertiseID))
	if err != nil || expertise == nil {
		s.EmitServerError("DeleteUserExpertise: expertise not found", err)
		return
	}
	if expertise.UserID.Int != s.user.ID {
		return
	}
	err = s.io.service.DeleteUserExpertise(int32(expertiseID))
	if err != nil {
		s.EmitServerError("DeleteUserExpertise: delete error", err)
		return
	}
	s.Emit(a.Type, expertiseID)
}
