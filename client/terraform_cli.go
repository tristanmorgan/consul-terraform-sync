package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hashicorp/consul-nia/templates/tftmpl"
	"github.com/hashicorp/terraform-exec/tfexec"
)

// workspaceEnv is the environment variable to set a Terraform workspace
const workspaceEnv = "TF_WORKSPACE"

var _ Client = (*TerraformCLI)(nil)

// TerraformCLI is the client that wraps around terraform-exec
// to execute Terraform cli commands
type TerraformCLI struct {
	tf         terraformExec
	logLevel   string
	workingDir string
	workspace  string
	execPath   string
}

// TerraformCLIConfig configures the Terraform client
type TerraformCLIConfig struct {
	LogLevel   string
	ExecPath   string
	WorkingDir string
	Workspace  string
}

// NewTerraformCLI creates a terraform-exec client and configures and
// initializes a new Terraform client
func NewTerraformCLI(config *TerraformCLIConfig) (*TerraformCLI, error) {
	if config == nil {
		return nil, errors.New("TerraformCLIConfig cannot be nil - no meaningful default values")
	}

	tfPath := filepath.Join(config.ExecPath, "terraform")
	tf, err := tfexec.NewTerraform(config.WorkingDir, tfPath)
	if err != nil {
		return nil, err
	}

	client := &TerraformCLI{
		tf:         tf,
		logLevel:   config.LogLevel,
		execPath:   tfPath,
		workingDir: config.WorkingDir,
		workspace:  config.Workspace,
	}
	log.Printf("[TRACE] (client.terraformcli) created Terraform CLI client %s", client.GoString())

	return client, nil
}

// Init executes the cli command a `terraform init`
func (t *TerraformCLI) Init(ctx context.Context) error {
	if err := t.tf.Init(ctx); err != nil {
		return err
	}
	return t.workspaceNew(ctx)
}

// Apply executes the cli command `terraform apply` for a given workspace
func (t *TerraformCLI) Apply(ctx context.Context) error {

	// Set the workspace in the environment
	env := make(map[string]string)
	env[workspaceEnv] = t.workspace
	t.tf.SetEnv(env)

	tfvarFile := strings.TrimRight(tftmpl.TFVarsTmplFilename, ".tmpl")
	return t.tf.Apply(ctx, tfexec.VarFile(tfvarFile))
}

// Plan executes the cli command `terraform plan` for a given workspace
func (t *TerraformCLI) Plan(ctx context.Context) error {
	return t.tf.Plan(ctx)
}

// workspaceNew make the `terraform workspace new` cli call. At the time of writing
// terraform-exec package does not support this cli command yet:
// https://github.com/hashicorp/terraform-exec/issues/4
// TODO: revisit this and replace with terraform-exec when ready
// and update TestTerraformCLIInit which is currently being skipped
func (t *TerraformCLI) workspaceNew(ctx context.Context) error {
	// the args/env vars mimic the default ones set by terraform-exec
	args := []string{"workspace", "new", t.workspace, "-no-color"}
	workspaceEnv := fmt.Sprintf("%s=%s", workspaceEnv, t.workspace)
	env := []string{workspaceEnv, "TF_IN_AUTOMATION=1", "TF_LOG=", "CHECKPOINT_DISABLE="}

	cmd := exec.CommandContext(ctx, t.execPath, args...)
	cmd.Env = env
	cmd.Dir = t.workingDir

	// log out stdout, capture stderr in buffer
	cmd.Stdout = os.Stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		errStr := string(stderr.Bytes())
		if strings.Contains(errStr, "already exists") {
			log.Printf("[DEBUG] (client.terraformcli) workspace already exists: '%s'", t.workspace)
			return nil
		}
		return fmt.Errorf("Error creating workspace %s: %s: %s", t.workspace, err.Error(), errStr)
	}
	return nil
}

// GoString defines the printable version of this struct.
func (t *TerraformCLI) GoString() string {
	if t == nil {
		return "(*TerraformCLI)(nil)"
	}

	return fmt.Sprintf("&TerraformCLI{"+
		"LogLevel:%s, "+
		"WorkingDir:%s, "+
		"WorkSpace:%s, "+
		"}",
		t.logLevel,
		t.workingDir,
		t.workspace,
	)
}