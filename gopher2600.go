package main

import (
	"flag"
	"fmt"
	"gopher2600/debugger"
	"gopher2600/debugger/colorterm"
	"gopher2600/hardware"
	"gopher2600/television"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

func main() {
	mode := flag.String("mode", "DEBUG", "emulation mode: DEBUG, FPS, DISASM")
	flag.Parse()

	switch strings.ToUpper(*mode) {
	case "DEBUG":
		// create a new debugger with the choice of terminal
		// TODO: implement flag for plain terminal
		dbg, err := debugger.NewDebugger(new(colorterm.ColorTerminal))
		if err != nil {
			fmt.Printf("* error starting debugger (%s)\n", err)
			os.Exit(10)
		}

		// run initialisation script
		err = dbg.RunScript(".gopher2600/debuggerInit", true)
		if err != nil {
			fmt.Println(err)
			os.Exit(10)
		}

		// start debugger with choice of cartridge
		// TODO: implement command line selection of cartridge
		err = dbg.Start("roms/ball_test_card.bin")
		if err != nil {
			fmt.Println(err)
			os.Exit(10)
		}
	case "FPS":
		err := fps()
		if err != nil {
			fmt.Println(err)
			os.Exit(10)
		}
	case "DISASM":
		fmt.Printf("* not yet implemented")
		os.Exit(10)

	default:
		fmt.Printf("* unknown mode (%s)\n", strings.ToUpper(*mode))
		os.Exit(10)
	}

}

func fps() error {
	tv := new(television.DummyTV)
	if tv == nil {
		return fmt.Errorf("error creating television for fps profiler")
	}

	vcs, err := hardware.New(tv)
	if err != nil {
		return fmt.Errorf("error starting fps profiler (%s)", err)
	}

	err = vcs.AttachCartridge("roms/flappy.bin")
	if err != nil {
		return err
	}

	const cyclesPerFrame = 19912
	const numOfFrames = 180

	f, err := os.Create("cpu.profile")
	if err != nil {
		return err
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		return err
	}
	defer pprof.StopCPUProfile()

	cycles := cyclesPerFrame * numOfFrames
	startTime := time.Now()
	for cycles > 0 {
		stepCycles, _, err := vcs.Step(hardware.NullVideoCycleCallback)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("%d cycles completed\n", cycles)
			return nil
		}
		cycles -= stepCycles
	}

	fmt.Printf("%f fps\n", float64(numOfFrames)/time.Since(startTime).Seconds())

	return nil
}
