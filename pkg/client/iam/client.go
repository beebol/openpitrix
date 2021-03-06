// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package iam

import (
	"context"
	"fmt"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
	"openpitrix.io/openpitrix/pkg/util/stringutil"
)

type Client struct {
	pb.AccountManagerClient
}

func NewClient() (*Client, error) {
	conn, err := manager.NewClient(constants.IAMServiceHost, constants.IAMServicePort)
	if err != nil {
		return nil, err
	}
	return &Client{
		AccountManagerClient: pb.NewAccountManagerClient(conn),
	}, nil
}

func (c *Client) GetUsers(ctx context.Context, userIds []string) ([]*pb.User, error) {
	var internalUsers []*pb.User
	var noInternalUserIds []string
	for _, userId := range userIds {
		if stringutil.StringIn(userId, constants.InternalUsers) {
			internalUsers = append(internalUsers, &pb.User{
				UserId: pbutil.ToProtoString(userId),
				Role:   pbutil.ToProtoString(constants.RoleGlobalAdmin),
			})
		} else {
			noInternalUserIds = append(noInternalUserIds, userId)
		}
	}

	if len(noInternalUserIds) == 0 {
		return internalUsers, nil
	}

	response, err := c.DescribeUsers(ctx, &pb.DescribeUsersRequest{
		UserId: noInternalUserIds,
	})
	if err != nil {
		logger.Error(ctx, "Describe users %s failed: %+v", noInternalUserIds, err)
		return nil, err
	}
	if len(response.UserSet) != len(noInternalUserIds) {
		logger.Error(ctx, "Describe users %s with return count [%d]", userIds, len(response.UserSet)+len(internalUsers))
		return nil, fmt.Errorf("describe users %s with return count [%d]", userIds, len(response.UserSet)+len(internalUsers))
	}
	response.UserSet = append(response.UserSet, internalUsers...)
	return response.UserSet, nil
}
