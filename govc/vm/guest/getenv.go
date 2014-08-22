/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package guest

import (
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
)

type getenv struct {
	*GuestFlag
}

func init() {
	cli.Register(&getenv{GuestFlag: NewGuestFlag()})
}

func (cmd *getenv) Register(f *flag.FlagSet) {
}

func (cmd *getenv) Process() error {
	return nil
}

func (cmd *getenv) Run(f *flag.FlagSet) error {
	m, err := cmd.ProcessManager()
	if err != nil {
		return err
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	vars, err := m.ReadEnvironmentVariableInGuest(vm, cmd.Auth(), f.Args())
	if err != nil {
		return err
	}

	for _, v := range vars {
		fmt.Printf("%s\n", v)
	}

	return nil
}