// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalPandaUserRoleDao is an internal type for wrapping the internal DAO implementation.
type internalPandaUserRoleDao = *internal.PandaUserRoleDao

// pandaUserRoleDao is the data access object for the table panda_user_role.
// You can define custom methods on it to extend its functionality as needed.
type pandaUserRoleDao struct {
	internalPandaUserRoleDao
}

var (
	// PandaUserRole is a globally accessible object for table panda_user_role operations.
	PandaUserRole = pandaUserRoleDao{
		internal.NewPandaUserRoleDao(),
	}
)

// Add your custom methods and functionality below.
