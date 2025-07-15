package constant

const (
	_ = iota
	USER_STATUS_PENDING_ID
	USER_STATUS_ACTIVE_ID
	USER_STATUS_INACTIVE_ID
)

const (
	USER_STATUS_PENDING_SLUG  = "pending"
	USER_STATUS_ACTIVE_SLUG   = "active"
	USER_STATUS_INACTIVE_SLUG = "inactive"
)

const (
	_ = iota
	ROLE_SUPER_ADMIN_ID
	ROLE_ADMIN_ID
	ROLE_USER_ID
)

const (
	ROLE_SUPER_ADMIN_SLUG = "super-admin"
	ROLE_ADMIN_SLUG       = "admin"
	ROLE_USER_SLUG        = "user"
)
