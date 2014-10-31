package main

import (
	"os"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/bgentry/go-netrc/netrc"
	"github.com/heroku/hk/apps"
	"github.com/heroku/hk/cli"
	"github.com/heroku/hk/plugins"
)

var Cli = cli.NewCli(
	apps.Apps,
	apps.Info,
	plugins.Plugins,
)

func main() {
	defer handlePanic()
	plugins.Setup()
	for _, topic := range plugins.PluginTopics() {
		Cli.AddTopic(topic)
	}
	ctx, err := Cli.Parse(os.Args[1:])
	if err != nil {
		if err == cli.HelpErr {
			help()
		}
		cli.Errln(err)
		cli.Errf("USAGE: %s %s\n", os.Args[0], commandSignature(ctx.Topic, ctx.Command))
		os.Exit(2)
	}
	if ctx.Command.NeedsApp {
		if ctx.App == "" {
			ctx.App = app()
		}
		if app := os.Getenv("HEROKU_APP"); app != "" {
			ctx.App = app
		}
		if ctx.App == "" {
			AppNeededWarning()
		}
	}
	if ctx.Command.NeedsAuth {
		ctx.Auth.Username, ctx.Auth.Password = auth()
	}
	cli.Logf("Running %s\n", ctx)
	before := time.Now()
	ctx.Command.Run(ctx)
	cli.Logf("Finished in %s\n", (time.Since(before)))
}

func handlePanic() {
	if e := recover(); e != nil {
		cli.Errln("ERROR:", e)
		cli.Logln(debug.Stack())
		cli.Exit(1)
	}
}

func app() string {
	app, err := appFromGitRemote(remoteFromGitConfig())
	if err != nil {
		panic(err)
	}
	return app
}

func auth() (user, password string) {
	netrc, err := netrc.ParseFile(filepath.Join(cli.HomeDir, ".netrc"))
	if err != nil {
		panic(err)
	}
	auth := netrc.FindMachine("api.heroku.com")
	return auth.Login, auth.Password
}
