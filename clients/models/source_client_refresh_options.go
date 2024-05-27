package models

type SourceClientRefreshOptions struct {
	Force bool //force refresh even if token is not expired
}

func WithForce(force bool) func(*SourceClientRefreshOptions) {
	return func(s *SourceClientRefreshOptions) {
		s.Force = force
	}
}
