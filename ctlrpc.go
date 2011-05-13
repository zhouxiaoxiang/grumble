// Copyright (c) 2011 The Grumble Authors
// The use of this source code is goverened by a BSD-style
// license that can be found in the LICENSE-file.

package main

import (
	"os"
)

type ControlRPC struct {

}

// Start a server
func (c *ControlRPC) Start(Id int64, out *int) os.Error {
	server, exists := servers[Id]
	if !exists {
		return os.NewError("no such server")
	}
	_ = server
	return nil
}

// Stop a server
func (c *ControlRPC) Stop(Id int64, out *int) os.Error {
	server, exists := servers[Id]
	if !exists {
		return os.NewError("no such server")
	}
	_ = server
	return nil
}

type ConfigValue struct {
	Id    int64
	Key   string
	Value string
}

// Set a config value
func (c *ControlRPC) SetConfig(in *ConfigValue, out *ConfigValue) os.Error {
	server, exists := servers[in.Id]
	if !exists {
		return os.NewError("no such server")
	}
	server.cfg.Set(in.Key, in.Value)
	out.Id = in.Id
	out.Key = in.Key
	out.Value = in.Value
	return nil
}

// Get a config value
func (c *ControlRPC) GetConfig(in *ConfigValue, out *ConfigValue) os.Error {
	server, exists := servers[in.Id]
	if !exists {
		return os.NewError("no such server")
	}
	out.Id = in.Id
	out.Key = in.Key
	out.Value = server.cfg.StringValue(in.Key)
	return nil
}
