// Copyright © 2016 Kujtim Hoxha <kujtimii.h@gmail.com>
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

package tpl

import (
	"fmt"
	"github.com/Songmu/prompter"
	"github.com/kujtimiihoxha/plis-cli/fs"
	"github.com/spf13/afero"
	"os"
	"reflect"
)

func CopyTpl(data string, file string, dest string) error {
	exists, err := afero.Exists(fs.WorkingDirFs(), dest)
	if err != nil {
		return err
	}
	if !exists {
		fs.WorkingDirFs().MkdirAll(dest, os.ModePerm)
	}
	exists, err = afero.Exists(fs.WorkingDirFs(), dest+file)
	if err != nil {
		return err
	}
	if exists {
		r, _ := afero.ReadFile(fs.WorkingDirFs(), dest+file)
		if reflect.DeepEqual(r, []byte(data)) {
			fmt.Printf("The file `%s` is identical and is going to be ignored", dest+file)
			fmt.Println()
			return nil
		}
		if !prompter.YN(fmt.Sprintf("The file `%s` already exists do you want to replace it", dest+file), false) {
			return nil
		}
	}
	return afero.WriteFile(afero.NewBasePathFs(fs.WorkingDirFs(), dest), file, []byte(data), os.ModePerm)
}
