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
	"strings"

	"github.com/jetsetilly/gopher2600/hardware/tia/video"

	"github.com/inkyblackness/imgui-go/v2"
)

func (win *winTIA) drawMissile(missile int) {
	lz := win.img.lazy.Missile0
	ps := win.img.lazy.VCS.TIA.Video.Player0
	ms := win.img.lazy.VCS.TIA.Video.Missile0
	if missile != 0 {
		lz = win.img.lazy.Missile1
		ps = win.img.lazy.VCS.TIA.Video.Player1
		ms = win.img.lazy.VCS.TIA.Video.Missile1
	}

	imgui.Spacing()

	imgui.BeginGroup()
	imguiText("Colour")
	col := lz.Color
	if win.img.imguiSwatch(col) {
		win.popupPalette.request(&col, func() {
			win.img.lazy.Dbg.PushRawEvent(func() { ms.Color = col })

			// update player color too
			win.img.lazy.Dbg.PushRawEvent(func() { ps.Color = col })
		})
	}

	imguiText("Reset-to-Player")
	r2p := lz.ResetToPlayer
	if imgui.Checkbox("##resettoplayer", &r2p) {
		win.img.lazy.Dbg.PushRawEvent(func() { ms.ResetToPlayer = r2p })
	}

	imgui.SameLine()
	imguiText("Enabled")
	enb := lz.Enabled
	if imgui.Checkbox("##enabled", &enb) {
		win.img.lazy.Dbg.PushRawEvent(func() { ms.Enabled = enb })
	}
	imgui.EndGroup()

	imgui.Spacing()
	imgui.Spacing()

	// hmove value and slider
	imgui.BeginGroup()
	imguiText("HMOVE")
	imgui.SameLine()
	imgui.PushItemWidth(win.byteDim.X)
	hmove := fmt.Sprintf("%01x", lz.Hmove)
	if imguiHexInput("##hmove", !win.img.paused, 1, &hmove) {
		if v, err := strconv.ParseUint(hmove, 16, 8); err == nil {
			win.img.lazy.Dbg.PushRawEvent(func() { ps.Hmove = uint8(v) })
		}
	}
	imgui.PopItemWidth()

	imgui.SameLine()
	imgui.PushItemWidth(win.hmoveSliderWidth)
	hmoveSlider := int32(lz.Hmove) - 8
	if imgui.SliderIntV("##hmoveslider", &hmoveSlider, -8, 7, "%d") {
		win.img.lazy.Dbg.PushRawEvent(func() { ps.Hmove = uint8(hmoveSlider + 8) })
	}
	imgui.PopItemWidth()
	imgui.EndGroup()

	imgui.Spacing()
	imgui.Spacing()

	// nusiz
	imgui.BeginGroup()
	imgui.PushItemWidth(win.missileCopiesComboDim.X)
	if imgui.BeginComboV("##missilecopies", video.MissileCopies[lz.Copies], imgui.ComboFlagNoArrowButton) {
		for k := range video.MissileCopies {
			if imgui.Selectable(video.MissileCopies[k]) {
				v := uint8(k) // being careful about scope
				win.img.lazy.Dbg.PushRawEvent(func() {
					ms.Copies = v
					win.img.lazy.VCS.TIA.Video.UpdateNUSIZ(missile, true)
				})
			}
		}

		imgui.EndCombo()
	}
	imgui.PopItemWidth()

	imgui.SameLine()
	imgui.PushItemWidth(win.missileSizeComboDim.X)
	if imgui.BeginComboV("##missilesize", video.MissileSizes[lz.Size], imgui.ComboFlagNoArrowButton) {
		for k := range video.MissileSizes {
			if imgui.Selectable(video.MissileSizes[k]) {
				v := uint8(k) // being careful about scope
				win.img.lazy.Dbg.PushRawEvent(func() {
					ms.Size = v
					win.img.lazy.VCS.TIA.Video.UpdateNUSIZ(missile, true)
				})
			}
		}

		imgui.EndCombo()
	}
	imgui.PopItemWidth()

	imgui.SameLine()
	imguiText("NUSIZ")
	imgui.PushItemWidth(win.byteDim.X)
	nusiz := fmt.Sprintf("%02x", lz.Nusiz)
	if imguiHexInput("##nusiz", !win.img.paused, 2, &nusiz) {
		if v, err := strconv.ParseUint(nusiz, 16, 8); err == nil {
			win.img.lazy.Dbg.PushRawEvent(func() {
				ms.SetNUSIZ(uint8(v))

				// update player NUSIZ too
				ps.SetNUSIZ(uint8(v))
			})
		}
	}
	imgui.PopItemWidth()

	s := strings.Builder{}
	if lz.EncActive {
		s.WriteString("drawing ")
		if lz.EncSecondHalf {
			s.WriteString("2nd half of ")
		}
		switch lz.EncCpy {
		case 1:
			s.WriteString("1st copy")
		case 2:
			s.WriteString("2nd copy")
		}
	}
	imgui.SameLine()
	imgui.Text(s.String())
	imgui.EndGroup()

	imgui.Spacing()
	imgui.Spacing()

	// horizontal positioning
	imgui.BeginGroup()
	imgui.Text(fmt.Sprintf("Last reset at pixel %03d. Draws at pixel %03d", lz.ResetPixel, lz.HmovedPixel))
	if lz.MoreHmove {
		imgui.SameLine()
		imgui.Text("[currently moving]")
	}
	imgui.EndGroup()
}
