package cublas

//#include <cuda_runtime.h>
//#include <cublas_v2.h>
//#include "library_types.h"
import "C"

import (
//	"fmt"
)

// FFT type
type Type int

const (
	OP_N Type = C.CUBLAS_OP_N
	OP_T Type = C.CUBLAS_OP_T
	OP_C Type = C.CUBLAS_OP_C
)

const (
	GEMM_DFALT Type   = C.CUBLAS_GEMM_DFALT
	GEMM_DEFAULT Type = C.CUBLAS_GEMM_DEFAULT
	GEMM_ALGO0 Type = C.CUBLAS_GEMM_ALGO0
	GEMM_ALGO1 Type = C.CUBLAS_GEMM_ALGO1
	GEMM_ALGO2 Type = C.CUBLAS_GEMM_ALGO2
	GEMM_ALGO3 Type = C.CUBLAS_GEMM_ALGO3
	GEMM_ALGO4 Type = C.CUBLAS_GEMM_ALGO4
	GEMM_ALGO5 Type = C.CUBLAS_GEMM_ALGO5
	GEMM_ALGO6 Type = C.CUBLAS_GEMM_ALGO6
	GEMM_ALGO7 Type = C.CUBLAS_GEMM_ALGO7
)

const (
	R_16F Type = C.CUDA_R_16F
        C_16F Type = C.CUDA_C_16F
        R_32F Type = C.CUDA_R_32F
        C_32F Type = C.CUDA_C_32F
        R_64F Type = C.CUDA_R_64F
        C_64F Type = C.CUDA_C_64F
        R_8I  Type = C.CUDA_R_8I
        C_8I  Type = C.CUDA_C_8I
        R_8U  Type = C.CUDA_R_8U
        C_8U  Type = C.CUDA_C_8U
        R_32I Type = C.CUDA_R_32I
        C_32I Type = C.CUDA_C_32I
        R_32U Type = C.CUDA_R_32U
        C_32U Type = C.CUDA_C_32U
)

