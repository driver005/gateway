package registry

import "github.com/driver005/gateway/routes/admin"

func (m *Base) AdminAuth() *admin.Auth {
	if m.adminAuth == nil {
		m.adminAuth = admin.NewAuth(m)
	}
	return m.adminAuth
}
