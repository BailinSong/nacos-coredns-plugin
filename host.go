/*
 * Copyright 1999-2018 Alibaba Group Holding Ltd.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nacos_coredns_plugin

import "encoding/json"

type Instance struct {
	InstanceId  string
	ServiceName string
	Enabled     bool
	IP          string
	Port        int
	Weight      float64
	Valid       bool
	Unit        string
	AppUseType  string
	Site        string
	Metadata    map[string]string
	clusterName string
}

func (v *Instance) equals(v2 Instance) bool {
	return v.InstanceId == v2.InstanceId && v.Enabled == v2.Enabled && v.Weight == v2.Weight && v.Valid == v2.Valid
}

func (h Instance) String() string {
	bs, err := json.Marshal(&h)

	if err != nil {
		return ""
	}

	return string(bs)
}
