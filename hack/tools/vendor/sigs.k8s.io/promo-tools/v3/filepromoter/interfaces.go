/*
Copyright 2019 The Kubernetes Authors.

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

package filepromoter

import "context"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// SyncFileOp defines a synchronization operation.
//counterfeiter:generate . SyncFileOp
type SyncFileOp interface {
	Run(ctx context.Context) error
}
