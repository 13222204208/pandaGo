// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalPandaRoleDao is an internal type for wrapping the internal DAO implementation.
type internalPandaRoleDao = *internal.PandaRoleDao

// pandaRoleDao is the data access object for the table panda_role.
// You can define custom methods on it to extend its functionality as needed.
type pandaRoleDao struct {
	internalPandaRoleDao
}

var (
	// PandaRole is a globally accessible object for table panda_role operations.
	PandaRole = pandaRoleDao{
		internal.NewPandaRoleDao(),
	}
)

// Add your custom methods and functionality below.
