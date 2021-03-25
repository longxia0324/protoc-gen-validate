package shared

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
)

func RegisterFunctions(tpl *template.Template, params pgs.Parameters) {
	tpl.Funcs(map[string]interface{}{
		"disabled":      Disabled,
		"ignored":       Ignored,
		"required":      RequiredOneOf,
		"context":       rulesContext,
		"render":        Render(tpl),
		"has":           Has,
		"needs":         Needs,
		"fileneeds":     FileNeeds,
		"isSp3Optional": IsSp3Optional,
	})
}

func IsSp3Optional(oneof pgs.OneOf) bool {
	fs := oneof.Fields()
	if len(fs) != 1 {
		return false
	}
	return fs[0].Descriptor().Proto3Optional != nil && *fs[0].Descriptor().Proto3Optional
}
