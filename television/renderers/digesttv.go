package renderers

import (
	"crypto/sha1"
	"fmt"
	"gopher2600/errors"
	"gopher2600/hardware/tia/audio"
	"gopher2600/television"
	"os"
	"strings"
)

// DigestTV is an implementation of the television.Renderer interface with an
// embedded television for convenience. It generates a sha1 value of the
// image every frame. it does not display the image anywhere.
//
// Note that the use of sha1 is fine for this application because this is not a
// cryptographic task.
type DigestTV struct {
	television.Television
	digest    [sha1.Size]byte
	frameData []byte
	frameNum  int
}

// NewDigestTV initialises a new instance of DigestTV. For convenience, the
// television argument can be nil, in which case an instance of
// StellaTelevision will be created.
func NewDigestTV(tvType string, tv television.Television) (*DigestTV, error) {
	var err error

	// set up digest tv
	dtv := new(DigestTV)

	// create or attach television implementation
	if tv == nil {
		dtv.Television, err = television.NewStellaTelevision(tvType)
		if err != nil {
			return nil, err
		}
	} else {
		// check that the quoted tvType matches the specification of the
		// supplied BasicTelevision instance. we don't really need this but
		// becuase we're implying that tvType is required, even when an
		// instance of BasicTelevision has been supplied, the caller may be
		// expecting an error
		tvType = strings.ToUpper(tvType)
		if tvType != "AUTO" && tvType != tv.GetSpec().ID {
			return nil, errors.New(errors.DigestTV, "trying to piggyback a tv of a different spec")
		}
		dtv.Television = tv
	}

	// register ourselves as a television.Renderer
	dtv.AddPixelRenderer(dtv)

	// set attributes that depend on the television specification
	dtv.ChangeTVSpec()

	return dtv, nil
}

// ChangeTVSpec implements television.Television interface
func (dtv *DigestTV) ChangeTVSpec() error {
	// memory for frameData has to be sufficient for the entirety of the
	// screen plus the size of a fingerprint. we'll use the additional space to
	// chain fingerprint hashes
	dtv.frameData = make([]byte, len(dtv.digest)+((television.ClocksPerScanline+1)*(dtv.GetSpec().ScanlinesTotal+1)*3))
	return nil
}

// NewFrame implements television.Renderer interface
func (dtv *DigestTV) NewFrame(frameNum int) error {
	// chain fingerprints by copying the value of the last fingerprint
	// to the head of the screen data
	n := copy(dtv.frameData, dtv.digest[:])
	if n != len(dtv.digest) {
		return errors.New(errors.DigestTV, fmt.Sprintf("unexpected amount of data copied"))
	}
	dtv.digest = sha1.Sum(dtv.frameData)
	dtv.frameNum = frameNum
	return nil
}

// NewScanline implements television.Renderer interface
func (dtv *DigestTV) NewScanline(scanline int) error {
	return nil
}

// SetPixel implements television.Renderer interface
func (dtv *DigestTV) SetPixel(x, y int, red, green, blue byte, vblank bool) error {
	// preserve the first few bytes for a chained fingerprint
	offset := television.ClocksPerScanline * y * 3
	offset += x * 3

	if offset >= len(dtv.frameData) {
		return errors.New(errors.DigestTV, fmt.Sprintf("the coordinates (%d, %d) passed to SetPixel will cause an invalid access of the frameData array", x, y))
	}

	dtv.frameData[offset] = red
	dtv.frameData[offset+1] = green
	dtv.frameData[offset+2] = blue

	return nil
}

// SetAltPixel implements television.Renderer interface
func (dtv *DigestTV) SetAltPixel(x, y int, red, green, blue byte, vblank bool) error {
	return nil
}

// SetAudio implements television.AudioMixer interface
func (dtv *DigestTV) SetAudio(audio audio.Audio) error {
	return nil
}

func (dtv DigestTV) String() string {
	return fmt.Sprintf("%x", dtv.digest)
}

// ResetDigest resets the current digest value to 0
func (dtv *DigestTV) ResetDigest() {
	for i := range dtv.digest {
		dtv.digest[i] = 0
	}
}

// Save current frame data to filename - filename base supplied as an argument, the
// frame number and file extension is appended by the function
func (dtv *DigestTV) Save(fileNameBase string) error {
	// prepare filename for image
	outName := fmt.Sprintf("%s_digest_%d.bin", fileNameBase, dtv.frameNum)

	f, err := os.Open(outName)
	if f != nil {
		f.Close()
		return errors.New(errors.DigestTV, fmt.Sprintf("output file (%s) already exists", outName))
	}
	if err != nil && !os.IsNotExist(err) {
		return errors.New(errors.DigestTV, err)
	}

	f, err = os.Create(outName)
	if err != nil {
		return errors.New(errors.DigestTV, err)
	}
	defer f.Close()

	n, err := f.Write(dtv.frameData)
	if n != len(dtv.frameData) {
		return errors.New(errors.DigestTV, "output truncated")
	}
	if err != nil {
		return errors.New(errors.DigestTV, err)
	}

	return nil
}
