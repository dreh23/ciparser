/*
 * Johannes Amorosa, (C) 2016
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"os/exec"
	"strings"
)

func plattform() (string, error) {
	// uname -sm
	var (
		plattform []byte
		e         error
	)
	cmdName := "uname"
	cmdArgs := []string{"-s"}
	if plattform, e = exec.Command(cmdName, cmdArgs...).Output(); e != nil {
		return "", e
	}
	return strings.TrimSpace(string(plattform)), nil
}

func bitOS() (string, error) {
	// uname -sm
	var (
		bit []byte
		e   error
	)
	cmdName := "uname"
	cmdArgs := []string{"-m"}
	if bit, e = exec.Command(cmdName, cmdArgs...).Output(); e != nil {
		return "", e
	}
	return strings.TrimSpace(string(bit)), nil
}