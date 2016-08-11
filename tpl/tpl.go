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
	"gopkg.in/flosch/pongo2.v3"
	"strings"
)

func CopyTpl(pongoTpl *pongo2.TemplateSet, generator string, tpl string, dest string, context map[string]interface{}) error {
	tpl = strings.TrimPrefix(tpl,"/")
	parts := strings.Split(tpl,"/")
	if len(parts) <= 2{

	}
	return nil;
}//	t, err := pongoTpl.FromFile(helpers.GeneratorTemplateFile(generator, tpl))
//	if err != nil {
//		return err
//	}
//	s, err := t.Execute(pongo2.Context(context))
//	if err != nil {
//		return err
//	}
//	if (path.Ext(dest) != "") {
//		exists, err := afero.Exists(fs.WorkingDirFs(), dest)
//		if err != nil {
//			return err
//		}
//		if !exists {
//			ex, err := afero.Exists(fs.WorkingDirFs(), path.Dir(dest))
//			if err != nil {
//				return err
//			}
//			if !ex {
//				fs.WorkingDirFs().MkdirAll(path.Dir(dest), 755)
//			}
//		} else {
//			r, _ := afero.ReadFile(fs.WorkingDirFs(), dest)
//			if (reflect.DeepEqual(r, []byte(s))) {
//				fmt.Printf("The file `%s` is identical and is going to be ignored", dest)
//				fmt.Println()
//				return nil
//			}
//			if !prompter.YN(fmt.Sprintf("The file `%s` already exists do you want to replace it", dest), false) {
//				return nil
//			}
//		}
//
//
//		err = afero.WriteFile(afero.NewBasePathFs(fs.WorkingDirFs(), path.Dir(dest)), strings.Replace(dest, path.Dir(dest), "", 1), []byte(s), 755)
//		return err
//	}
//
//	if strings.HasPrefix(tpl, "/") {
//		tpl = strings.TrimPrefix(tpl, "/")
//	}
//	if tpl != "" || tpl != "." || tpl != "/" {
//		pp := strings.Split(tpl, "/")
//		tpl = ""
//		for _, v := range pp[1:] {
//			tpl += v + "/"
//		}
//		tpl = strings.TrimSuffix(tpl, "/")
//	}
//	destPath := path.Dir(dest + "/" + tpl);
//	exists, err := afero.Exists(fs.WorkingDirFs(), destPath)
//	if err != nil {
//		return err
//	}
//	if !exists {
//		fs.WorkingDirFs().MkdirAll(destPath, 755)
//	}
//
//	pth := destPath + "/" + strings.Replace(tpl, ".tpl", "", 1)
//	exists, err = afero.Exists(fs.WorkingDirFs(), pth)
//	if exists {
//		r, _ := afero.ReadFile(fs.WorkingDirFs(), pth)
//		if (reflect.DeepEqual(r, []byte(s))) {
//			fmt.Printf("The file `%s` is identical and is going to be ignored", pth)
//			fmt.Println()
//			return nil
//		}
//		if !prompter.YN(fmt.Sprintf("The file `%s` already exists do you want to replace it", pth), false) {
//			return nil
//		}
//	}
//	fmt.Println( destPath)
//	err = afero.WriteFile(afero.NewBasePathFs(fs.WorkingDirFs(), destPath), strings.Replace(pth, path.Dir(pth), "", 1), []byte(s), 755)
//	return err
//}
//func CopyAll(pongoTpl *pongo2.TemplateSet, generator string, tpl string, dest string, context map[string]interface{}) error {
//	res,err := getTemplates(generator, tpl)
//	if err != nil {
//		return err
//	}
//	for _, v := range res {
//		err := CopyTpl(pongoTpl, generator, v, dest, context)
//		if err != nil {
//			return err;
//		}
//	}
//	return nil
//}
//func getTemplates(generator  string, p string) ([]string, error) {
//	files := []string{}
//	err := afero.Walk(fs.WorkingDirFs(), helpers.GeneratorTemplateFile(generator, p), func(path string, info os.FileInfo, err error) error {
//		if !info.IsDir() {
//			path = strings.Replace(path, "\\", "/", -1)
//			files = append(files, strings.Replace(path, strings.Replace(helpers.GeneratorTemplatePath(generator), "\\", "/", -1), "", -1))
//		}
//		return nil
//	})
//	return files, err
//}