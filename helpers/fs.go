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

package helpers

import (
	"github.com/spf13/viper"
)

func GeneratorsPath() string {
	return BasePath() + viper.GetString("dir.generators") + "/"
}
func RootGeneratorPath(generator string) string {
	return GeneratorsPath() + generator + "/"
}
func RootGeneratorConfig(generator string) string {
	return GeneratorsPath() + generator + "/" + generator + ".json"
}
func RootGeneratorScript(generator string) string {
	return GeneratorsPath() + generator + "/" + generator + ".plis"
}
func ChildGeneratorConfigPath(generator string) string {
	return RootGeneratorPath(generator) + "config" + "/"
}
func GeneratorScriptsPath(generator string) string {
	return RootGeneratorPath(generator) + "scripts" + "/"
}
func GeneratorTemplatesPath(generator string) string {
	return RootGeneratorPath(generator) + "templates" + "/"
}
func ChildGeneratorConfig(generator string, child string) string {
	return RootGeneratorPath(generator) + "config" + "/" + child + ".json"
}
func ChildGeneratorScript(generator string, child string) string {
	return RootGeneratorPath(generator) + "scripts" + "/" + child + ".plis"
}
func GeneratorUserConfigPath(command string) string {
	return BasePath() + viper.GetString("dir.user") + "/" +
		"config" + "/" + command + ".json"
}
func GeneratorModulesPath(generator string) string {
	return RootGeneratorPath(generator) + "modules" + "/"
}
func GeneratorTemplatePath(generator string) string {
	return RootGeneratorPath(generator) + "templates" + "/"
}
func GeneratorTemplateFile(generator string, file string) string {
	return RootGeneratorPath(generator) + "templates" + "/" + file
}
func GeneratorModulesFile(generator string, module string) string {
	return GeneratorModulesPath(generator) + module + ".plis"
}
func BasePath() string {
	return viper.GetString("dir.base") + "/"
}
