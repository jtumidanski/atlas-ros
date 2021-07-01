package statistics

import (
	"atlas-ros/wz"
	"atlas-ros/xml"
	"errors"
	"fmt"
	"strconv"
)

func readStatistics(id uint32) (*Model, error) {
	w, err := findReactor(id)
	if err != nil {
		return nil, err
	}
	rd, err := xml.Read(w.Path())
	if err != nil {
		return nil, err
	}

	// read linked reactor
	info, err := rd.ChildByName("info")

	if info == nil {
		return NewModel().AddState(0, []ReactorState{{theType: 999, reactorItem: nil, activeSkills: nil, nextState: 0}}, -1), nil
	}

	link := info.GetString("link", "")
	if link != "" {
		lid, err := strconv.Atoi(link)
		if err != nil {
			return nil, err
		}
		return readStatistics(uint32(lid))
	}

	loadArea := info.GetIntegerWithDefault("activateByTouch", 0) != 0

	rid, err := rd.ChildByName("0")
	i := byte(0)
	m := NewModel()
	for rid != nil {
		areaSet := false
		sdl := make([]ReactorState, 0)
		ed, _ := rid.ChildByName("event")
		if ed != nil {
			timeout := ed.GetIntegerWithDefault("timeOut", -1)

			for _, md := range ed.ChildNodes {
				t := uint32(md.GetIntegerWithDefault("type", 0))
				var ri *ReactorItem
				if t == 100 {
					itemId := uint32(md.GetIntegerWithDefault("0", 0))
					quantity := uint16(md.GetIntegerWithDefault("1", 0))
					ri = &ReactorItem{itemId: itemId, quantity: quantity}
					if !areaSet || loadArea {
						m = m.SetTL(md.GetPoint("tl", 0, 0))
						m = m.SetRB(md.GetPoint("rb", 0, 0))
						areaSet = true
					}
				}
				skillIds := make([]uint32, 0)
				activeSkillId, _ := md.ChildByName("activeSkillID")
				if activeSkillId != nil {
					for _, s := range activeSkillId.ChildNodes {
						skillIds = append(skillIds, uint32(md.GetIntegerWithDefault(s.Name, 0)))
					}
				}
				ns := byte(md.GetIntegerWithDefault("state", 0))
				sdl = append(sdl, ReactorState{theType: t, reactorItem: ri, activeSkills: skillIds, nextState: ns})
			}

			m = m.AddState(i, sdl, timeout)
		}
		i++
		rid, _ = rd.ChildByName(string(i))
	}
	return m, nil
}

func findReactor(id uint32) (*wz.FileEntry, error) {
	path := fmt.Sprintf("%07d", id)
	if val, ok := wz.GetFileCache().GetFile(path + ".img.xml"); ok == nil {
		return val, nil
	}
	return nil, errors.New(fmt.Sprintf("reactor %d not found", id))
}
