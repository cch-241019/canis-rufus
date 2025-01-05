package canis_rufus

import (
	"canis-rufus/value"
	stdflag "flag"
	"fmt"
	"io"
	"os"
	"sort"
)

/*
* @author: Chen Chiheng
* @date: 2025/1/3 20:33:50
* @description:
**/

type Flag struct {
	Name                string
	Shorthand           string
	Usage               string
	Default             string
	Deprecated          string
	NoOptDefVal         string
	ShorthandDeprecated string
	EnvVars             []string
	Value               value.Value
	Changed             bool
	Validate            func(flag *Flag) (Validator, error)
	Hidden              bool
	Annotations         map[string]string
}

func (f *Flag) defaultIsZeroValue() bool {
	switch f.Value.(type) {
	case *value.IntValue:
		return f.Default == "0"
	default:
		switch f.Value.String() {
		case "false":
			return true
		case "0":
			return true
		}
		return false
	}
}

type ErrorHanding int

const (
	ContinueOnError ErrorHanding = iota // 继续执行
	ExitOnError                         // 退出程序
	PanicOnError                        // 触发panic
)

type NormalizedName string

type ParseErrorsWhitelist struct {
	UnknownFlags bool
}

type FlagSet struct {
	Usage                func()
	SortFlags            bool
	ParseErrorsWhitelist ParseErrorsWhitelist
	name                 string // 命令的名称
	args                 []string
	actual               map[NormalizedName]*Flag
	argsLenAtDash        int
	addStdFlagSets       []stdflag.FlagSet
	errorHandling        ErrorHanding
	formal               map[NormalizedName]*Flag
	interspersed         bool
	normalizedNameFunc   func(f *FlagSet, name string) NormalizedName
	output               io.Writer
	orderedFormal        []*Flag // 按标志的添加顺序保存
	orderedActual        []*Flag
	parsed               bool
	sortedFormal         []*Flag
	sortedActual         []*Flag
	shorthands           map[byte]*Flag
}

func sortFlags(flags map[NormalizedName]*Flag) []*Flag {
	list := make(sort.StringSlice, len(flags))
	i := 0
	for name := range flags {
		list[i] = string(name)
		i++
	}
	list.Sort()
	result := make([]*Flag, len(list))
	for i, name := range list {
		result[i] = flags[NormalizedName(name)]
	}
	return result
}

func (f *FlagSet) SetNormalizeFunc(fn func(f *FlagSet, name string) NormalizedName) {
	f.normalizedNameFunc = fn
	f.sortedFormal = f.sortedFormal[:0]
	for current, flag := range f.formal {
		normalized := f.normalizeFlagName(flag.Name)
		if current == normalized {
			continue
		}
		flag.Name = string(normalized)
		delete(f.formal, current)
		f.formal[normalized] = flag
		if _, isSet := f.actual[current]; isSet {
			delete(f.actual, current)
			f.actual[normalized] = flag
		}
	}
}

func (f *FlagSet) normalizeFlagName(name string) NormalizedName {
	fn := f.GetNormalizeFunc()
	return fn(f, name)
}

func (f *FlagSet) GetNormalizeFunc() func(*FlagSet, string) NormalizedName {
	if f.normalizedNameFunc != nil {
		return f.normalizedNameFunc
	}
	return func(f *FlagSet, name string) NormalizedName { return NormalizedName(name) }
}

func (f *FlagSet) out() io.Writer {
	if f.output == nil {
		return os.Stderr
	}
	return f.output
}

func (f *FlagSet) SetOutput(output io.Writer) {
	f.output = output
}

func (f *FlagSet) VisitAll(fn func(*Flag)) {
	if len(f.formal) == 0 {
		return
	}

	var flags []*Flag
	if f.SortFlags {
		if len(f.formal) != len(f.sortedFormal) {
			f.sortedFormal = sortFlags(f.formal)
		}
		flags = f.sortedFormal
	} else {
		flags = f.orderedFormal
	}

	for _, flag := range flags {
		fn(flag)
	}
}

func (f *FlagSet) HasFlags() bool {
	return len(f.formal) > 0
}

func (f *FlagSet) HasAvailableFlags() bool {
	for _, flag := range f.formal {
		if !flag.Hidden {
			return true
		}
	}
	return false
}

func (f *FlagSet) Visit(fn func(*Flag)) {
	if len(f.actual) == 0 {
		return
	}

	var flags []*Flag
	if f.SortFlags {
		if len(f.actual) != len(f.sortedActual) {
			f.sortedActual = sortFlags(f.actual)
		}
		flags = f.sortedActual
	} else {
		flags = f.orderedActual
	}

	for _, flag := range flags {
		fn(flag)
	}
}

func (f *FlagSet) lookup(name NormalizedName) *Flag {
	return nil
}

func (f *FlagSet) Lookup(name string) *Flag {
	return f.lookup(f.normalizeFlagName(name))
}

// Set 为标志设定值
func (f *FlagSet) Set(name, value string) error {
	normalized := f.normalizeFlagName(name)
	flag, isExist := f.formal[normalized]
	if !isExist {
		return fmt.Errorf("no such flag -%v", name)
	}

	err := flag.Value.Set(value)
	if err != nil {
		var flagName string
		if flag.Shorthand != "" {
			flagName = fmt.Sprintf("-%s, --%s", flag.Shorthand, flag.Name)
		} else {
			flagName = fmt.Sprintf("--%s", flag.Name)
		}
		return fmt.Errorf("invalid argument %q for %q flag: %v", value, flagName, err)
	}

	if !flag.Changed {
		if f.actual == nil {
			f.actual = make(map[NormalizedName]*Flag)
		}
		f.actual[normalized] = flag
		f.orderedActual = append(f.orderedActual, flag)

		flag.Changed = true
	}

	return nil
}

// Changed 只要显示设置过这个标志就会返回true
func (f *FlagSet) Changed(name string) bool {
	flag := f.Lookup(name)
	if flag == nil {
		return false
	}
	return flag.Changed
}

type parseFunc func(flag *Flag, value string) error

func (f *FlagSet) parseLongArg(s string, args []string, fn parseFunc) ([]string, error) {
	return nil, nil
}

func (f *FlagSet) parseShortArg(s string, args []string, fn parseFunc) ([]string, error) {
	shorthands := s[1:]

	for len(shorthands) > 0 {
		// todo
	}
	return nil, nil
}

func (f *FlagSet) parseArgs(args []string, fn parseFunc) (err error) {
	for len(args) > 0 {
		s := args[0]
		args = args[1:]
		if len(s) == 0 || s[0] != '-' || len(s) == 1 {
			if !f.interspersed {
				f.args = append(f.args, s)
				f.args = append(f.args, args...)
				return
			}
			f.args = append(f.args, s)
			continue
		}

		if s[1] == '-' {
			if len(s) == 2 {
				f.argsLenAtDash = len(f.args)
				f.args = append(f.args, args...)
				break
			}
			args, err = f.parseLongArg(s, args, fn)
		} else {
			args, err = f.parseShortArg(s, args, fn)
		}
		if err != nil {
			return
		}
	}
	return
}

func (f *FlagSet) AddFlag(flag *Flag) {
	normalizedFlagName := f.normalizedNameFunc(f, flag.Name)
	_, isExist := f.formal[normalizedFlagName]
	if isExist {
		msg := fmt.Sprintf("%s flag redeinfed: %s", f.name, flag.Name)
		_, _ = fmt.Fprintln(f.out(), msg)
		panic(msg)
	}
}

func (f *FlagSet) Init(name string, errorHanding ErrorHanding) {
	f.name = name
	f.errorHandling = errorHanding
	f.argsLenAtDash = -1
}

func (f *FlagSet) Parsed() bool {
	return f.parsed
}

var CommandLine = NewFlagSet(os.Args[0], ExitOnError)

func Set(name, value string) error {
	return CommandLine.Set(name, value)
}

func NewFlagSet(name string, errorHanding ErrorHanding) *FlagSet {
	f := &FlagSet{
		name:          name,
		errorHandling: errorHanding,
		argsLenAtDash: -1,
		interspersed:  true,
		SortFlags:     true,
	}
	return f
}
