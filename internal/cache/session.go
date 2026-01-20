package cache

type SessionManager struct {
	cache *Service
}

func NewSessionManager(cache *Service) *SessionManager {
	return &SessionManager{
		cache: cache,
	}
}
