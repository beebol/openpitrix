// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package repo

import (
	"context"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
)

func NewRepoManagerClient(ctx context.Context) (pb.RepoManagerClient, error) {
	conn, err := manager.NewClient(ctx, constants.RepoManagerHost, constants.RepoManagerPort)
	if err != nil {
		return nil, err
	}
	return pb.NewRepoManagerClient(conn), err
}
