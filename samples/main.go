package main
import(
	"github.com/bigbigbigfish/gocuda/cu"
	"github.com/bigbigbigfish/gocuda/cublas"
	"fmt"
	"C"
	"strconv"
	"unsafe"
)

const (
	m = 4
	n = 4
	k = 4
)


func main(){
	var alpha_t, beta_t  int32
	alpha_t = 1
	beta_t  = 0
	cu.Init(0);
	fmt.Println("Hello, I am you GPU:", cu.Device(0).Name())
	fmt.Println("Number of devices: " + strconv.Itoa(cu.DeviceGetCount()))
	fmt.Println("Free memory: " + strconv.FormatInt(cu.DeviceGet(0).TotalMem(),10))

	blas :=cublas.NewCublas()
	defer blas.Destroy()
	
	h_A := make([]int8,m*k)
	h_B := make([]int8,k*n)
	h_C := make([]int32,m*n)
	
	for i :=0;i<m*n;i++{
		h_A[i] = int8(i)
		h_B[i] = int8(i)
		h_C[i] = 0
		fmt.Printf("[%d]%d:%d:%d\n",i,h_A[i],h_B[i],h_C[i])
	}
	
	d_A := cu.MemAlloc(m*k)
	d_B := cu.MemAlloc(k*n)
	d_C := cu.MemAlloc(4*m*n)

	cu.MemcpyHtoD(d_A,unsafe.Pointer(&h_A[0]),m*k)
	cu.MemcpyHtoD(d_B,unsafe.Pointer(&h_B[0]),k*n)
//func ( h BlasHandle)GemmEx(transa,transb Type,m,n,k int,alpha,dA unsafe.Pointer,Atype Type,lda int,dB unsafe.Pointer,Btype Type,ldb int,
//                        beta,dC unsafe.Pointer,Ctype Type,ldc int,computype,algo Type){
//	
	blas.GemmEx(cublas.OP_N,cublas.OP_N,m,n,k,unsafe.Pointer(&alpha_t),unsafe.Pointer(d_A),cublas.R_8I,m,unsafe.Pointer(d_B),cublas.R_8I,n,unsafe.Pointer(&beta_t),unsafe.Pointer(d_C),cublas.R_32I,k,cublas.R_32I,cublas.GEMM_DFALT)

	cu.MemcpyDtoH(unsafe.Pointer(&h_C[0]),d_C,4*m*n)
	for i :=0;i<m*n;i++{
		fmt.Printf("[%d]%d\n",i,h_C[i])
	}

	cu.MemFree(d_A)
	cu.MemFree(d_B)
	cu.MemFree(d_C)
}
