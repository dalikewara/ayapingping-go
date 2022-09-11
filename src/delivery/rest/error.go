package rest

import (
	"github.com/dalikewara/ayapingping-go/v2/src/library/errs"
	"net/http"
)

var ErrorComposeData = errs.NewWithHttpStatus("REST01", "error when compose data", http.StatusInternalServerError)
