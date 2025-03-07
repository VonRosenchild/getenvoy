// Copyright 2019 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package binary

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// Run execs the file at the path with the args passed
func Run(path string, args []string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("unable to stat %q: %v", path, err)
	}
	_, filename := filepath.Split(path)
	// #nosec -> passthrough by design
	if err := syscall.Exec(path, append([]string{filename}, args...), os.Environ()); err != nil {
		return fmt.Errorf("unable to exec %q: %v", path, err)
	}
	return nil
}
