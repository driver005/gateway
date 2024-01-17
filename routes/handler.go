package routes

type Routes struct {
	r Registry
}

func New(r Registry) *Routes {
	return &Routes{
		r: r,
	}
}

func (r Routes) SetAdminRoutes() {
	route := r.r.AdminRouter()
	r.r.AdminAuth().SetRoutes(route)

	r.r.AdminUpload().SetRoutes(route)
	r.r.AdminUser().SetRoutes(route)
	r.r.AdminVariant().SetRoutes(route)
}
