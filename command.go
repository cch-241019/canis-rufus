package canis_rufus

import (
	"bytes"
	"io"
)

/*
* @author: Chen Chiheng
* @date: 2025/1/3 22:15:20
* @description:
**/

type Group struct {
	ID    string
	Title string
}

type Command struct {
	Use                        string
	Aliases                    []string
	SuggestFor                 []string
	Short                      string
	GroupID                    string
	Long                       string
	Example                    string
	ValidArgs                  []string
	ValidArgsFunction          func(*Command, []string, string) ([]string, ShellComDirective)
	Args                       PositionalArgs
	ArgAliases                 []string
	BashCompletionFunc         string
	Deprecated                 string
	Annotations                map[string]string
	Version                    string
	PersistentPreRun           func(*Command, []string)
	PersistentPreRunE          func(*Command, []string) error
	PreRun                     func(*Command, []string)
	PreRunEr                   func(*Command, []string) error
	Run                        func(*Command, []string)
	RunE                       func(*Command, []string) error
	PostRun                    func(*Command, []string)
	PostRunE                   func(*Command, []string) error
	PersistentPostRun          func(*Command, []string)
	PersistentPostRunE         func(*Command, []string) error
	TraversChildren            bool
	Hidden                     bool
	SilenceErrors              bool
	SilenceUsage               bool
	DisableFlagParsing         bool
	DisableAutoGenTag          bool
	DisableFlagsInUseLine      bool
	DisableSuggestions         bool
	SuggestionsMinimumDistance int
	commandGroups              []*Group
	flagErrorBuf               *bytes.Buffer
	flags                      *FlagSet
	persistentFlags            *FlagSet
	localFlags                 *FlagSet
	parentFlags                *FlagSet
	normalizedNameFunc         func(*FlagSet, string) NormalizedName
	usageFunc                  func(*Command) error
	usageTemplate              string
	flagErrorFunc              func(*Command, error) error
	helpTemplate               string
	helpFunc                   func(*Command, []string)
	helpCommand                *Command
	helpCommandGroupID         string
	completionCommandGroupID   string
	versionTemplate            string
	errPrefix                  string
	inReader                   io.Reader
	outWriter                  io.Writer
	errWriter                  io.Writer
	commands                   []*Command
	parent                     *Command
	commandsMaxUseLen          int
	commandsMaxCommandPathLen  int
	commandsMaxNameLen         int
}

func (c *Command) AddFlag(flags ...*Flag) {

}
