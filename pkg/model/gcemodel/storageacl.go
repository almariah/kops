/*
Copyright 2016 The Kubernetes Authors.

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

package gcemodel

import (
	"fmt"

	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/gce"
	"k8s.io/kops/upup/pkg/fi/cloudup/gcetasks"
	"k8s.io/kops/util/pkg/vfs"
)

// StorageAclBuilder configures storage acls
type StorageAclBuilder struct {
	*GCEModelContext
	Cloud     gce.GCECloud
	Lifecycle *fi.Lifecycle
}

var _ fi.ModelBuilder = &NetworkModelBuilder{}

// Build creates the tasks that set up storage acls
func (b *StorageAclBuilder) Build(c *fi.ModelBuilderContext) error {
	clusterPath := b.Cluster.Spec.ConfigBase
	p, err := vfs.Context.BuildVfsPath(clusterPath)
	if err != nil {
		return fmt.Errorf("cannot parse cluster path %q: %v", clusterPath, err)
	}

	serviceAccount, err := b.Cloud.ServiceAccount()
	if err != nil {
		return fmt.Errorf("error fetching ServiceAccount: %v", err)
	}

	switch p := p.(type) {
	case *vfs.GSPath:
		// It's not ideal that we have to do this at the bucket level,
		// but GCS doesn't seem to have a way to do subtrees (like AWS IAM does)
		// Note this permission only lets us list objects, not read them
		c.AddTask(&gcetasks.StorageBucketAcl{
			Name:      s("serviceaccount-statestore-list"),
			Lifecycle: b.Lifecycle,
			Bucket:    s(p.Bucket()),
			Entity:    s("user-" + serviceAccount),
			Role:      s("READER"),
		})
	}

	return nil
}
