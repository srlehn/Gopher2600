package video

import (
	"gopher2600/hardware/memory"
	"gopher2600/hardware/memory/vcssymbols"
	"gopher2600/hardware/tia/polycounter"
	"gopher2600/hardware/tia/video/future"
)

// Video contains all the components of the video sub-system of the VCS TIA chip
type Video struct {
	colorClock *polycounter.Polycounter

	vblank *bool

	// collision matrix
	collisions *collisions

	// playfield
	Playfield *playfield

	// sprite objects
	Player0  *playerSprite
	Player1  *playerSprite
	Missile0 *missileSprite
	Missile1 *missileSprite
	Ball     *ballSprite

	// there's a slight delay when changing the state of video objects. we're
	// using two future instances to emulate what happens in the 2600. the
	// first is OnFutureColorClock, which *ticks* every video cycle. we use this for
	// writing playfield bits, player bits and enable flags for missiles and
	// the ball.
	//
	// the second future instance is OnFutureMotionClock. this is for those
	// writes that only occur during the "motion clock". eg. resetting sprite
	// positions
	OnFutureColorClock  future.Group
	OnFutureMotionClock future.Group
}

// colors to use for debugging
const (
	debugColBackground = uint8(0x02) // black (stella uses a light grey)
	debugColBall       = uint8(0xb4) // cyan
	debugColPlayfield  = uint8(0x62) // purple
	debugColPlayer0    = uint8(0x32) // red
	debugColPlayer1    = uint8(0x12) // gold
	debugColMissile0   = uint8(0xf2) // orange
	debugColMissile1   = uint8(0xd2) // green
)

// NewVideo is the preferred method of initialisation for the Video structure
func NewVideo(colorClock *polycounter.Polycounter, mem memory.ChipBus, vblank *bool) *Video {
	vd := &Video{colorClock: colorClock, vblank: vblank}

	// collision matrix
	vd.collisions = newCollision(mem)

	// playfield
	vd.Playfield = newPlayfield(vd.colorClock)

	// sprite objects
	vd.Player0 = newPlayerSprite("player0", vd.colorClock)
	if vd.Player0 == nil {
		return nil
	}
	vd.Player1 = newPlayerSprite("player1", vd.colorClock)
	if vd.Player1 == nil {
		return nil
	}
	vd.Missile0 = newMissileSprite("missile0", vd.colorClock)
	if vd.Missile0 == nil {
		return nil
	}
	vd.Missile1 = newMissileSprite("missile1", vd.colorClock)
	if vd.Missile1 == nil {
		return nil
	}
	vd.Ball = newBallSprite("ball", vd.colorClock)
	if vd.Ball == nil {
		return nil
	}

	// connect player 0 and player 1 to each other
	vd.Player0.otherPlayer = vd.Player1
	vd.Player1.otherPlayer = vd.Player0

	// connect missile sprite to its parent player sprite
	vd.Missile0.parentPlayer = vd.Player0
	vd.Missile1.parentPlayer = vd.Player1

	return vd
}

// TickFutures is called *every* video clock
func (vd *Video) TickFutures(sprites bool) {
	// resolve delayed write operations
	vd.OnFutureColorClock.Tick()

	if sprites {
		vd.OnFutureMotionClock.Tick()
	}
}

// TickPlayfield is called *every* video clock
func (vd *Video) TickPlayfield() {
	// tick playfield forward
	vd.Playfield.tick()
}

// TickSprites moves all video elements forward one video cycle and is only
// called when motion clock is active
func (vd *Video) TickSprites() {
	vd.Player0.tick()
	vd.Player1.tick()
	vd.Missile0.tick()
	vd.Missile1.tick()
	vd.Ball.tick()
}

// PrepareSpritesForHMOVE should be called whenever HMOVE is triggered
func (vd *Video) PrepareSpritesForHMOVE() {
	vd.Player0.prepareForHMOVE()
	vd.Player1.prepareForHMOVE()
	vd.Missile0.prepareForHMOVE()
	vd.Missile1.prepareForHMOVE()
	vd.Ball.prepareForHMOVE()
}

// EndHMOVE is called whenever HMOVE activity is manually unset
func (vd *Video) EndHMOVE() {
	vd.Player0.endHMOVE()
	vd.Player1.endHMOVE()
	vd.Missile0.endHMOVE()
	vd.Missile1.endHMOVE()
	vd.Ball.endHMOVE()
}

// ResolveHorizMovement is only called when HMOVE is active
func (vd *Video) ResolveHorizMovement(count int) {
	vd.Player0.resolveHMOVE(count)
	vd.Player1.resolveHMOVE(count)
	vd.Missile0.resolveHMOVE(count)
	vd.Missile1.resolveHMOVE(count)
	vd.Ball.resolveHMOVE(count)
}

// ForceHMOVE is an ungodly hack
func (vd *Video) ForceHMOVE(adjustment int) {
	vd.Player0.forceHMOVE(adjustment)
	vd.Player1.forceHMOVE(adjustment)
	vd.Missile0.forceHMOVE(adjustment)
	vd.Missile1.forceHMOVE(adjustment)
	vd.Ball.forceHMOVE(adjustment)
}

// Pixel returns the color of the pixel at the current time. it will default
// to returning the background color if no sprite or playfield pixel is
// present. it also sets the collision registers
// - it need not be called therefore during VBLANK or HBLANK
func (vd *Video) Pixel() (uint8, uint8) {
	bgc := vd.Playfield.backgroundColor
	pfu, pfc := vd.Playfield.pixel()
	blu, blc := vd.Ball.pixel()
	p0u, p0c := vd.Player0.pixel()
	p1u, p1c := vd.Player1.pixel()
	m0u, m0c := vd.Missile0.pixel()
	m1u, m1c := vd.Missile1.pixel()

	// collisions
	if m0u && p1u {
		vd.collisions.cxm0p |= 0x80
		vd.collisions.SetMemory(vcssymbols.CXM0P)
	}
	if m0u && p0u {
		vd.collisions.cxm0p |= 0x40
		vd.collisions.SetMemory(vcssymbols.CXM0P)
	}

	if m1u && p0u {
		vd.collisions.cxm1p |= 0x80
		vd.collisions.SetMemory(vcssymbols.CXM1P)
	}
	if m1u && p1u {
		vd.collisions.cxm1p |= 0x40
		vd.collisions.SetMemory(vcssymbols.CXM1P)
	}

	if p0u && pfu {
		vd.collisions.cxp0fb |= 0x80
		vd.collisions.SetMemory(vcssymbols.CXP0FB)
	}
	if p0u && blu {
		vd.collisions.cxp0fb |= 0x40
		vd.collisions.SetMemory(vcssymbols.CXP0FB)
	}

	if p1u && pfu {
		vd.collisions.cxp1fb |= 0x80
		vd.collisions.SetMemory(vcssymbols.CXP1FB)
	}
	if p1u && blu {
		vd.collisions.cxp1fb |= 0x40
		vd.collisions.SetMemory(vcssymbols.CXP1FB)
	}

	if m0u && pfu {
		vd.collisions.cxm0fb |= 0x80
		vd.collisions.SetMemory(vcssymbols.CXM0FB)
	}
	if m0u && blu {
		vd.collisions.cxm0fb |= 0x40
		vd.collisions.SetMemory(vcssymbols.CXM0FB)
	}

	if m1u && pfu {
		vd.collisions.cxm1fb |= 0x80
		vd.collisions.SetMemory(vcssymbols.CXM1FB)
	}
	if m1u && blu {
		vd.collisions.cxm1fb |= 0x40
		vd.collisions.SetMemory(vcssymbols.CXM1FB)
	}

	if blu && pfu {
		vd.collisions.cxblpf |= 0x80
		vd.collisions.SetMemory(vcssymbols.CXBLPF)
	}
	// no bit 6 for CXBLPF

	if p0u && p1u {
		vd.collisions.cxppmm |= 0x80
		vd.collisions.SetMemory(vcssymbols.CXPPMM)
	}
	if m0u && m1u {
		vd.collisions.cxppmm |= 0x40
		vd.collisions.SetMemory(vcssymbols.CXPPMM)
	}

	var col, dcol uint8

	// apply priorities to get pixel color
	if vd.Playfield.priority {
		if pfu { // priority 1
			col = pfc
			dcol = debugColPlayfield
		} else if blu {
			col = blc
			dcol = debugColBall
		} else if p0u { // priority 2
			col = p0c
			dcol = debugColPlayer0
		} else if m0u {
			col = m0c
			dcol = debugColMissile0
		} else if p1u { // priority 3
			col = p1c
			dcol = debugColPlayer1
		} else if m1u {
			col = m1c
			dcol = debugColMissile1
		} else {
			col = bgc
			dcol = debugColBackground
		}

	} else {
		if p0u { // priority 1
			col = p0c
			dcol = debugColPlayer0
		} else if m0u {
			col = m0c
			dcol = debugColMissile0
		} else if p1u { // priority 2
			col = p1c
			dcol = debugColPlayer1
		} else if m1u {
			col = m1c
			dcol = debugColMissile1
		} else if blu { // priority 3
			col = blc
			dcol = debugColBall
		} else if pfu {
			if vd.Playfield.scoremode == true {
				if vd.Playfield.screenRegion == 2 {
					col = p1c
				} else {
					col = p0c
				}
			} else {
				col = pfc
			}
			dcol = debugColPlayfield
		} else {
			col = bgc
			dcol = debugColBackground
		}
	}

	// priority 4
	return col, dcol
}

func createTriggerList(playerSize uint8) []int {
	var triggerList []int
	switch playerSize {
	case 0x0, 0x05, 0x07:
		// empty trigger list - single sprite of varying widths
	case 0x01:
		triggerList = []int{4}
	case 0x02:
		triggerList = []int{8}
	case 0x03:
		triggerList = []int{4, 8}
	case 0x04:
		triggerList = []int{16}
	case 0x06:
		triggerList = []int{8, 16}
	}
	return triggerList
}

// ReadVideoMemory checks the TIA memory for changes to registers that are
// interesting to the video sub-system. all changes happen immediately except
// for those where a "schedule" function is called.
//
// returns true if memory has been serviced
func (vd *Video) ReadVideoMemory(register string, value uint8) bool {
	switch register {
	default:
		return false

	case "VBLANK":
		vd.OnFutureColorClock.Schedule(delayVBLANK, func() {
			*vd.vblank = (value&0x02 == 0x02)
		}, "setting VBLANK")

	// colours
	case "COLUP0":
		// TODO: write delay?
		vd.Player0.color = value & 0xfe
		vd.Missile0.color = value & 0xfe
	case "COLUP1":
		// TODO: write delay?
		vd.Player1.color = value & 0xfe
		vd.Missile1.color = value & 0xfe

	// playfield
	case "COLUBK":
		// this delay works and fixes a graphical issue with the "Keystone
		// Kapers" rom. I'm not entirely sure this is the correct fix however.
		// and I'm definitely not sure about the delay time
		vd.OnFutureColorClock.Schedule(delayWritePlayfield, func() {
			vd.Playfield.backgroundColor = value & 0xfe
		}, "setting COLUBK")
	case "COLUPF":
		// similar to COLUBK this fixes a bug with "Pressure Cooker"
		vd.OnFutureColorClock.Schedule(delayWritePlayfield, func() {
			vd.Playfield.foregroundColor = value & 0xfe
			vd.Ball.color = value & 0xfe
		}, "setting COLUPF")
	case "CTRLPF":
		// TODO: write delay?
		vd.Ball.size = (value & 0x30) >> 4
		vd.Playfield.reflected = value&0x01 == 0x01
		vd.Playfield.scoremode = value&0x02 == 0x02
		vd.Playfield.priority = value&0x04 == 0x04
	case "REFP0":
		// TODO: write delay?
		vd.Player0.reflected = value&0x08 == 0x08
	case "REFP1":
		// TODO: write delay?
		vd.Player1.reflected = value&0x08 == 0x08
	case "PF0":
		vd.Playfield.scheduleWrite(0, value, &vd.OnFutureColorClock)
	case "PF1":
		vd.Playfield.scheduleWrite(1, value, &vd.OnFutureColorClock)
	case "PF2":
		vd.Playfield.scheduleWrite(2, value, &vd.OnFutureColorClock)

	// ball sprite
	case "ENABL":
		vd.Ball.scheduleEnable(value&0x02 == 0x02, &vd.OnFutureColorClock)
	case "RESBL":
		vd.Ball.scheduleReset(&vd.OnFutureMotionClock)
	case "VDELBL":
		vd.Ball.scheduleVerticalDelay(value&0x01 == 0x01, &vd.OnFutureMotionClock)

	// player sprites
	case "GRP0":
		vd.Player0.scheduleWrite(value, &vd.OnFutureColorClock)
	case "GRP1":
		vd.Player1.scheduleWrite(value, &vd.OnFutureColorClock)
	case "RESP0":
		vd.Player0.scheduleReset(&vd.OnFutureMotionClock)
	case "RESP1":
		vd.Player1.scheduleReset(&vd.OnFutureMotionClock)
	case "VDELP0":
		vd.Player0.scheduleVerticalDelay(value&0x01 == 0x01, &vd.OnFutureMotionClock)
	case "VDELP1":
		vd.Player1.scheduleVerticalDelay(value&0x01 == 0x01, &vd.OnFutureMotionClock)

	// missile sprites
	case "ENAM0":
		vd.Missile0.scheduleEnable(value&0x02 == 0x02, &vd.OnFutureColorClock)
	case "ENAM1":
		vd.Missile1.scheduleEnable(value&0x02 == 0x02, &vd.OnFutureColorClock)
	case "RESM0":
		vd.Missile0.scheduleReset(&vd.OnFutureMotionClock)
	case "RESM1":
		vd.Missile1.scheduleReset(&vd.OnFutureMotionClock)
	case "RESMP0":
		vd.Missile0.scheduleResetToPlayer(value&0x02 == 0x002, &vd.OnFutureColorClock)
	case "RESMP1":
		vd.Missile1.scheduleResetToPlayer(value&0x02 == 0x002, &vd.OnFutureColorClock)

	// player & missile sprites
	case "NUSIZ0":
		vd.OnFutureColorClock.Schedule(delayNUSIZ, func() {
			vd.Missile0.size = (value & 0x30) >> 4
			vd.Player0.size = value & 0x07
			vd.Player0.triggerList = createTriggerList(vd.Player0.size)
			vd.Missile0.triggerList = vd.Player0.triggerList
		}, "adjusting NUSIZ0")
	case "NUSIZ1":
		vd.OnFutureColorClock.Schedule(delayNUSIZ, func() {
			vd.Missile1.size = (value & 0x30) >> 4
			vd.Player1.size = value & 0x07
			vd.Player1.triggerList = createTriggerList(vd.Player1.size)
			vd.Missile1.triggerList = vd.Player1.triggerList
		}, "adjusting NUSIZ1")

	// clear collisions
	case "CXCLR":
		vd.collisions.clear()

	// horizontal movement
	case "HMCLR":
		vd.Player0.horizMovement = 0x08
		vd.Player1.horizMovement = 0x08
		vd.Missile0.horizMovement = 0x08
		vd.Missile1.horizMovement = 0x08
		vd.Ball.horizMovement = 0x08

	// horizontal movement values range from -8 to +7 for convenience we
	// convert this to the range 0 to 15

	case "HMP0":
		// TODO: write delay?
		vd.Player0.horizMovement = (int(value) ^ 0x80) >> 4
	case "HMP1":
		// TODO: write delay?
		vd.Player1.horizMovement = (int(value) ^ 0x80) >> 4
	case "HMM0":
		// TODO: write delay?
		vd.Missile0.horizMovement = (int(value) ^ 0x80) >> 4
	case "HMM1":
		// TODO: write delay?
		vd.Missile1.horizMovement = (int(value) ^ 0x80) >> 4
	case "HMBL":
		// TODO: write delay?
		vd.Ball.horizMovement = (int(value) ^ 0x80) >> 4
	}

	return true
}
