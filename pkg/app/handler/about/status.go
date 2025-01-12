package about

import (
	"html/template"
	"net/http"
	"soko/pkg/app/utils"
	"soko/pkg/database"
	"soko/pkg/models"
	"time"
)

// Index shows the landing page of the about pages
func Status(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(
		template.Must(
			template.New("status").
				Funcs(template.FuncMap{
					"timeSince": func(t time.Time) time.Duration {
						return time.Since(t).Round(time.Second)
					},
				}).
				ParseGlob("web/templates/layout/*.tmpl")).
			ParseGlob("web/templates/about/status.tmpl"))

	var applicationData []*models.Application
	database.DBCon.Model(&applicationData).Order("id").Select()

	templates.ExecuteTemplate(w, "status.tmpl", struct {
		Header       models.Header
		Application  models.Application
		Applications []*models.Application
	}{
		Header:       models.Header{Title: "About – ", Tab: "about"},
		Application:  utils.GetApplicationData(),
		Applications: applicationData,
	})
}
