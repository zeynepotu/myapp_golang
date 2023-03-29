package handlers

import (
	"github.com/zeynepotu/go-course/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.gohtml")

}
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.gohtml")
}
