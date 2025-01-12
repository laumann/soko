package utils

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"soko/pkg/models"
	"strings"
)

type stabilization struct {
	XMLName  xml.Name `xml:"stabilization" json:"-"`
	Category string   `xml:"category" json:"category"`
	Package  string   `xml:"package" json:"package"`
	Version  string   `xml:"version" json:"version"`
	Message  string   `xml:"message" json:"message"`
}

func (s stabilization) String() string {
	return s.Category + "/" + s.Package + "-" + s.Version + " # " + s.Message
}

func StabilizationExport(w http.ResponseWriter, pageUrl string, gpackages []*models.Package) {
	result := make([]stabilization, 0)
	for _, gpackage := range gpackages {
		for _, version := range gpackage.Versions {
			for _, pkgcheck := range version.PkgCheckResults {
				result = append(result, stabilization{
					Category: pkgcheck.Category,
					Package:  pkgcheck.Package,
					Version:  pkgcheck.Version,
					Message:  pkgcheck.Message,
				})
			}
		}
	}

	_, extension, _ := strings.Cut(pageUrl, ".")
	switch extension {
	case "json":
		b, err := json.Marshal(result)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	case "xml":
		b, err := xml.Marshal(struct {
			XMLName  xml.Name `xml:"xml"`
			Packages []stabilization
		}{Packages: result})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/xml")
		w.Write(b)
	case "list":
		var lines string
		for _, pkg := range result {
			lines += pkg.String() + "\n"
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(lines))
	}
}
