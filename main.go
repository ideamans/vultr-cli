package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/ideamans/go-llm-cli-kit/llmcmd"
	"github.com/spf13/cobra"
	"github.com/vultr/govultr/v3"
	"golang.org/x/oauth2"

	"github.com/ideamans/vultr-cli/internal/llmdocs"
)

//go:generate go run . gen-llmdocs

// pluginVersion is the released version of this CLI. It is also the version
// recorded in plugins/vultr-cli/.claude-plugin/plugin.json — a test enforces
// that the two agree, and the release workflow enforces that both agree with
// the git tag. Bump it in the same commit as the tag.
const pluginVersion = "0.3.0"

// Overridden at release time by goreleaser via -ldflags.
var (
	version = pluginVersion
	commit  = "none"
	date    = "unknown"
)

var (
	apiKeyFlag    string
	forceIPv4Flag bool
)

// llmConfig describes the `vultr llm` subcommand.
func llmConfig() llmcmd.Config {
	return llmcmd.Config{Docs: llmdocs.Docs()}
}

// newRoot assembles the full command tree without executing it, so that both
// main and the catalog generator work from the same definition.
func newRoot() *cobra.Command {
	root := &cobra.Command{
		Use:           "vultr",
		Short:         "vultr is a CLI that wraps the entire Vultr API (via govultr)",
		Long:          "vultr is an aws-style CLI covering every Vultr API operation.\n\nUsage pattern: vultr <service> <operation> [args...]\nJSON body arguments accept inline JSON, @file.json, or - for stdin.\nSet VULTR_API_KEY in your environment (or pass --api-key).",
		Version:       fmt.Sprintf("%s (commit %s, built %s)", version, commit, date),
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.PersistentFlags().StringVar(&apiKeyFlag, "api-key", "", "Vultr API key (defaults to VULTR_API_KEY environment variable)")
	root.PersistentFlags().BoolVarP(&forceIPv4Flag, "ipv4", "4", false, "force IPv4 connections to the API (useful with IP-restricted API keys)")
	buildCommands(root)

	llmcmd.AddTo(root, llmConfig())
	root.AddCommand(newGenerateCommand())
	return root
}

func main() {
	// --llm anywhere on the command line prints the reference and exits,
	// bypassing cobra so it keeps working regardless of subcommand position.
	// Deprecated in favour of `vultr llm`, but removing it would break every
	// existing caller.
	if handled, err := llmcmd.HandleLegacy(os.Args[1:], llmConfig(), os.Stdout); handled {
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		return
	}

	if err := newRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func newClient() (*govultr.Client, error) {
	key := apiKeyFlag
	if key == "" {
		key = os.Getenv("VULTR_API_KEY")
	}
	if key == "" {
		return nil, errors.New("no API key: set VULTR_API_KEY or pass --api-key")
	}
	ctx := context.Background()
	if forceIPv4Flag {
		dialer := &net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}
		base := &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return dialer.DialContext(ctx, "tcp4", addr)
				},
			},
			Timeout: 60 * time.Second,
		}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, base)
	}
	config := &oauth2.Config{}
	ts := config.TokenSource(ctx, &oauth2.Token{AccessToken: key})
	return govultr.NewClient(oauth2.NewClient(ctx, ts)), nil
}

func printJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
