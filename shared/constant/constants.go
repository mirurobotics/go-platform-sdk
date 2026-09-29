// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/mirurobotics/go-platform-sdk/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type ConfigInstance string // Always "config_instance"
type ConfigSchema string   // Always "config_schema"
type ConfigType string     // Always "config_type"
type Deployment string     // Always "deployment"
type Device string         // Always "device"
type GitCommit string      // Always "git_commit"
type Group string          // Always "group"
type List string           // Always "list"
type Principal string      // Always "principal"
type Release string        // Always "release"

func (c ConfigInstance) Default() ConfigInstance { return "config_instance" }
func (c ConfigSchema) Default() ConfigSchema     { return "config_schema" }
func (c ConfigType) Default() ConfigType         { return "config_type" }
func (c Deployment) Default() Deployment         { return "deployment" }
func (c Device) Default() Device                 { return "device" }
func (c GitCommit) Default() GitCommit           { return "git_commit" }
func (c Group) Default() Group                   { return "group" }
func (c List) Default() List                     { return "list" }
func (c Principal) Default() Principal           { return "principal" }
func (c Release) Default() Release               { return "release" }

func (c ConfigInstance) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c ConfigSchema) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c ConfigType) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Deployment) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Device) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c GitCommit) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Group) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c List) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Principal) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Release) MarshalJSON() ([]byte, error)        { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
