package websocket

import (
	"github.com/savsgio/atreugo"
)

// GetUserInvestment Get user investment
func GetUserInvestment(s *Socket, a *Action) {
	data, ok := a.Data.(atreugo.JSON)
	if ok {
		investmentID, ok := data["id"]
		if ok {
			investment, err := s.io.service.GetUserInvestment(int32(investmentID.(float64)))
			if err != nil {
				s.EmitServerError("GetUserInvestment", err)
			} else {
				s.Emit(a.Type, investment)
			}
		}
	}
}

// GetUserInvestments Get user investment
func GetUserInvestments(s *Socket, a *Action) {
	if s.userID != 0 {
		investments, err := s.io.service.GetUserInvestments(s.userID)
		if err != nil {
			s.EmitServerError("GetUserInvestments", err)
		} else {
			s.Emit(a.Type, *investments)
		}
	}
}

// AddUserInvestment Add user investment
func AddUserInvestment(s *Socket, a *Action) {
	data, ok := a.Data.(atreugo.JSON)
	if ok && s.userID != 0 {
		typeof, tOk := data["type"]
		reference, nOk := data["reference"]
		year, yOk := data["year"]
		if tOk && nOk && yOk {
			investment, err := s.io.service.AddUserInvestment(s.userID, typeof.(string), reference.(string), int32(year.(float64)))
			if err != nil {
				s.EmitServerError("AddUserInvestment", err)
			} else {
				s.Emit(a.Type, investment)
			}
		}
	}
}

// UpdateUserInvestment Update user investment
func UpdateUserInvestment(s *Socket, a *Action) {
	data, ok := a.Data.(atreugo.JSON)
	if !ok || s.userID == 0 {
		return
	}
	investmentID, ok := data["id"]
	reference, rOk := data["reference"]
	year, yOk := data["year"]
	if !ok || !rOk || !yOk {
		return
	}
	investment, err := s.io.service.GetUserInvestment(int32(investmentID.(float64)))
	if err != nil || investment == nil {
		s.EmitServerError("UpdateUserInvestment: investment not found", err)
		return
	}
	if investment.UserID.Int != s.userID {
		return
	}
	investment, err = s.io.service.UpdateUserInvestment(int32(investmentID.(float64)), reference.(string), int32(year.(float64)))
	if err != nil {
		s.EmitServerError("UpdateUserInvestment: update user error", err)
		return
	}
	s.Emit(a.Type, investment)
}

// DeleteUserInvestment Delete user investment
func DeleteUserInvestment(s *Socket, a *Action) {
	data, ok := a.Data.(atreugo.JSON)
	if !ok {
		return
	}
	investmentID, ok := data["id"]
	if !ok {
		return
	}
	investment, err := s.io.service.GetUserInvestment(int32(investmentID.(float64)))
	if err != nil || investment == nil {
		s.EmitServerError("DeleteUserInvestment: investment not found", err)
		return
	}
	if investment.UserID.Int != s.userID {
		return
	}
	err = s.io.service.DeleteUserInvestment(int32(investmentID.(float64)))
	if err != nil {
		s.EmitServerError("DeleteUserInvestment: delete error", err)
		return
	}
	s.Emit(a.Type, investmentID)
}
