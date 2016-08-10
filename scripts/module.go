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

package scripts

import (
	anko_core "github.com/mattn/anko/builtins"
	anko_encoding_json "github.com/mattn/anko/builtins/encoding/json"
	anko_errors "github.com/mattn/anko/builtins/errors"
	anko_flag "github.com/mattn/anko/builtins/flag"
	anko_fmt "github.com/mattn/anko/builtins/fmt"
	anko_io "github.com/mattn/anko/builtins/io"
	anko_io_ioutil "github.com/mattn/anko/builtins/io/ioutil"
	anko_math "github.com/mattn/anko/builtins/math"
	anko_math_rand "github.com/mattn/anko/builtins/math/rand"
	anko_net "github.com/mattn/anko/builtins/net"
	anko_net_http "github.com/mattn/anko/builtins/net/http"
	anko_net_url "github.com/mattn/anko/builtins/net/url"
	anko_os "github.com/mattn/anko/builtins/os"
	anko_os_exec "github.com/mattn/anko/builtins/os/exec"
	anko_os_signal "github.com/mattn/anko/builtins/os/signal"
	anko_path "github.com/mattn/anko/builtins/path"
	anko_path_filepath "github.com/mattn/anko/builtins/path/filepath"
	anko_regexp "github.com/mattn/anko/builtins/regexp"
	anko_runtime "github.com/mattn/anko/builtins/runtime"
	anko_sort "github.com/mattn/anko/builtins/sort"
	anko_strings "github.com/mattn/anko/builtins/strings"
	anko_time "github.com/mattn/anko/builtins/time"
	"github.com/mattn/anko/vm"
	"github.com/spf13/cobra"
	"fmt"
	"github.com/kujtimiihoxha/plis-cli/generators"
	"strconv"
	"os"
	"github.com/spf13/afero"
	"github.com/kujtimiihoxha/plis-cli/fs"
	"github.com/kujtimiihoxha/plis-cli/helpers"
	"encoding/json"
)

var pkgs  map[string]func(env *vm.Env) *vm.Env;

var myEnv *vm.Env;

func Build(env *vm.Env, generator *generators.PlisGenerator, args []string) {
	anko_core.Import(env)
	pkgs = map[string]func(env *vm.Env) *vm.Env{
		"encoding/json": anko_encoding_json.Import,
		"errors":        anko_errors.Import,
		"flag":          anko_flag.Import,
		"fmt":           anko_fmt.Import,
		"io":            anko_io.Import,
		"io/ioutil":     anko_io_ioutil.Import,
		"math":          anko_math.Import,
		"math/rand":     anko_math_rand.Import,
		"net":           anko_net.Import,
		"net/http":      anko_net_http.Import,
		"net/url":       anko_net_url.Import,
		"os":            anko_os.Import,
		"os/exec":       anko_os_exec.Import,
		"os/signal":     anko_os_signal.Import,
		"path":          anko_path.Import,
		"path/filepath": anko_path_filepath.Import,
		"regexp":        anko_regexp.Import,
		"runtime":       anko_runtime.Import,
		"sort":          anko_sort.Import,
		"strings":       anko_strings.Import,
		"time":          anko_time.Import,
		"plis":          func(env *vm.Env) *vm.Env {
			return plisModule(env, generator, args)
		},
	}

	env.Define("import", func(s string) interface{} {
		if loader, ok := pkgs[s]; ok {
			m := loader(env)
			return m
		} else if loader, ok := pkgs[generator.GetRootParent().Config.Name + "/" + s]; ok {
			//Search for the root function ex. angular2 don't use the current all packages will be in the root func.
			m := loader(env)
			return m
		}
		panic(fmt.Sprintf("package '%s' not found", s))
	})
	env.Define("register", func(m string, obj map[string]interface{}) {
		if _, ok :=pkgs[m]; ok  {
			panic(fmt.Sprintf("package '%s' already exists", m))
		} else if _, ok := pkgs[generator.GetRootParent().Config.Name + "/" + m] ; ok {
			panic(fmt.Sprintf("package '%s' already exists", m))
		}
		//Search for the root function ex. angular2 don't use the current all packages will be in the root func.
		pkgs[generator.GetRootParent().Config.Name + "/" + m] = func(env *vm.Env) *vm.Env {
			md := env.NewModule(m)
			for k, v := range obj {
				md.Define(k, v)
			}
			return md
		}
	})
}
func addUserConfig(plis *vm.Env, command string) {
	data,err := afero.ReadFile(fs.WorkingDirFs(),helpers.GeneratorUserConfigPath(command))
	if err != nil {
		fmt.Println(fmt.Sprintf("Could not read config from `%s`", helpers.GeneratorUserConfigPath(command)))
		os.Exit(-1)
	}
	config := &map[string]interface{}{}
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Println(fmt.Sprintf("Could not read json from `%s`", helpers.GeneratorUserConfigPath(command)))
		os.Exit(-1)
	}
	plis.Define("UserConfig",*config)
}
func plisModule(env *vm.Env, gen *generators.PlisGenerator, args []string) *vm.Env {
	plis := env.NewPackage("plis")
	arguments := map[string]interface{}{};
	flags := map[string]interface{}{};
	if gen.Config.Arguments != nil && len(*gen.Config.Arguments) > 0 {
		for i, v := range args {
			arguments[(*gen.Config.Arguments)[i].Name] = argumentToType(v, (*gen.Config.Arguments)[i].Name, (*gen.Config.Arguments)[i].Type)
		}
	}
	if (gen.Config.Flags != nil) {
		for _, v := range *gen.Config.Flags {
			flags[v.Long] = flagToType(v.Long, gen.Cmd, v.Type)
		}
	}
	plis.Define("Args", arguments)
	plis.Define("Flags", flags)
	addUserConfig(plis,gen.GetRootParent().Config.Name)
	return plis
}
func flagToType(name string, cmd *cobra.Command, tp string) interface{} {
	switch tp {
	case "string":
		v, _ := cmd.Flags().GetString(name)
		return v
	case "int":
		v, _ := cmd.Flags().GetInt64(name)
		return v
	case "float":
		v, _ := cmd.Flags().GetFloat64(name)
		return v
	case "bool":
		v, _ := cmd.Flags().GetBool(name)
		return v
	default:
		return ""
	}
}
func argumentToType(arg string, name string, tp string) interface{} {
	switch tp {
	case "string":
		return arg
	case "int":
		number, err := strconv.ParseInt(arg, 10, 0)
		if err != nil {
			fmt.Println(fmt.Sprintf("The argument '%s' must be an INT type", name))
			os.Exit(-1)
		}
		return number
	case "float":
		floatNumber, err := strconv.ParseFloat(arg, 0)
		if err != nil {
			fmt.Println(fmt.Sprintf("The argument '%s' must be an Float type", name))
			os.Exit(-1)
		}
		return floatNumber
	case "bool":
		boolValue, err := strconv.ParseBool(arg)
		if err != nil {
			fmt.Println(fmt.Sprintf("The argument '%s' must be true/false", name))
			os.Exit(-1)
		}
		return boolValue
	default:
		return arg
	}
}