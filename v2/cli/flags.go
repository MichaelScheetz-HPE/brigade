package main

import "github.com/urfave/cli/v2"

const (
	flagAborted     = "aborted"
	flagAnyPhase    = "any-phase"
	flagBrowse      = "browse"
	flagCanceled    = "canceled"
	flagFailed      = "failed"
	flagFile        = "file"
	flagID          = "id"
	flagInsecure    = "insecure"
	flagNonTerminal = "non-terminal"
	flagOutput      = "output"
	flagPassword    = "password"
	flagPayload     = "payload"
	flagPayloadFile = "payload-file"
	flagPending     = "pending"
	flagProject     = "project"
	flagRoot        = "root"
	flagRunning     = "running"
	flagServer      = "server"
	flagSet         = "set"
	flagSource      = "source"
	flagSucceeded   = "succeeded"
	flagTerminal    = "terminal"
	flagTimedOut    = "timedout"
	flagType        = "type"
	flagUnknown     = "unknown"
	flagUnset       = "unset"
	flagYes         = "yes"
)

const (
	flagOutputJSON  = "json"
	flagOutputTable = "table"
	flagOutputYAML  = "yaml"
)

var (
	cliFlagOutput = &cli.StringFlag{
		Name:    flagOutput,
		Aliases: []string{"o"},
		Usage: "Return output in the specified format; supported formats: table, " +
			"yaml, json",
		Value: flagOutputTable,
	}
)
