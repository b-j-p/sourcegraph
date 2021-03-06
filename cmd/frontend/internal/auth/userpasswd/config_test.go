package userpasswd

import (
	"testing"

	"github.com/sourcegraph/sourcegraph/pkg/conf"
	"github.com/sourcegraph/sourcegraph/schema"
)

func TestValidateCustom(t *testing.T) {
	tests := map[string]struct {
		input        conf.Unified
		wantProblems []string
	}{
		"single": {
			input: conf.Unified{Critical: schema.CriticalConfiguration{
				AuthProviders: []schema.AuthProviders{
					{Builtin: &schema.BuiltinAuthProvider{Type: "builtin"}},
				},
			}},
			wantProblems: nil,
		},
		"multiple": {
			input: conf.Unified{Critical: schema.CriticalConfiguration{
				AuthProviders: []schema.AuthProviders{
					{Builtin: &schema.BuiltinAuthProvider{Type: "builtin"}},
					{Builtin: &schema.BuiltinAuthProvider{Type: "builtin"}},
				},
			}},
			wantProblems: []string{"at most 1"},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			conf.TestValidator(t, test.input, validateConfig, test.wantProblems)
		})
	}
}
