//authorization provides Role-Based Access Control (RBAC) like functionality
//// in order to restrict resource access to authorised client. It currently has
//// two built-in conditional permission checker types, however it accepts custom
//// ones from outside.

package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sleep-tracker/pkg/errs"
	"sleep-tracker/pkg/httpx"
	"strings"
)

const HeaderXPermissions = "X-Permissions"

type checker interface {
	IsSatisfied(perms string) bool
}

// And requires all permission to be match.
type And struct {
	Permissions []string
}

// IsSatisfied checks if all the required permissions have been present in the
// HTTP request header.
func (a And) IsSatisfied(strPerms string) bool {
	if strPerms == "" || len(a.Permissions) == 0 {
		return false
	}

	perms := strings.Split(strPerms, " ")
	if len(perms) == 0 {
		return false
	}

	listPerms := make(map[string]bool, len(perms))
	for _, perm := range perms {
		listPerms[perm] = true
	}

	for _, perm := range a.Permissions {
		if _, ok := listPerms[perm]; !ok {
			return false
		}
	}

	return true
}

// Or requires at least one permission match.
type Or struct {
	Permissions []string
}

// IsSatisfied checks if at least one of the required permissions has been
// present in the HTTP request header.
func (o Or) IsSatisfied(xPerms string) bool {
	if xPerms == "" || len(o.Permissions) == 0 {
		return false
	}

	perms := strings.Split(xPerms, " ")
	if len(perms) == 0 {
		return false
	}

	listPerms := make(map[string]bool, len(perms))
	for _, perm := range perms {
		listPerms[perm] = true
	}

	for _, perm := range o.Permissions {
		if _, ok := listPerms[perm]; ok {
			return true
		}
	}

	return false
}

// CreateHeaderValue accepts list of permissions and construct a standard
// space separated string value to go with the "X-Permissions" header.
func CreateHeaderValue(perms []string) string {
	return strings.Join(perms, " ")
}

// Authorize - accepts a built-in or a custom checker type and instructs it to
// check if the required permissions were satisfied or not.
func Authorize(c checker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var perms string
		perms, _ = ctx.MustGet(HeaderXPermissions).(string)
		if ok := c.IsSatisfied(perms); !ok {
			err := errs.NewBadRequest("failed to get the permissions for the the service")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, httpx.ApiJson{
				Error: []error{err},
			})
			return
		}

		ctx.Next()
	}
}
