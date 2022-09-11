package repository

import (
	"github.com/dalikewara/ayapingping-go/v2/src/library/errs"
	"net/http"
)

var ErrorRoleNoDataFound = errs.NewWithHttpStatus("RP01", "no role data found", http.StatusOK)
