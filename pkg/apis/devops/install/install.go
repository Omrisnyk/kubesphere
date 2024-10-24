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
package install

import (
	"github.com/emicklei/go-restful"
	urlruntime "k8s.io/apimachinery/pkg/util/runtime"
	devopsv1alpha2 "kubesphere.io/kubesphere/pkg/apis/devops/v1alpha2"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
)

func init() {
	Install(runtime.Container)
}

func Install(c *restful.Container) {
	urlruntime.Must(devopsv1alpha2.AddToContainer(c))
}
