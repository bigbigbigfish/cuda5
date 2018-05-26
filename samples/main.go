package main
import(
	"github.com/bigbigbigfish/gocuda/cu"
	"github.com/bigbigbigfish/gocuda/cublas"
	"fmt"
	"C"
	"strconv"
)

func main(){
	cu.Init(0);
	fmt.Println("Hello, I am you GPU:", cu.Device(0).Name())
	fmt.Println("Number of devices: " + strconv.Itoa(cu.DeviceGetCount()))
	fmt.Println("Free memory: " + strconv.FormatInt(cu.DeviceGet(0).TotalMem(),10))

	blas :=cublas.NewCublas()
	defer blas.Destroy()

	

}
