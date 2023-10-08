package usecases

type Authorization interface {
}

type AuthService struct {
	repo Authorization
}

func NewAuthService(repo Authorization) *AuthService {
	return &AuthService{repo: repo}
}
