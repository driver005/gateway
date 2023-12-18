package registry

import (
	"context"
	"time"

	"github.com/driver005/gateway/database"
	"github.com/driver005/gateway/logger"
	"github.com/driver005/gateway/sql"

	_ "github.com/driver005/gateway/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Base struct {
	l            *logger.Logger
	al           *logger.Logger
	buildVersion string
	buildHash    string
	buildDate    string
	r            Registry
	trc          trace.Tracer
	database     sql.Database
	db           *database.Handler
}

func (m *Base) with(r Registry) *Base {
	m.r = r
	return m
}

func (m *Base) WithBuildInfo(version, hash, date string) Registry {
	m.buildVersion = version
	m.buildHash = hash
	m.buildDate = date
	return m.r
}

func (m *Base) BuildVersion() string {
	return m.buildVersion
}

func (m *Base) BuildDate() string {
	return m.buildDate
}

func (m *Base) BuildHash() string {
	return m.buildHash
}

func (m *Base) WithLogger(l *logger.Logger) Registry {
	m.l = l
	return m.r
}

func (m *Base) Logger() *logger.Logger {
	if m.l == nil {
		m.l = logger.New("ORY Hydra", m.BuildVersion())
	}
	return m.l
}

func (m *Base) Tracer(ctx context.Context) trace.Tracer {
	if m.trc == nil {
		tp := otel.GetTracerProvider()
		m.trc = tp.Tracer("github.com/driver005/gateway", trace.WithInstrumentationVersion(m.BuildVersion()))
	}

	return m.trc
}

func (m *Base) AuditLogger() *logger.Logger {
	if m.al == nil {
		m.al = logger.NewAudit("ORY Hydra", m.BuildVersion())
	}
	return m.al
}

func (m *Base) RegisterRoutes(router *fiber.App) {
	_ = router.Group("/api/v1")
}

func (m *Base) Db() *database.Handler {
	if m.db == nil {
		m.db = database.NewHandler(m.r)
	}
	return m.db
}


func (m *Base) Database() sql.Database {
	return m.database
}

func (m *Base) Setup() {

	public := fiber.New(fiber.Config{
		ServerHeader:   "Fiber",
		AppName:        "Test App v1.0.1",
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		StrictRouting:  true,
		ReadBufferSize: 4096 * 10,
		Views:          html.New("./views", ".html"),
	})

	public.Use(favicon.New())
	public.Use(cors.New())
	// public.Use(csrf.New())
	public.Use(fiber_logger.New())
	// public.Use(limiter.New())
	public.Get("/swagger/*", swagger.HandlerDefault)
	// public.Get("/swagger/*", swagger.New(swagger.Config{
	// 	Layout: ,
	// }))

	m.RegisterRoutes(public)

	public.Listen("localhost:80")
}
