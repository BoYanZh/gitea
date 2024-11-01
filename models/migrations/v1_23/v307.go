// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_23 //nolint

import (
	"code.gitea.io/gitea/modules/timeutil"

	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

type improveNotificationTableIndicesAction struct {
	ID     int64 `xorm:"pk autoincr"`
	UserID int64 `xorm:"INDEX NOT NULL"`
	RepoID int64 `xorm:"INDEX NOT NULL"`

	Status uint8 `xorm:"SMALLINT INDEX NOT NULL"`
	Source uint8 `xorm:"SMALLINT INDEX NOT NULL"`

	IssueID   int64  `xorm:"INDEX NOT NULL"`
	CommitID  string `xorm:"INDEX"`
	CommentID int64

	UpdatedBy int64 `xorm:"INDEX NOT NULL"`

	CreatedUnix timeutil.TimeStamp `xorm:"created INDEX NOT NULL"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated INDEX NOT NULL"`
}

// TableName sets the name of this table
func (*improveNotificationTableIndicesAction) TableName() string {
	return "notification"
}

// TableIndices implements xorm's TableIndices interface
func (*improveNotificationTableIndicesAction) TableIndices() []*schemas.Index {
	usuuIndex := schemas.NewIndex("u_s_uu", schemas.IndexType)
	usuuIndex.AddColumn("user_id", "status", "updated_unix")
	indices := []*schemas.Index{usuuIndex}

	return indices
}

func ImproveNotificationTableIndices(x *xorm.Engine) error {
	return x.Sync(&improveNotificationTableIndicesAction{})
}
