package service

import (
	"github.com/dalikewara/ayapingping-go/src/library/errs"
	"net/http"
)

var ErrorRoleSystemUserNotAdmin = errs.NewWithHttpStatus("SVC01", "the user is a system user, but the role is not an admin", http.StatusOK)
