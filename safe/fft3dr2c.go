package safe

import (
	"fmt"
	"github.com/bigbigbigfish/gocuda/cufft"
)

// 3D single-precission real-to-complex FFT plan.
type FFT3DR2CPlan struct {
	fftplan
	size3D
}

// 3D single-precission real-to-complex FFT plan.
func FFT3DR2C(Nx, Ny, Nz int) FFT3DR2CPlan {
	handle := cufft.Plan3d(Nx, Ny, Nz, cufft.R2C)
	handle.SetCompatibilityMode(cufft.COMPATIBILITY_NATIVE)
	return FFT3DR2CPlan{fftplan{handle, 0}, size3D{Nx, Ny, Nz}}
}

// Execute the FFT plan. Synchronized.
// src and dst are 3D arrays stored 1D arrays.
func (p FFT3DR2CPlan) Exec(src Float32s, dst Complex64s) {
	oksrclen := p.InputLen()
	if src.Len() != oksrclen {
		panic(fmt.Errorf("size mismatch: expecting src len %v, got %v", oksrclen, src.Len()))
	}
	okdstlen := p.OutputLen()
	if dst.Len() != okdstlen {
		panic(fmt.Errorf("size mismatch: expecting dst len %v, got %v", okdstlen, dst.Len()))
	}
	p.handle.ExecR2C(src.Pointer(), dst.Pointer())
	p.stream.Synchronize() //!
}

// 3D size of the input array.
func (p FFT3DR2CPlan) InputSize() (Nx, Ny, Nz int) {
	return p.size3D[0], p.size3D[1], p.size3D[2]
}

// 3D size of the output array.
func (p FFT3DR2CPlan) OutputSize() (Nx, Ny, Nz int) {
	return p.size3D[0], p.size3D[1], p.size3D[2]/2 + 1
}

// Required length of the (1D) input array.
func (p FFT3DR2CPlan) InputLen() int {
	return prod3(p.InputSize())
}

// Required length of the (1D) output array.
func (p FFT3DR2CPlan) OutputLen() int {
	return prod3(p.OutputSize())
}
