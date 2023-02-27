/*
 * Copyright (c) 2022-2023 Zander Schwid & Co. LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package properties

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CompanyData interface {

	MakeDir(appName string) (string, error)

	GetDir(appName string) string

	LoadJsonFile(appName, fileName string, v interface{}) (bool, error)

	SaveJsonFile(appName, fileName string, v interface{}) error

}

type implCompanyData struct {
	companyName string
}

func Locate(companyName string) CompanyData {
	return &implCompanyData{companyName: companyName}
}

func (t *implCompanyData) MakeDir(appName string) (string, error) {
	dir := AppDataDir(t.companyName, appName)
	err := os.MkdirAll(dir, 0700)
	return dir, err
}

func (t *implCompanyData) GetDir(appName string) string {
	return AppDataDir(t.companyName, appName)
}

func (t *implCompanyData) LoadJsonFile(appName, fileName string, v interface{}) (bool, error) {
	dir := AppDataDir(t.companyName, appName)
	fullPath := filepath.Join(dir, fileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false, nil
	}
	blob, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(blob, v)
	return err == nil, err
}

func (t *implCompanyData) SaveJsonFile(appName, fileName string, v interface{}) error {
	blob, err := json.Marshal(v)
	if err != nil {
		return err
	}
	dir, err := t.MakeDir(appName)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(dir, fileName), blob, 0660)
}
