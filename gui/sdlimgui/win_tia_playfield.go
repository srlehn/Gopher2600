// This file is part of Gopher2600.
//
// Gopher2600 is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Gopher2600 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Gopher2600.  If not, see <https://www.gnu.org/licenses/>.
//
// *** NOTE: all historical versions of this file, as found in any
// git repository, are also covered by the licence, even when this
// notice is not present ***

package sdlimgui

import (
	"fmt"
	"strconv"

	"github.com/jetsetilly/gopher2600/hardware/tia/video"

	"github.com/inkyblackness/imgui-go/v2"
)

func (win *winTIA) drawPlayfield() {
	lz := win.img.lazy.Playfield
	pf := win.img.lazy.VCS.TIA.Video.Playfield
	bl := win.img.lazy.VCS.TIA.Video.Ball

	imgui.Spacing()

	// foreground color indicator. when clicked popup palette is requested. on
	// selection in palette, missile color is changed too
	imgui.BeginGroup()

	imguiText("Foreground")
	fgCol := lz.ForegroundColor
	if win.img.imguiSwatch(fgCol) {
		win.popupPalette.request(&fgCol, func() {
			win.img.lazy.Dbg.PushRawEvent(func() { pf.ForegroundColor = fgCol })
			// update ball color too
			win.img.lazy.Dbg.PushRawEvent(func() { bl.Color = fgCol })
		})
	}

	// background color indicator. when clicked popup palette is requested
	imguiText("Background")
	bgCol := lz.BackgroundColor
	if win.img.imguiSwatch(bgCol) {
		win.popupPalette.request(&bgCol, func() {
			win.img.lazy.Dbg.PushRawEvent(func() { pf.BackgroundColor = bgCol })
		})
	}
	imgui.EndGroup()

	imgui.Spacing()
	imgui.Spacing()

	// playfield control bits
	imgui.BeginGroup()
	imguiText("Reflected")
	ref := lz.Reflected
	if imgui.Checkbox("##reflected", &ref) {
		win.img.lazy.Dbg.PushRawEvent(func() {
			pf.Reflected = ref
			win.img.lazy.VCS.TIA.Video.UpdateCTRLPF()
		})
	}
	imgui.SameLine()
	imguiText("Scoremode")
	sm := lz.Scoremode
	if imgui.Checkbox("##scoremode", &sm) {
		win.img.lazy.Dbg.PushRawEvent(func() {
			pf.Scoremode = sm
			win.img.lazy.VCS.TIA.Video.UpdateCTRLPF()
		})
	}
	imgui.SameLine()
	imguiText("Priority")
	pri := lz.Priority
	if imgui.Checkbox("##priority", &pri) {
		win.img.lazy.Dbg.PushRawEvent(func() {
			pf.Priority = pri
			win.img.lazy.VCS.TIA.Video.UpdateCTRLPF()
		})
	}

	imgui.SameLine()
	imguiText("CTRLPF")
	imgui.SameLine()
	imgui.PushItemWidth(win.byteDim.X)
	ctrlpf := fmt.Sprintf("%02x", lz.Ctrlpf)
	if imguiHexInput("##ctrlpf", !win.img.paused, 2, &ctrlpf) {
		if v, err := strconv.ParseUint(ctrlpf, 16, 8); err == nil {
			win.img.lazy.Dbg.PushRawEvent(func() {
				pf.SetCTRLPF(uint8(v))

				// update ball copy of CTRLPF too
				bl.SetCTRLPF(uint8(v))
			})
		}
	}
	imgui.PopItemWidth()
	imgui.EndGroup()

	imgui.Spacing()
	imgui.Spacing()

	// playfield data
	imgui.BeginGroup()
	imguiText("PF0")
	imgui.SameLine()
	seq := newDrawlistSequence(win.img, imgui.Vec2{X: imgui.FrameHeight(), Y: imgui.FrameHeight()}, 0.1)
	pf0d := lz.PF0
	for i := 0; i < 4; i++ {
		var col uint8
		if (pf0d<<i)&0x80 == 0x80 {
			col = lz.ForegroundColor
		} else {
			col = lz.BackgroundColor
			seq.nextItemDepressed = true
		}
		if seq.rectFilled(col) {
			pf0d ^= 0x80 >> i
			win.img.lazy.Dbg.PushRawEvent(func() { pf.SetPF0(pf0d) })
		}
		seq.sameLine()
	}
	seq.end()

	imgui.SameLine()
	imguiText("PF1")
	imgui.SameLine()
	seq.start()
	pf1d := lz.PF1
	for i := 0; i < 8; i++ {
		var col uint8
		if (pf1d<<i)&0x80 == 0x80 {
			col = lz.ForegroundColor
		} else {
			col = lz.BackgroundColor
			seq.nextItemDepressed = true
		}
		if seq.rectFilled(col) {
			pf1d ^= 0x80 >> i
			win.img.lazy.Dbg.PushRawEvent(func() { pf.SetPF1(pf1d) })
		}
		seq.sameLine()
	}
	seq.end()

	imgui.SameLine()
	imguiText("PF2")
	imgui.SameLine()
	seq.start()
	pf2d := lz.PF2
	for i := 0; i < 8; i++ {
		var col uint8
		if (pf2d<<i)&0x80 == 0x80 {
			col = lz.ForegroundColor
		} else {
			col = lz.BackgroundColor
			seq.nextItemDepressed = true
		}
		if seq.rectFilled(col) {
			pf2d ^= 0x80 >> i
			win.img.lazy.Dbg.PushRawEvent(func() { pf.SetPF2(pf2d) })
		}
		seq.sameLine()
	}
	seq.end()
	imgui.EndGroup()

	imgui.Spacing()
	imgui.Spacing()

	// playfield data as a sequence
	imgui.BeginGroup()
	imguiText("Sequence")
	seq = newDrawlistSequence(win.img, imgui.Vec2{X: imgui.FrameHeight() * 0.5, Y: imgui.FrameHeight()}, 0.1)

	// first half of the playfield
	for _, v := range lz.Data {
		var col uint8
		if v {
			col = lz.ForegroundColor
		} else {
			col = lz.BackgroundColor
		}
		seq.rectFilled(col)
		seq.sameLine()
	}

	// second half of the playfield
	for i, v := range lz.Data {
		// correct for reflected playfield
		if lz.Reflected {
			v = lz.Data[len(lz.Data)-1-i]
		}

		var col uint8
		if v {
			col = lz.ForegroundColor
		} else {
			col = lz.BackgroundColor
		}
		seq.rectFilled(col)
		seq.sameLine()
	}
	seq.end()

	// playfield index pointer
	if lz.Region != video.RegionOffScreen {
		idx := lz.Idx
		if lz.Region == video.RegionRight {
			idx += len(lz.Data)
		}

		p1 := imgui.Vec2{
			X: seq.offsetX(idx),
			Y: imgui.CursorScreenPos().Y,
		}

		dl := imgui.WindowDrawList()
		dl.AddCircleFilled(p1, imgui.FontSize()*0.20, win.idxPointer)
	}
	imgui.EndGroup()
}
