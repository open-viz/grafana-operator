/*
Copyright AppsCode Inc. and Contributors

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

package framework

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	api "go.openviz.dev/grafana-tools/apis/openviz/v1alpha1"
)

func (f *Framework) GetGrafanaDashboard() (*api.GrafanaDashboard, error) {
	db := &api.GrafanaDashboard{}
	if err := f.cc.Get(context.TODO(), client.ObjectKey{Namespace: f.namespace, Name: f.name}, db); err != nil {
		return nil, err
	}
	return db, nil
}

func (f *Framework) CreateGrafanaDashboard(db *api.GrafanaDashboard) error {
	return f.cc.Create(context.TODO(), db)
}

func (f *Framework) DeleteGrafanaDashboard(db *api.GrafanaDashboard) error {
	return f.cc.Delete(context.TODO(), db)
}
