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

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestNacosClient_GetDomain(t *testing.T) {
	//s := `{"dom":"hello123","cacheMillis":10000,"useSpecifiedURL":false,"hosts":[{"valid":true,"marked":false,"metadata":{},"instanceId":"","port":81,"ip":"2.2.2.2","weight":1.0,"enabled":true}],"checksum":"c7befb32f3bb5b169f76efbb0e1f79eb1542236821437","lastRefTime":1542236821437,"env":"","clusters":""}`
	s := `{
"metadata": {},
"dom": "nacos.test.3",
"cacheMillis": 3000,
"useSpecifiedURL": false,
"hosts": [
{
"valid": true,
"marked": false,
"metadata": {},
"instanceId": "192.168.25.129#8848#KanBan#DEFAULT_GROUP@@nacos.test.3",
"port": 8848,
"healthy": true,
"ip": "2.2.2.2",
"clusterName": "KanBan",
"weight": 1,
"ephemeral": true,
"serviceName": "nacos.test.3",
"enabled": true
}
],
"name": "DEFAULT_GROUP@@nacos.test.3",
"checksum": "b9ca0722f872f1d3c2d7ee0213c9ba9c",
"lastRefTime": 1583376088992,
"env": "",
"clusters": ""
}
`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.EscapedPath() == "/nacos/v1/ns/api/srvIPXT" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(s))
		} else if req.URL.EscapedPath() == "/nacos/v1/ns/api/allDomNames" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
"count": 1,
"doms": {
"public": [
"nacos.test.3"
]
}
}`))
		}

	}))

	port, _ := strconv.Atoi(strings.Split(server.URL, ":")[2])

	defer server.Close()

	vc := NacosClient{NewConcurrentMap(), UDPServer{}, ServerManager{}, port}
	vc.udpServer.vipClient = &vc
	vc.SetServers([]string{strings.Split(strings.Split(server.URL, "http://")[1], ":")[0]})
	vc.getAllDomNames()
	instance := vc.SrvInstance("nacos.test.3", "127.0.0.1")
	if strings.Compare(instance.IP, "2.2.2.2") == 0 {
		t.Log("Passed")
	}
}

func Test_getInstens(t *testing.T) {

	s := `{
"metadata": {},
"dom": "nacos.test.3",
"cacheMillis": 3000,
"useSpecifiedURL": false,
"hosts": [
{
"valid": true,
"marked": false,
"metadata": {
"protocol": "http"
},
"instanceId": "192.168.25.127#8843#KanBan#DEFAULT_GROUP@@nacos.test.3",
"port": 8843,
"healthy": true,
"ip": "192.168.25.127",
"clusterName": "KanBan",
"weight": 2,
"ephemeral": true,
"serviceName": "nacos.test.3",
"enabled": true
},
{
"valid": true,
"marked": false,
"metadata": {
"protocol": "http"
},
"instanceId": "192.168.25.129#8841#KanBan#DEFAULT_GROUP@@nacos.test.3",
"port": 8841,
"healthy": true,
"ip": "192.168.25.129",
"clusterName": "KanBan",
"weight": 1,
"ephemeral": true,
"serviceName": "nacos.test.3",
"enabled": true
},
{
"valid": true,
"marked": false,
"metadata": {
"protocol": "http"
},
"instanceId": "192.168.25.245#8842#KanBan#DEFAULT_GROUP@@nacos.test.3",
"port": 8842,
"healthy": true,
"ip": "192.168.25.245",
"clusterName": "KanBan",
"weight": 2,
"ephemeral": true,
"serviceName": "nacos.test.3",
"enabled": true
}
],
"name": "DEFAULT_GROUP@@nacos.test.3",
"checksum": "1cbfc5f52b1ea28f60288dfa46344f17",
"lastRefTime": 1583487192431,
"env": "",
"clusters": ""
}

`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.EscapedPath() == "/nacos/v1/ns/api/srvIPXT" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(s))
		} else if req.URL.EscapedPath() == "/nacos/v1/ns/api/allDomNames" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
"count": 1,
"doms": {
"public": [
"nacos.test.3"
]
}
}`))
		}

	}))

	port, _ := strconv.Atoi(strings.Split(server.URL, ":")[2])

	defer server.Close()

	vc := NacosClient{NewConcurrentMap(), UDPServer{}, ServerManager{}, port}
	vc.udpServer.vipClient = &vc
	vc.SetServers([]string{strings.Split(strings.Split(server.URL, "http://")[1], ":")[0]})
	vc.getAllDomNames()
	for i := 1; i < 10; i++ {
		t.Log("SrvInstances:" + string(i))
		instances := vc.SrvInstances("nacos.test.3", "127.0.0.1")
		for _, v := range instances {
			t.Log(v.String())
		}
	}

}
