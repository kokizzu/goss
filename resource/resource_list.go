// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/goss-org/goss/system"
	"github.com/goss-org/goss/util"
)

type AddrMap map[string]*Addr

func (r AddrMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Addr, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewAddr(ctx, sr, sys, config)
	res, err := NewAddr(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r AddrMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Addr, system.Addr, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewAddr(ctx, sr, sys, util.Config{})
	res, err := NewAddr(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *AddrMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Addr{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Addr
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *AddrMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Addr{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Addr
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type CommandMap map[string]*Command

func (r CommandMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Command, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewCommand(ctx, sr, sys, config)
	res, err := NewCommand(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r CommandMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Command, system.Command, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewCommand(ctx, sr, sys, util.Config{})
	res, err := NewCommand(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *CommandMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Command{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Command
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *CommandMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Command{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Command
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type DNSMap map[string]*DNS

func (r DNSMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*DNS, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewDNS(ctx, sr, sys, config)
	res, err := NewDNS(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r DNSMap) AppendSysResourceIfExists(sr string, sys *system.System) (*DNS, system.DNS, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewDNS(ctx, sr, sys, util.Config{})
	res, err := NewDNS(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *DNSMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := DNS{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*DNS
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *DNSMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := DNS{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*DNS
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type FileMap map[string]*File

func (r FileMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*File, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewFile(ctx, sr, sys, config)
	res, err := NewFile(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r FileMap) AppendSysResourceIfExists(sr string, sys *system.System) (*File, system.File, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewFile(ctx, sr, sys, util.Config{})
	res, err := NewFile(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *FileMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := File{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*File
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *FileMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := File{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*File
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type GossfileMap map[string]*Gossfile

func (r GossfileMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Gossfile, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewGossfile(ctx, sr, sys, config)
	res, err := NewGossfile(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r GossfileMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Gossfile, system.Gossfile, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewGossfile(ctx, sr, sys, util.Config{})
	res, err := NewGossfile(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *GossfileMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Gossfile{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Gossfile
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *GossfileMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Gossfile{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Gossfile
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type GroupMap map[string]*Group

func (r GroupMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Group, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewGroup(ctx, sr, sys, config)
	res, err := NewGroup(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r GroupMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Group, system.Group, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewGroup(ctx, sr, sys, util.Config{})
	res, err := NewGroup(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *GroupMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Group{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Group
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *GroupMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Group{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Group
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type PackageMap map[string]*Package

func (r PackageMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Package, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewPackage(ctx, sr, sys, config)
	res, err := NewPackage(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r PackageMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Package, system.Package, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewPackage(ctx, sr, sys, util.Config{})
	res, err := NewPackage(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *PackageMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Package{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Package
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *PackageMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Package{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Package
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type PortMap map[string]*Port

func (r PortMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Port, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewPort(ctx, sr, sys, config)
	res, err := NewPort(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r PortMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Port, system.Port, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewPort(ctx, sr, sys, util.Config{})
	res, err := NewPort(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *PortMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Port{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Port
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *PortMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Port{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Port
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type ProcessMap map[string]*Process

func (r ProcessMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Process, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewProcess(ctx, sr, sys, config)
	res, err := NewProcess(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r ProcessMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Process, system.Process, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewProcess(ctx, sr, sys, util.Config{})
	res, err := NewProcess(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *ProcessMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Process{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Process
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *ProcessMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Process{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Process
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type ServiceMap map[string]*Service

func (r ServiceMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Service, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewService(ctx, sr, sys, config)
	res, err := NewService(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r ServiceMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Service, system.Service, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewService(ctx, sr, sys, util.Config{})
	res, err := NewService(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *ServiceMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Service{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Service
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *ServiceMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Service{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Service
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type UserMap map[string]*User

func (r UserMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*User, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewUser(ctx, sr, sys, config)
	res, err := NewUser(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r UserMap) AppendSysResourceIfExists(sr string, sys *system.System) (*User, system.User, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewUser(ctx, sr, sys, util.Config{})
	res, err := NewUser(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *UserMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := User{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*User
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *UserMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := User{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*User
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type KernelParamMap map[string]*KernelParam

func (r KernelParamMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*KernelParam, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewKernelParam(ctx, sr, sys, config)
	res, err := NewKernelParam(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r KernelParamMap) AppendSysResourceIfExists(sr string, sys *system.System) (*KernelParam, system.KernelParam, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewKernelParam(ctx, sr, sys, util.Config{})
	res, err := NewKernelParam(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *KernelParamMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := KernelParam{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*KernelParam
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *KernelParamMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := KernelParam{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*KernelParam
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type MountMap map[string]*Mount

func (r MountMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Mount, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewMount(ctx, sr, sys, config)
	res, err := NewMount(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r MountMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Mount, system.Mount, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewMount(ctx, sr, sys, util.Config{})
	res, err := NewMount(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *MountMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Mount{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Mount
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *MountMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Mount{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Mount
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type InterfaceMap map[string]*Interface

func (r InterfaceMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Interface, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewInterface(ctx, sr, sys, config)
	res, err := NewInterface(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r InterfaceMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Interface, system.Interface, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewInterface(ctx, sr, sys, util.Config{})
	res, err := NewInterface(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *InterfaceMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := Interface{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Interface
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *InterfaceMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := Interface{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*Interface
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

type HTTPMap map[string]*HTTP

func (r HTTPMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*HTTP, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewHTTP(ctx, sr, sys, config)
	res, err := NewHTTP(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r HTTPMap) AppendSysResourceIfExists(sr string, sys *system.System) (*HTTP, system.HTTP, bool, error) {
	ctx := context.WithValue(context.Background(), idKey{}, sr)
	sysres := sys.NewHTTP(ctx, sr, sys, util.Config{})
	res, err := NewHTTP(sysres, util.Config{})
	if err != nil {
		return nil, nil, false, err
	}
	if e, _ := sysres.Exists(); !e {
		return res, sysres, false, nil
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true, nil
}

func (ret *HTTPMap) UnmarshalJSON(data []byte) error {
	// Curried json.Unmarshal
	unmarshal := func(i interface{}) error {
		if err := json.Unmarshal(data, i); err != nil {
			return err
		}
		return nil
	}

	// Validate configuration
	zero := HTTP{}
	whitelist, err := util.WhitelistAttrs(zero, util.JSON)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*HTTP
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}

func (ret *HTTPMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// Validate configuration
	zero := HTTP{}
	whitelist, err := util.WhitelistAttrs(zero, util.YAML)
	if err != nil {
		return err
	}
	if err := util.ValidateSections(unmarshal, zero, whitelist); err != nil {
		return err
	}

	var tmp map[string]*HTTP
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	typ := reflect.TypeOf(zero)
	typs := strings.Split(typ.String(), ".")[1]
	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*ret = tmp
	return nil
}
