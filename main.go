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

	"github.com/spf13/cobra"
	"github.com/vultr/govultr/v3"
	"golang.org/x/oauth2"
)

// Overridden at release time by goreleaser via -ldflags.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	apiKeyFlag    string
	forceIPv4Flag bool
)

func main() {
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
	root.PersistentFlags().Bool("llm", false, "print the full machine-oriented reference for AI (LLMs) and exit")
	buildCommands(root)

	// --llm anywhere on the command line prints the LLM guide and exits,
	// bypassing cobra so it works regardless of subcommand position.
	for _, a := range os.Args[1:] {
		if a == "--" {
			break
		}
		if a == "--llm" {
			fmt.Print(llmHelp())
			return
		}
	}

	if err := root.Execute(); err != nil {
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
