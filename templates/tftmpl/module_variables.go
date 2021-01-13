package tftmpl

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/hashicorp/consul-terraform-sync/templates/hcltmpl"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// LoadModuleVariables loads Terraform input variables from a file.
func LoadModuleVariables(filePath string) (hcltmpl.Variables, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return ParseModuleVariables(content, filePath)
}

// ParseModuleVariables parses bytes representing Terraform input variables
// for a module. It encodes the content into cty.Value types. Invalid HCL
// syntax and unsupported Terraform variable types result in an error.
func ParseModuleVariables(content []byte, filename string) (hcltmpl.Variables, error) {
	p := hclparse.NewParser()

	hclFile, diag := p.ParseHCL(content, filename)
	if diag.HasErrors() {
		return nil, diag
	}

	attrs, diag := hclFile.Body.JustAttributes()
	if diag.HasErrors() {
		return nil, diag
	}

	variables := make(hcltmpl.Variables)
	var diags hcl.Diagnostics
	for k, attr := range attrs {
		val, diag := attr.Expr.Value(&hcl.EvalContext{})
		if diag.HasErrors() {
			diags = diags.Extend(diag)
			continue
		}
		variables[k] = val
	}

	if diags.HasErrors() {
		return nil, diags
	}

	return variables, nil
}

// NewModuleVariablesTF writes content used for variables.module.tf of a
// Terraform root module. These variable defintions correspond to variables
// that are passed as arguments within the module block.
func newModuleVariablesTF(w io.Writer, input *RootModuleInputData) error {
	writePreamble(w, input.Task, ModuleVarsFilename)

	hclFile := hclwrite.NewEmptyFile()
	rootBody := hclFile.Body()
	rootBody.AppendNewline()

	lastIdx := len(input.Variables) - 1
	for i, name := range input.Variables.Keys() {
		v := input.Variables[name]
		vType := v.Type()

		vBody := rootBody.AppendNewBlock("variable", []string{name}).Body()
		vBody.SetAttributeValue("default", cty.NullVal(vType))

		rawTypeAttr := fmt.Sprintf("type = %s", variableTypeString(v, vType))
		vBody.AppendUnstructuredTokens(hclwrite.Tokens{{
			Type:  hclsyntax.TokenNil,
			Bytes: []byte(rawTypeAttr),
		}})
		vBody.AppendNewline()
		if i != lastIdx {
			rootBody.AppendNewline()
		}
	}

	// Format the file before writing
	content := hclFile.Bytes()
	content = hclwrite.Format(content)
	_, err := w.Write(content)
	return err
}
