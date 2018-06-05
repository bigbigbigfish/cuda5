package cublas

// This file provides access to CUDA driver error statuses (type CUresult).

//#include <cublas_v2.h>
import "C"

import (
	"fmt"
)

// CUDA error status.
// CUDA error statuses are not returned by functions but checked and passed to
// panic() when not successful. If desired, they can be caught by
// recover().
type Result int

// Message string for the error
func (err Result) String() string {
	str, ok := errorString[err]
	if !ok {
		return "Unknown CUresult: " + fmt.Sprint(int(err))
	}
	return str
}

const (
	SUCCESS                              Result = C.CUBLAS_STATUS_SUCCESS
	ERROR_NOT_INITIALIZED                Result = C.CUBLAS_STATUS_NOT_INITIALIZED
	ERROR_ALLOC_FAILED                   Result = C.CUBLAS_STATUS_ALLOC_FAILED
	ERROR_INVALID_VALUE                  Result = C.CUBLAS_STATUS_INVALID_VALUE
	ERROR_ARCH_MISMATCH                  Result = C.CUBLAS_STATUS_ARCH_MISMATCH
	ERROR_MAPPING_ERROR                  Result = C.CUBLAS_STATUS_MAPPING_ERROR
	ERROR_EXECUTION_FAILED               Result = C.CUBLAS_STATUS_EXECUTION_FAILED
	ERROR_INTERNAL_ERROR                 Result = C.CUBLAS_STATUS_INTERNAL_ERROR
	ERROR_NOT_SUPPORTED                  Result = C.CUBLAS_STATUS_NOT_SUPPORTED
	ERROR_LICENSE_ERROR                  Result = C.CUBLAS_STATUS_LICENSE_ERROR
)

// Map with error strings for Result error numbers
var errorString map[Result]string = map[Result]string{
	SUCCESS:                             "CUBLAS_STATUS_SUCCESS",
	ERROR_NOT_INITIALIZED:                "CUBLAS_STATUS_NOT_INITIALIZED",
	ERROR_ALLOC_FAILED:                   "CUBLAS_STATUS_ALLOC_FAILED",
	ERROR_INVALID_VALUE:                  "CUBLAS_STATUS_INVALID_VALUE",
	ERROR_ARCH_MISMATCH:                  "CUBLAS_STATUS_ARCH_MISMATCH",
	ERROR_MAPPING_ERROR:                  "CUBLAS_STATUS_MAPPING_ERROR",
	ERROR_EXECUTION_FAILED:               "CUBLAS_STATUS_EXECUTION_FAILED",
	ERROR_INTERNAL_ERROR:                 "CUBLAS_STATUS_INTERNAL_ERROR",
	ERROR_NOT_SUPPORTED:                  "CUBLAS_STATUS_NOT_SUPPORTED",
	ERROR_LICENSE_ERROR:                  "CUBLAS_STATUS_LICENSE_ERROR"}
