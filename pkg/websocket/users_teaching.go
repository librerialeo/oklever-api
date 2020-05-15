package websocket

// GetUserTeachingSignature Get user signature
func GetUserTeachingSignature(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		signatureID, ok := data["id"].(float64)
		if ok {
			signature, err := s.io.service.GetUserTeachingSignature(int32(signatureID))
			if err != nil {
				s.EmitServerError("GetUserTeachingSignature", err)
			} else {
				s.Emit(a.Type, signature)
			}
		}
	}
}

// GetUserTeachingSignatures Get user signature
func GetUserTeachingSignatures(s *Socket, a *Action) {
	if s.userID != 0 {
		signatures, err := s.io.service.GetUserTeachingSignatures(s.userID)
		if err != nil {
			s.EmitServerError("GetUserTeachingSignatures", err)
		} else {
			s.Emit(a.Type, *signatures)
		}
	}
}

// AddUserTeachingSignature Add user signature
func AddUserTeachingSignature(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		degreeID, ok := data["degree"].(float64)
		name, ok := data["name"].(string)
		if ok {
			signature, err := s.io.service.AddUserTeachingSignature(s.userID, int32(degreeID), name)
			if err != nil {
				s.EmitServerError("AddUserTeachingSignature", err)
			} else {
				s.Emit(a.Type, signature)
			}
		}
	}
}

// UpdateUserTeachingSignature Update user signature
func UpdateUserTeachingSignature(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok || s.userID == 0 {
		return
	}
	signatureID, ok := data["id"].(float64)
	name, ok := data["name"].(string)
	if !ok {
		return
	}
	signature, err := s.io.service.GetUserTeachingSignature(int32(signatureID))
	if err != nil || signature == nil {
		s.EmitServerError("UpdateUserTeachingSignature: signature not found", err)
		return
	}
	if signature.UserID.Int != s.userID {
		return
	}
	signature, err = s.io.service.UpdateUserTeachingSignature(int32(signatureID), name)
	if err != nil {
		s.EmitServerError("UpdateUserTeachingSignature: update user error", err)
		return
	}
	s.Emit(a.Type, signature)
}

// DeleteUserTeachingSignature Delete user signature
func DeleteUserTeachingSignature(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok {
		return
	}
	signatureID, ok := data["id"].(float64)
	if !ok {
		return
	}
	signature, err := s.io.service.GetUserTeachingSignature(int32(signatureID))
	if err != nil || signature == nil {
		s.EmitServerError("DeleteUserTeachingSignature: signature not found", err)
		return
	}
	if signature.UserID.Int != s.userID {
		return
	}
	err = s.io.service.DeleteUserTeachingSignature(int32(signatureID))
	if err != nil {
		s.EmitServerError("DeleteUserTeachingSignature: delete error", err)
		return
	}
	s.Emit(a.Type, signatureID)
}

// GetUserTeachingInstitution Get user institution
func GetUserTeachingInstitution(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		institutionID, ok := data["id"].(float64)
		if ok {
			institution, err := s.io.service.GetUserTeachingInstitution(int32(institutionID))
			if err != nil {
				s.EmitServerError("GetUserTeachingInstitution", err)
			} else {
				s.Emit(a.Type, institution)
			}
		}
	}
}

// GetUserTeachingInstitutions Get user institution
func GetUserTeachingInstitutions(s *Socket, a *Action) {
	if s.userID != 0 {
		institutions, err := s.io.service.GetUserTeachingInstitutions(s.userID)
		if err != nil {
			s.EmitServerError("GetUserTeachingInstitutions", err)
		} else {
			s.Emit(a.Type, *institutions)
		}
	}
}

// AddUserTeachingInstitution Add user institution
func AddUserTeachingInstitution(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		name, ok := data["name"].(string)
		if ok {
			institution, err := s.io.service.AddUserTeachingInstitution(s.userID, name)
			if err != nil {
				s.EmitServerError("AddUserTeachingInstitution", err)
			} else {
				s.Emit(a.Type, institution)
			}
		}
	}
}

// UpdateUserTeachingInstitution Update user institution
func UpdateUserTeachingInstitution(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok || s.userID == 0 {
		return
	}
	institutionID, ok := data["id"].(float64)
	name, ok := data["name"].(string)
	if !ok {
		return
	}
	institution, err := s.io.service.GetUserTeachingInstitution(int32(institutionID))
	if err != nil || institution == nil {
		s.EmitServerError("UpdateUserTeachingInstitution: institution not found", err)
		return
	}
	if institution.UserID.Int != s.userID {
		return
	}
	institution, err = s.io.service.UpdateUserTeachingInstitution(int32(institutionID), name)
	if err != nil {
		s.EmitServerError("UpdateUserTeachingInstitution: update user error", err)
		return
	}
	s.Emit(a.Type, institution)
}

// DeleteUserTeachingInstitution Delete user institution
func DeleteUserTeachingInstitution(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok {
		return
	}
	institutionID, ok := data["id"].(float64)
	if !ok {
		return
	}
	institution, err := s.io.service.GetUserTeachingInstitution(int32(institutionID))
	if err != nil || institution == nil {
		s.EmitServerError("DeleteUserTeachingInstitution: institution not found", err)
		return
	}
	if institution.UserID.Int != s.userID {
		return
	}
	err = s.io.service.DeleteUserTeachingInstitution(int32(institutionID))
	if err != nil {
		s.EmitServerError("DeleteUserTeachingInstitution: delete error", err)
		return
	}
	s.Emit(a.Type, institutionID)
}

// GetUserExperience get user teaching months
func GetUserExperience(s *Socket, a *Action) {
	if s.userID != 0 {
		experience, err := s.io.service.GetUserExperience(s.userID)
		if err != nil {
			s.EmitServerError("GetUserExperience", err)
		} else {
			s.Emit(a.Type, *experience)
		}
	}
}

// SetUserExperience set user teaching months
func SetUserExperience(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		months, ok := data["months"].(float64)
		if ok {
			experience, err := s.io.service.SetUserExperience(s.userID, int32(months))
			if err != nil {
				s.EmitServerError("SetUserExperience", err)
			} else {
				s.Emit(a.Type, *experience)
			}
		}
	}
}
