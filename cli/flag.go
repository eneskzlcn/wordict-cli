package cli

import "strconv"

type FlagValue interface {
	String() string
	Set(string) error
	Type() string
}
type Flag struct {
	Value FlagValue
	Name string
	Shorthand string
	Usage string
	IsNoOption bool
	NoOptionDefaultValue string
}
func NewBoolFlag(value bool, name,shorthand,usage,noOptionDefaultValue string, isNoOption bool) *Flag {
	return &Flag{
		Value:     newBoolFlagValue(value),
		Name:      name,
		Shorthand: shorthand,
		Usage:     usage,
		IsNoOption: isNoOption,
		NoOptionDefaultValue: noOptionDefaultValue,
	}
}
func NewStringFlag(value string, name,shorthand,usage,noOptionDefaultValue string, isNoOption bool) *Flag {
	return &Flag{
		Value:     newStringFlagValue(value),
		Name:      name,
		Shorthand: shorthand,
		Usage:     usage,
		IsNoOption: isNoOption,
		NoOptionDefaultValue: noOptionDefaultValue,
	}
}
func NewIntFlag(value int, name,shorthand,usage,noOptionDefaultValue string, isNoOption bool) *Flag {
	return &Flag{
		Value:     newIntFlagValue(value),
		Name:      name,
		Shorthand: shorthand,
		Usage:     usage,
		IsNoOption: isNoOption,
		NoOptionDefaultValue: noOptionDefaultValue,
	}
}

type BoolFlag bool

func (b *BoolFlag) String() string {
	return strconv.FormatBool(bool(*b))
}
func newBoolFlagValue(value bool) FlagValue{
	a := BoolFlag(value)
	return &a
}
func (b *BoolFlag) Set(value string) error {
	val, err := strconv.ParseBool(value)
	*b = BoolFlag(val)
	if err != nil {
		return err
	}
	return nil
}

func (b *BoolFlag) Type() string {
	return "bool"
}

type StringFlag string

func newStringFlagValue(value string) FlagValue {
	a := StringFlag(value)
	return &a
}
func (s *StringFlag) String() string {
	return string(*s)
}

func (s *StringFlag) Set(value string) error {
	*s = StringFlag(value)
	return nil
}

func (s *StringFlag) Type() string {
	return "string"
}

type IntFlag int

func newIntFlagValue(value int) FlagValue {
	a := IntFlag(value)
	return &a
}
func (i *IntFlag) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

func (i *IntFlag) Set(value string) error {
	val, err := strconv.ParseInt(value,10, 32)
	if err != nil {
		return err
	}
	*i = IntFlag(val)
	return nil
}

func (i *IntFlag) Type() string {
	return "int"
}



