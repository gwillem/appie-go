package main

import (
	"testing"

	"github.com/jessevdk/go-flags"
)

func TestCountryFlagEmptyWhenOmitted(t *testing.T) {
	// When omitted the flag is empty; country resolution (stored config, else
	// nl) happens in the library, not via a CLI default.
	globalOpts.Country = ""
	p := flags.NewParser(&globalOpts, flags.HelpFlag)
	p.SubcommandsOptional = true
	if _, err := p.ParseArgs([]string{}); err != nil {
		t.Fatalf("parse: %v", err)
	}
	if globalOpts.Country != "" {
		t.Errorf("Country = %q, want empty when omitted", globalOpts.Country)
	}
}

func TestCountryFlagParsesBE(t *testing.T) {
	globalOpts.Country = ""
	p := flags.NewParser(&globalOpts, flags.HelpFlag)
	p.SubcommandsOptional = true
	if _, err := p.ParseArgs([]string{"--country", "be"}); err != nil {
		t.Fatalf("parse: %v", err)
	}
	if globalOpts.Country != "be" {
		t.Errorf("Country = %q, want be", globalOpts.Country)
	}
}
