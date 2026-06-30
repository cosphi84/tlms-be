package auth

import "github.com/casbin/casbin/v3"

type Service struct {
	e *casbin.Enforcer
}

func NewService(e *casbin.Enforcer) *Service {
	if e == nil {
		panic("auth.NewService: nil enforcer")
	}
	return &Service{
		e: e,
	}
}

func (s *Service) Enforce(sub, obj, act string) (bool, error) {
	return s.e.Enforce(sub, obj, act)
}

func (s *Service) GrantRole(email, role string) (bool, error) {
	return s.e.AddGroupingPolicy(email, role)
}

func (s *Service) RevokeRole(email, role string) (bool, error) {
	return s.e.RemoveGroupingPolicy(email, role)
}

func (s *Service) GrantPermission(role, obj, act string) (bool, error) {
	return s.e.AddPolicy(role, obj, act)
}

func (s *Service) RevokePermission(role, obj, act string) (bool, error) {
	return s.e.RemovePolicy(role, obj, act)
}
