package utils

import authz "github.com/zeiss/fiber-authz"

const (
	// PermissionAdmin grants all permissions on a team
	PermissionAdmin = authz.AuthzAction("admin")
	// PermissionSuperAdmin grants all permissions
	PermissionSuperAdmin = authz.AuthzAction("superadmin")
	// PermissionCreate grants the ability to create
	PermissionCreate = authz.AuthzAction("create")
	// PermissionDelete grants the ability to delete
	PermissionDelete = authz.AuthzAction("delete")
	// PermissionEdit grants the ability to edit
	PermissionEdit = authz.AuthzAction("edit")
	// PermissionView grants the ability to read
	PermissionView = authz.AuthzAction("view")
)

const (
	// RoleAdmin grants all permissions on a team
	RoleAdmin = authz.AuthzAction("Admin")
	// RoleSuperAdmin grants all permissions
	RoleSuperAdmin = authz.AuthzAction("Super Admin")
	// RoleOwner grants all permissions
	RoleOwner = authz.AuthzAction("Owner")
	// Editor grants the ability to edit
	RoleEditor = authz.AuthzAction("Editor")
	// RoleViewer grants the ability to read
	RoleViewer = authz.AuthzAction("Viewer")
)
