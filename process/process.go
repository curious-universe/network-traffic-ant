/*
Copyright © 2021 curious-universe

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

package process

import (
	"github.com/curious-universe/go-ps"
	"github.com/curious-universe/network-traffic-ant/nerror"
)

// FindProcess find the process by binary program
func FindProcessByName(key string) (p *ps.Process, err error) {
	keyAllProcess := FindKeyAllProcess(key)

	if len(keyAllProcess) == 1 {
		return &keyAllProcess[0], nil
	} else if len(keyAllProcess) > 1 {
		return nil, nerror.ErrTooManySameNameProcess
	} else {
		return nil, nerror.ErrNotFoundProcess
	}
}

func FindProcessByNameAndPid(key string, pid int) (p *ps.Process, err error) {
	keyAllProcess := FindKeyAllProcess(key)

	if len(keyAllProcess) > 0 {
		for _, pp := range keyAllProcess {
			if pp.Pid() == pid {
				return &pp, nil
			}
		}
	}

	return nil, nerror.ErrNotFoundProcess
}

func FindKeyAllProcess(key string) []ps.Process {
	allProcess := make([]ps.Process, 0)
	procs, _ := ps.Processes()
	for _, proc := range procs {
		if proc.Executable() == key {
			allProcess = append(allProcess, proc)
		}
	}
	return allProcess
}
