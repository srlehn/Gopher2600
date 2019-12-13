package cartridge

import (
	"crypto/sha1"
	"fmt"
	"gopher2600/cartridgeloader"
	"gopher2600/errors"
	"gopher2600/hardware/memory/bus"
	"gopher2600/hardware/memory/memorymap"
	"strings"
)

// Cartridge defines the information and operations for a VCS cartridge
type Cartridge struct {
	bus.DebuggerBus
	bus.CPUBus

	Filename string
	Hash     string

	// the specific cartridge data, mapped appropriately to the memory
	// interfaces
	mapper cartMapper
}

// NewCartridge is the preferred method of initialisation for the cartridge
// type
func NewCartridge() *Cartridge {
	cart := &Cartridge{}
	cart.Eject()
	return cart
}

func (cart Cartridge) String() string {
	return fmt.Sprintf("%s [%s]", cart.Filename, cart.mapper)
}

// Peek is an implementation of memory.DebuggerBus
func (cart Cartridge) Peek(addr uint16) (uint8, error) {
	addr &= memorymap.OriginCart - 1
	return cart.mapper.read(addr)
}

// Poke is an implementation of memory.DebuggerBus
func (cart Cartridge) Poke(addr uint16, data uint8) error {
	return errors.New(errors.UnpokeableAddress, addr)
}

// Read is an implementation of memory.CPUBus
func (cart *Cartridge) Read(addr uint16) (uint8, error) {
	// * optimisation: called a lot. pointer to Cartridge to prevent duffcopy

	addr &= memorymap.OriginCart - 1
	return cart.mapper.read(addr)
}

// Write is an implementation of memory.CPUBus
func (cart *Cartridge) Write(addr uint16, data uint8) error {
	addr &= memorymap.OriginCart - 1
	return cart.mapper.write(addr, data)
}

// Eject removes memory from cartridge space and unlike the real hardware,
// attaches a bank of empty memory - for convenience of the debugger
func (cart *Cartridge) Eject() {
	cart.Filename = ejectedName
	cart.Hash = ejectedHash
	cart.mapper = newEjected()
}

// IsEjected returns true if no cartridge is attached
func (cart *Cartridge) IsEjected() bool {
	return cart.Hash == ejectedHash
}

// Attach the cartridge loader to the VCS and make available the data to the CPU
// bus
func (cart *Cartridge) Attach(cartload cartridgeloader.Loader) error {
	data, err := cartload.Load()
	if err != nil {
		return err
	}

	// note name of cartridge
	cart.Filename = cartload.Filename
	cart.mapper = newEjected()

	// generate hash
	cart.Hash = fmt.Sprintf("%x", sha1.Sum(data))

	// check that the hash matches the expected value
	if cartload.Hash != "" && cartload.Hash != cart.Hash {
		return errors.New(errors.CartridgeError, "unexpected hash value")
	}

	// how cartridges are mapped into the 4k space can differs dramatically.
	// the following implementation details have been cribbed from Kevin
	// Horton's "Cart Information" document [sizes.txt]

	cartload.Format = strings.ToUpper(cartload.Format)

	if cartload.Format == "" || cartload.Format == "AUTO" {
		return cart.fingerprint(data)
	}

	addSuperchip := false

	switch cartload.Format {
	case "2k":
		cart.mapper, err = newAtari2k(data)
	case "4k":
		cart.mapper, err = newAtari4k(data)
	case "F8":
		cart.mapper, err = newAtari8k(data)
	case "F6":
		cart.mapper, err = newAtari16k(data)
	case "F4":
		cart.mapper, err = newAtari32k(data)

	case "2k+SC":
		cart.mapper, err = newAtari2k(data)
		addSuperchip = true
	case "4k+SC":
		cart.mapper, err = newAtari4k(data)
		addSuperchip = true
	case "F8+SC":
		cart.mapper, err = newAtari8k(data)
		addSuperchip = true
	case "F6+SC":
		cart.mapper, err = newAtari16k(data)
		addSuperchip = true
	case "F4+SC":
		cart.mapper, err = newAtari32k(data)
		addSuperchip = true

	case "FA":
		cart.mapper, err = newCBS(data)
	case "FE":
		// TODO
	case "E0":
		cart.mapper, err = newparkerBros(data)
	case "E7":
		cart.mapper, err = newMnetwork(data)
	case "3F":
		cart.mapper, err = newTigervision(data)
	case "AR":
		// TODO
	}

	if addSuperchip {
		if superchip, ok := cart.mapper.(optionalSuperchip); ok {
			if !superchip.addSuperchip() {
				err = errors.New(errors.CartridgeError, "error adding superchip")
			}
		} else {
			err = errors.New(errors.CartridgeError, "error adding superchip")
		}
	}

	return err
}

// Initialise the cartridge
func (cart *Cartridge) Initialise() {
	cart.mapper.initialise()
}

// NumBanks returns the number of banks in the catridge
func (cart Cartridge) NumBanks() int {
	return cart.mapper.numBanks()
}

// GetBank returns the current bank number for the specified address
func (cart Cartridge) GetBank(addr uint16) int {
	addr &= memorymap.OriginCart - 1
	return cart.mapper.getBank(addr)
}

// SetBank maps the specified address such that it references the specified
// bank. For many cart mappers this just means switching banks for the entire
// cartridge
func (cart *Cartridge) SetBank(addr uint16, bank int) error {
	addr &= memorymap.OriginCart - 1
	return cart.mapper.setBank(addr, bank)
}

// SaveState notes and returns the current state of the cartridge (RAM
// contents, selected bank)
func (cart *Cartridge) SaveState() interface{} {
	return cart.mapper.saveState()
}

// RestoreState retuns the state of the cartridge to a previously known state
func (cart *Cartridge) RestoreState(state interface{}) error {
	return cart.mapper.restoreState(state)
}

// RAM returns a read only instance of any cartridge RAM
func (cart Cartridge) RAM() []uint8 {
	return cart.mapper.ram()
}

// Listen for data at the specified address. very wierd requirement of the
// tigervision cartridge format. If there was a better way of implementing the
// tigervision format, there'd be no need for this function.
func (cart Cartridge) Listen(addr uint16, data uint8) {
	cart.mapper.listen(addr, data)
}