package converter

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/tkuchiki/tfmbm/resource"
)

func ConvertProjectIamMembers(bindings resource.BindingResources) error {
	tmpl := `{{$project  := .Project -}}{{$role  := .Role -}}
{{range $index, $member := .Members}}
resource "google_project_iam_member" "{{toID $member $role}}" {
  project = "{{$project}}"
  role    = "{{$role}}"
  member  = "{{$member}}"
}
{{end}}
`

	t, err := template.New("").Funcs(template.FuncMap{
		"toID": func(member, role string) string {
			account := strings.SplitN(member, ":", 2)
			m := strings.SplitN(account[1], "@", 2)
			localPart := strings.ReplaceAll(m[0], "-", "_")
			localPart = strcase.ToSnake(strings.ReplaceAll(localPart, ".", "_"))

			r := strings.SplitN(role, "/", 2)
			r2 := strings.ReplaceAll(r[1], ".", "_")
			r2 = strcase.ToSnake(r2)

			return fmt.Sprintf(`%s_is_%s`, localPart, r2)
		},
	}).Parse(tmpl)
	if err != nil {
		return err
	}

	for _, binding := range bindings {
		err := t.Execute(os.Stdout, *binding)
		if err != nil {
			return err
		}
	}

	return nil
}
