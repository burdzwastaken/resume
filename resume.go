package main

import (
	"log"
	"os"
	"text/template"

	"github.com/BurntSushi/toml"
)

type Resume struct {
	Header     Header       `toml:"header"`
	Skills     []Skill      `toml:"skills"`
	OpenSource []OpenSource `toml:"opensource"`
	Experience []Experience `toml:"experience"`
	Education  []Education  `toml:"education"`
}

type Header struct {
	Name     string `toml:"name"`
	Phone    string `toml:"phone"`
	Email    string `toml:"email"`
	Handle   string `toml:"handle"`
	LinkedIn string `toml:"linkedin"`
	GitHub   string `toml:"github"`
	Website  string `toml:"website"`
	Summary  string `toml:"summary"`
}

type Skill struct {
	Title string   `toml:"title"`
	Items []string `toml:"items"`
}

type OpenSource struct {
	Project     string `toml:"project"`
	URL         string `toml:"url"`
	Description string `toml:"description"`
}

type Experience struct {
	Employer         string   `toml:"employer"`
	Role             string   `toml:"role"`
	Location         string   `toml:"location"`
	Timeframe        string   `toml:"timeframe"`
	Responsibilities []string `toml:"responsibilities"`
}

type Education struct {
	YearStart    int    `toml:"year_start"`
	YearComplete int    `toml:"year_complete"`
	Description  string `toml:"description"`
}

const resumeTemplate = `---
name: {{ .Header.Name }}
left-column:
  - '{{ .Header.Phone }}'
  - '{{ .Header.Email }}'
  - '{{ .Header.Handle }}'
right-column:
  - '{{ .Header.LinkedIn }}'
  - '{{ .Header.GitHub }}'
  - '{{ .Header.Website }}'
  - 'Last Updated: \today'
---

# Summary

{{ .Header.Summary }}

# Skills

{{ range .Skills }}**{{ .Title }}**: {{ range $i, $v := .Items }}{{ if $i }} · {{ end }}{{ $v }}{{ end }}

{{ end }}
# Open Source

{{ range .OpenSource }}**[{{ .Project }}]({{ .URL }})**: {{ .Description }}

{{ end }}
# Experience

{{ range .Experience }}## {{ .Employer }}
{{ .Role }}
{{ .Location }}
{{ .Timeframe }}

{{ range .Responsibilities }}* {{ . }}
{{ end }}
{{ end }}
# Education

{{ range .Education }}*{{ .YearStart }}*-*{{ .YearComplete }}*
: {{ .Description }}

{{ end }}
*References available upon request.*
`

func main() {
	var resume Resume
	if _, err := toml.DecodeFile("resume.toml", &resume); err != nil {
		log.Fatal("Error decoding resume.toml: ", err)
	}

	tmpl, err := template.New("resume").Parse(resumeTemplate)
	if err != nil {
		log.Fatal("Error parsing template: ", err)
	}

	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal("Error creating README.md: ", err)
	}

	if err := tmpl.Execute(f, resume); err != nil {
		_ = f.Close()
		log.Fatal("Error executing template: ", err)
	}

	if err := f.Close(); err != nil {
		log.Fatal("Error closing README.md: ", err)
	}
}
