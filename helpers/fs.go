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
	"github.com/spf13/afero"
)

func GeneratorsPath() string {
	return BasePath() + viper.GetString("dir.generators") + afero.FilePathSeparator
}
func RootGeneratorPath(generator string) string {
	return GeneratorsPath() + generator + afero.FilePathSeparator
}
func RootGeneratorConfig(generator string) string {
	return GeneratorsPath() + generator + afero.FilePathSeparator + generator + ".json"
}
func RootGeneratorScript(generator string) string {
	return GeneratorsPath() + generator + afero.FilePathSeparator + generator + ".plis"
}
func ChildGeneratorConfigPath(generator string) string {
	return RootGeneratorPath(generator) + "config" + afero.FilePathSeparator
}
func GeneratorScriptsPath(generator string) string {
	return RootGeneratorPath(generator) + "scripts" + afero.FilePathSeparator
}
func GeneratorTemplatesPath(generator string) string {
	return RootGeneratorPath(generator) + "templates" + afero.FilePathSeparator
}
func ChildGeneratorConfig(generator string, child string) string {
	return RootGeneratorPath(generator) + "config" + afero.FilePathSeparator + child + ".json"
}
func ChildGeneratorScript(generator string, child string) string {
	return RootGeneratorPath(generator) + "scripts" + afero.FilePathSeparator + child + ".plis"
}
func GeneratorUserConfigPath(command string) string {
	return BasePath() + viper.GetString("dir.user") + afero.FilePathSeparator +
		"config" + afero.FilePathSeparator + command + ".json"
}
func GeneratorModulesPath(generator string) string {
	return RootGeneratorPath(generator) + "modules" + afero.FilePathSeparator
}
func GeneratorModulesFile(generator string, module string) string {
	return GeneratorModulesPath(generator)  + module + ".plis"
}
func BasePath() string {
	return viper.GetString("dir.base") + afero.FilePathSeparator
}
