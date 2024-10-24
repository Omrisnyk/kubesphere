/*

 Copyright 2019 The KubeSphere Authors.

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
package main

import (
	"kubesphere.io/kubesphere/cmd/ks-apiserver/app"
	"log"
	// Install apis
	_ "kubesphere.io/kubesphere/pkg/apis/devops/install"
	_ "kubesphere.io/kubesphere/pkg/apis/logging/install"
	_ "kubesphere.io/kubesphere/pkg/apis/metrics/install"
	_ "kubesphere.io/kubesphere/pkg/apis/monitoring/install"
	_ "kubesphere.io/kubesphere/pkg/apis/operations/install"
	_ "kubesphere.io/kubesphere/pkg/apis/resources/install"
	_ "kubesphere.io/kubesphere/pkg/apis/servicemesh/metrics/install"
	_ "kubesphere.io/kubesphere/pkg/apis/tenant/install"
	_ "kubesphere.io/kubesphere/pkg/apis/terminal/install"
)

func main() {

	cmd := app.NewAPIServerCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
