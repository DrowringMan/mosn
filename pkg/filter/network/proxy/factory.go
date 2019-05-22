/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proxy

import (
	"context"

	"sofastack.io/sofa-mosn/pkg/api/v2"
	"sofastack.io/sofa-mosn/pkg/config"
	"sofastack.io/sofa-mosn/pkg/filter"
	"sofastack.io/sofa-mosn/pkg/proxy"
	"sofastack.io/sofa-mosn/pkg/types"
)

func init() {
	filter.RegisterNetwork(v2.DEFAULT_NETWORK_FILTER, CreateProxyFactory)
}

type genericProxyFilterConfigFactory struct {
	Proxy *v2.Proxy
}

func (gfcf *genericProxyFilterConfigFactory) CreateFilterChain(context context.Context, clusterManager types.ClusterManager, callbacks types.NetWorkFilterChainFactoryCallbacks) {
	p := proxy.NewProxy(context, gfcf.Proxy, clusterManager)
	callbacks.AddReadFilter(p)
}

func CreateProxyFactory(conf map[string]interface{}) (types.NetworkFilterChainFactory, error) {
	return &genericProxyFilterConfigFactory{
		Proxy: config.ParseProxyFilter(conf),
	}, nil
}
