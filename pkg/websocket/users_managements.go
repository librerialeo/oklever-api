package websocket

import "github.com/savsgio/atreugo"

// GetUserManagement Get user management
func GetUserManagement(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		managementID, ok := data["id"].(float64)
		if ok {
			management, err := s.io.service.GetUserManagement(int32(managementID))
			if err != nil {
				s.EmitServerError("GetUserManagement", err)
			} else {
				s.Emit(a.Type, management.ParseForUser())
			}
		}
	}
}

// GetUserManagements Get user management
func GetUserManagements(s *Socket, a *Action) {
	if s.user.ID != 0 {
		managements, err := s.io.service.GetUserManagements(s.user.ID)
		if err != nil {
			s.EmitServerError("GetUserManagements", err)
		} else {
			parsedManagements := []atreugo.JSON{}
			for _, management := range *managements {
				parsedManagements = append(parsedManagements, management.ParseForUser())
			}
			s.Emit(a.Type, parsedManagements)
		}
	}
}

// AddUserManagement Add user management
func AddUserManagement(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.user.ID != 0 {
		job, jOk := data["job"].(string)
		institution, iOk := data["institution"].(string)
		months, mOk := data["months"].(float64)
		if jOk && iOk && mOk {
			management, err := s.io.service.AddUserManagement(s.user.ID, job, institution, int16(months))
			if err != nil {
				s.EmitServerError("AddUserManagement", err)
			} else {
				s.Emit(a.Type, management.ParseForUser())
			}
		}
	}
}

// UpdateUserManagement Update user management
func UpdateUserManagement(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok || s.user.ID == 0 {
		return
	}
	managementID, ok := data["id"].(float64)
	job, jOk := data["job"].(string)
	institution, rOk := data["institution"].(string)
	months, mOk := data["months"].(float64)
	if !ok || !rOk || !mOk || !jOk {
		return
	}
	management, err := s.io.service.GetUserManagement(int32(managementID))
	if err != nil || management == nil {
		s.EmitServerError("UpdateUserManagement: management not found", err)
		return
	}
	if management.UserID.Int != s.user.ID {
		return
	}
	management, err = s.io.service.UpdateUserManagement(int32(managementID), job, institution, int16(months))
	if err != nil {
		s.EmitServerError("UpdateUserManagement: update user error", err)
		return
	}
	s.Emit(a.Type, management.ParseForUser())
}

// DeleteUserManagement Delete user management
func DeleteUserManagement(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok {
		return
	}
	managementID, ok := data["id"].(float64)
	if !ok {
		return
	}
	management, err := s.io.service.GetUserManagement(int32(managementID))
	if err != nil || management == nil {
		s.EmitServerError("DeleteUserManagement: management not found", err)
		return
	}
	if management.UserID.Int != s.user.ID {
		return
	}
	err = s.io.service.DeleteUserManagement(int32(managementID))
	if err != nil {
		s.EmitServerError("DeleteUserManagement: delete error", err)
		return
	}
	s.Emit(a.Type, managementID)
}
