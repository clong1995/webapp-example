package main

import (
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/log/jotrotate"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/driver"
	"github.com/richardwilkes/webapp/stdmenu"
)

func main() {
	cmdline.AppName = "Example"
	cmdline.AppCmdName = "example"
	cmdline.AppVersion = "0.2"
	cmdline.CopyrightYears = "2018-2019"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.AppIdentifier = "com.trollworks.webapp.example"

	args, err := webapp.Initialize(driver.ForPlatform())
	jot.FatalIfErr(err)

	cl := cmdline.New(true)
	jotrotate.ParseAndSetup(cl)

	webapp.WillFinishStartupCallback = finishStartup

	// Start only returns on error
	jot.FatalIfErr(webapp.Start(args, nil, nil))
	atexit.Exit(0)
}

func finishStartup() {
	wnd, err := webapp.NewWindow(webapp.StdWindowMask, webapp.MainDisplay().UsableBounds, "Example", "https://youtube.com")
	jot.FatalIfErr(err)
	if bar, global, first := webapp.MenuBarForWindow(wnd); !global || first {
		stdmenu.FillMenuBar(bar, nil, nil, true)
	}
	wnd.ToFront()
}
