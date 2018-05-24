package safe

// INTERNAL
// Base implementation for all FFT plans.

import (
	"github.com/bigbigbigfish/gocuda/cu"
	"github.com/bigbigbigfish/gocuda/cufft"
)

// Base implementation for all FFT plans.
type fftplan struct {
	handle cufft.Handle
	stream cu.Stream
}

// For the sake of embedding.
type size1D int

// Returns the logical size of the FFT:
// the number of elements (real or complex)
// it transforms.
func (s size1D) Size() int { return int(s) }

// For the sake of embedding.
type size3D [3]int

// Returns the logical size of the FFT:
// the number of elements (real or complex)
// it transforms.
func (s size3D) Size() (Nx, Ny, Nz int) { return s[0], s[1], s[2] }

func prod3(x, y, z int) int {
	return x * y * z
}

// Releases all resources associated with the FFT plan.
func (p fftplan) Destroy() { p.handle.Destroy() }

// Associates a CUDA stream with the FFT plan.
// If a stream is set, plan.Stream().Synchronize() can
// to be called to wait for the execution to finish.
func (p fftplan) SetStream(stream cu.Stream) {
	p.handle.SetStream(stream)
	p.stream = stream
}

// Returns the CUDA stream associated with the FFT plan.
func (p fftplan) Stream() cu.Stream {
	return p.stream
}
