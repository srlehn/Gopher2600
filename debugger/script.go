package debugger

import (
	"fmt"
	"gopher2600/debugger/ui"
	"gopher2600/errors"
	"io/ioutil"
	"os"
	"strings"
)

func (dbg *Debugger) loadScript(scriptfile string) ([]string, error) {

	// open script and defer closing
	sf, err := os.Open(scriptfile)
	if err != nil {
		return nil, errors.NewFormattedError(errors.ScriptFileCannotOpen, err)
	}
	defer func() {
		_ = sf.Close()
	}()

	buffer, err := ioutil.ReadAll(sf)
	if err != nil {
		return nil, errors.NewFormattedError(errors.ScriptFileError, err)
	}

	// convert buffer to an array of lines
	s := fmt.Sprintf("%s", buffer)
	return strings.Split(s, "\n"), nil
}

// RunScript uses a text file as a source for a sequence of commands
func (dbg *Debugger) RunScript(scriptfile string, silent bool) error {

	// the silent flag passed to this function is meant to silence commands for
	// the duration of the script only. store existing state of dbg.silent so we
	// can restore it when script has concluded
	uiSilentRestore := dbg.uiSilent
	dbg.uiSilent = silent
	defer func() {
		dbg.uiSilent = uiSilentRestore
	}()

	// load file
	lines, err := dbg.loadScript(scriptfile)
	if err != nil {
		return err
	}

	// parse each line as user input
	for i := 0; i < len(lines); i++ {
		if strings.Trim(lines[i], " ") != "" {
			if !silent {
				dbg.print(ui.Script, lines[i])
			}
			next, err := dbg.parseInput(lines[i])
			if err != nil {
				return errors.NewFormattedError(errors.ScriptFileError, err)
			}
			if next {
				// make sure run state is still sane
				dbg.runUntilHalt = false

				return errors.NewFormattedError(errors.ScriptRunError, strings.ToUpper(lines[i]), scriptfile, i)
			}
		}
	}

	return nil
}
