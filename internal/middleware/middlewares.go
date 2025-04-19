package middleware

import (
	"github.com/ayoadeoye1/go-clean-archi/config"
	"github.com/ayoadeoye1/go-clean-archi/internal/auth"
	"github.com/ayoadeoye1/go-clean-archi/internal/session"
	"github.com/ayoadeoye1/go-clean-archi/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	sessUC  session.UCSession
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{sessUC: sessUC, authUC: authUC, cfg: cfg, origins: origins, logger: logger}
}
