package cublas

//#include <cuda_runtime.h>
//#include <cublas_v2.h>
import "C"

import (
        "github.com/bigbigbigfish/gocuda/cu"
        "unsafe"
)


type BlasHandle struct{
	HandlePtr unsafe.Pointer
	Destroyed bool
}

func NewCublas() *BlasHandle {
        var handle C.cublasHandle_t
	err := Result(C.cublasCreate(&handle))
        if err != SUCCESS {
                panic(err)
        }
        return &BlasHandle{HandlePtr:unsafe.Pointer(handle)}
}

func ( h BlasHandle)GemmEx(transa,transb Type,m,n,k int,alpha,dA unsafe.Pointer,Atype Type,lda int,dB unsafe.Pointer,Btype Type,ldb int,
			beta,dC unsafe.Pointer,Ctype Type,ldc int,computype,algo Type){
	err := Result(C.cublasGemmEx(C.cublasHandle_t(h.HandlePtr),
		      C.cublasOperation_t(transa),C.cublasOperation_t(transb),
		      C.int(m),C.int(n),C.int(k),
		      alpha,dA,C.cudaDataType(Atype), C.int(lda),
		      dB,C.cudaDataType(Btype),
		      C.int(ldb),beta,dC,C.cudaDataType(Ctype),C.int(ldc),
		      C.cudaDataType(computype),C.cublasGemmAlgo_t(algo)))
        if err != SUCCESS {
                panic(err)
        }
}

func (h BlasHandle)SetStream(stream cu.Stream) {
        err := Result(C.cublasSetStream(
                C.cublasHandle_t(h.HandlePtr),
                C.cudaStream_t(unsafe.Pointer(uintptr(stream)))))
        if err != SUCCESS {
                panic(err)
        }
}

func (h BlasHandle) Destroy() {
        err := Result(C.cublasDestroy(C.cublasHandle_t(h.HandlePtr)))
        if err != SUCCESS {
                panic(err)
        }
}

