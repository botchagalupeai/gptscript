package cli

import (
	"fmt"

	"github.com/gptscript-ai/gptscript/pkg/cache"
	"github.com/gptscript-ai/gptscript/pkg/config"
	"github.com/gptscript-ai/gptscript/pkg/credentials"
	"github.com/spf13/cobra"
)

type Delete struct {
	root *GPTScript
}

func (c *Delete) Customize(cmd *cobra.Command) {
	cmd.Use = "delete <tool name>"
	cmd.Aliases = []string{"rm", "del"}
	cmd.SilenceUsage = true
	cmd.Short = "Delete a stored credential"
	cmd.Args = cobra.ExactArgs(1)
}

func (c *Delete) Run(_ *cobra.Command, args []string) error {
	opts, err := c.root.NewGPTScriptOpts()
	if err != nil {
		return err
	}
	opts.Cache = cache.Complete(opts.Cache)

	cfg, err := config.ReadCLIConfig(c.root.ConfigFile)
	if err != nil {
		return fmt.Errorf("failed to read CLI config: %w", err)
	}

	store, err := credentials.NewStore(cfg, c.root.CredentialContext, opts.Cache.CacheDir)
	if err != nil {
		return fmt.Errorf("failed to get credentials store: %w", err)
	}

	if err = store.Remove(args[0]); err != nil {
		return fmt.Errorf("failed to remove credential: %w", err)
	}
	return nil
}
