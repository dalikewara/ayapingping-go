package constant

import "net/http"

const SystemUserRole = "admin"

var SystemUserIds = []int{1, 2, 3}

const RESTOkCode = "00"
const RESTOkMessage = "Ok"
const RESTOkHttpStatus = http.StatusOK
const RESTNotOkHttpStatus = http.StatusInternalServerError
