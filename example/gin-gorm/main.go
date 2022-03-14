package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

const (
	indexTmpl   = "index"
	profileTmpl = "profile"
)

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	db.Create(&User{Username: "world"})
	db.Create(&User{Username: "foo-bar"})

	handler := &Handler{
		db: db,
	}

	router := gin.Default()
	router.SetHTMLTemplate(parseTemplates())
	router.Use(otelgin.Middleware("service-name"))
	router.GET("/", handler.Index)
	router.GET("/hello/:username", handler.Hello)

	if err := router.Run("localhost:9999"); err != nil {
		log.Print(err)
	}
}

type Handler struct {
	db *gorm.DB
}

func (h *Handler) Index(c *gin.Context) {
	ctx := c.Request.Context()
	otelgin.HTML(c, http.StatusOK, indexTmpl, gin.H{
		"traceURL": otelplay.TraceURL(trace.SpanFromContext(ctx)),
	})
}

func (h *Handler) Hello(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.Param("username")
	user := new(User)
	if err := h.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		c.Error(err)
		return
	}

	otelgin.HTML(c, http.StatusOK, profileTmpl, gin.H{
		"username": user.Username,
		"traceURL": otelplay.TraceURL(trace.SpanFromContext(ctx)),
	})
}

type User struct {
	ID       int64 `gorm:"primaryKey"`
	Username string
}

func parseTemplates() *template.Template {
	indexTemplate := `
		<html>
		<p>Here are some routes for you:</p>
		<ul>
			<li><a href="/hello/world">Hello world</a></li>
			<li><a href="/hello/foo-bar">Hello foo-bar</a></li>
		</ul>
		<p><a href="{{ .traceURL }}" target="_blank">{{ .traceURL }}</a></p>
		</html>
	`
	t := template.Must(template.New(indexTmpl).Parse(indexTemplate))

	profileTemplate := `
		<html>
		<h3>Hello {{ .username }}</h3>
		<p><a href="{{ .traceURL }}" target="_blank">{{ .traceURL }}</a></p>
		</html>
	`
	return template.Must(t.New(profileTmpl).Parse(profileTemplate))
}
