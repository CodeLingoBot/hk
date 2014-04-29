[hk][hk] is a new command line client to the Heroku platform, designed to be
fast and lightweight. While hk is intended to replace the
[Heroku Ruby CLI][ruby-cli], it's designed as a completely new tool rather than
a drop-in replacement.

## Installing hk

hk is distributed as an auto-updating executable. Once you've installed it, hk
will periodically perform a version check in the background. If a newer version
is found, hk will update itself automatically.

hk is currently only being distributed for Mac OS X and Linux users. More
user-friendly installers, including one for Windows, are planned.

To install a pre-built binary release, run the following in a command line
terminal:

```term
L=/usr/local/bin/hk && curl -sL -A "`uname -sp`" https://hk.heroku.com/hk.gz | zcat >$L && chmod +x $L
```

If you wish, you can customize the install location by modifying the `L=`
variable at the beginning of that command.

The URL [https://hk.heroku.com/hk.gz](https://hk.heroku.com/hk.gz) will attempt
to detect your OS and CPU architecture based on the User-Agent, then redirect
you to the latest release for your platform.

If you've installed hk on a machine that already had the Heroku Ruby CLI, you
can start using hk immediately. If it's a new machine or you've never logged in,
you'll need to do so by running `hk login`.

## Getting help with hk

hk has a simple help system. The most common commands are listed in the basic
help output, which is also available via `hk help`:

```term
$ hk help
Usage: hk <command> [-a app] [options] [arguments]


Commands:

    create          create an app
    apps            list apps
    dynos           list dynos
    releases        list releases
    release-info    show release info
    rollback        roll back to a previous release
    addons          list addons
    addon-add       add an addon
    addon-destroy   destroy an addon
    scale           change dyno quantities and sizes
    restart         restart dynos
    set             set env var
    unset           unset env var
    env             list env vars
    run             run a process in a dyno
    log             stream app log lines
    info            show app info
    rename          rename an app
    destroy         destroy an app
    domains         list domains
    domain-add      add a domain
    domain-remove   remove a domain
    version         show hk version

Run 'hk help [command]' for details.


Additional help topics:

    commands  list all commands with usage
    environ   environment variables used by hk
    plugins   interface to plugin commands
    more      additional commands, less frequently used
    about     information about hk (e.g. copyright, license, etc.)
```

Commands that are used less frequently are listed under `hk help more`. For any
specific command, you can run `hk help <command>` to get the detailed help and
usage info for that command.

## Important differences from the Heroku Ruby CLI

The Heroku Ruby CLI organized its commands under nested namespaces, separated
with colons (i.e. `domains:add`). Most, but not all, of these namespaces were
pluralized.

hk, however, uses a simple, flat command space (i.e. `domains` and
`domain-add`). Commands use pluralized nouns where it's logical to do so, such
as lists of items (`apps`, `dynos`, `addons`, `releases`). The rest of the
commands are named with singular nouns because they deal with a single resource
(`addon-add`, `domain-remove`, `release-info`).

The Heroku Ruby CLI allowed users to specify an app for commands via a `-a
<app>` flag, or via a git remote with `-r <remote>`. hk combines these flags
into a single `-a <app or remote>` flag.

## Guide to hk commands for Heroku users

Many commands are similar in both the Heroku Ruby CLI and hk. However, some
commands have different names and take different arguments. This is a list of
frequently used commands, showing how to accomplish the same thing with either
CLI.

### Apps

#### Create an app

```term
$ heroku create myapp
```

```term
$ hk create myapp
```

With a region specified:

```term
$ heroku create --region eu myapp
```

```term
$ hk create -r eu myapp
```

#### List apps

```term
$ heroku list
```

```term
$ hk apps
```

#### Show an app's info

```term
$ heroku info
```

```term
$ hk info
```

#### Destroy an app

```term
$ heroku destroy -a myapp
```

```term
$ hk destroy myapp
```

The `heroku destroy` command can infer your app name from the current
directory's git remotes. For safety, however, `hk destroy` always requires you
to specify the name of the app you want to destroy.

This command can permanently destroy data, so it prompts for confirmation.

#### Rename an app

```term
$ heroku rename -a oldappname newappname
```

```term
$ hk rename oldappname newappname
```

#### View your application log

```term
$ heroku logs --tail
```

```term
$ hk log
```

The `hk log` command follows your application log stream by default (which
required a `--tail` flag in the Toolbelt).

### Dynos

#### Change dyno scale

```term
$ heroku ps:scale web=2 worker=4:PX
```

```term
$ hk scale web=2 worker=4:PX
```

This command is mostly identical in hk, except that it doesn't support scaling
by relative increments (i.e. `web+2`).

#### List dynos

```term
$ heroku ps
```

```term
$ hk dynos
```

### App configuration (environment)

#### Set an environment variable on an app

```term
$ heroku config:set KEY=value
```

```term
$ hk set KEY=value
```

#### List app's environment settings

```term
$ heroku config
```

```term
$ hk env
```

#### Show a single environment variable

```term
$ heroku config:get KEY
```

```term
$ hk get KEY
```

#### Unset an environment variable on an app

```term
$ heroku config:unset KEY1 KEY2
```

```term
$ hk unset KEY1 KEY2
```

### Domain Names

#### List domain names

```term
$ heroku domains
```

```term
$ hk domains
```

#### Add a domain name

```term
$ heroku domains:add www.test.com
```

```term
$ hk domain-add www.test.com
```

#### Remove a domain name

```term
$ heroku domains:remove www.test.com
```

```term
$ hk domain-remove www.test.com
```

### Add-ons

#### List add-ons on an app

```term
$ heroku addons
```

```term
$ hk addons
```

#### Add an add-on

```term
$ heroku addons:add heroku-postgresql
```

```term
$ hk addon-add heroku-postgresql
```

With additional provisioning options:

```term
$ heroku addons:add heroku-postgresql --fork RED
```

```term
$ hk addon-add heroku-postgresql fork=red
```

Additional add-on config is provided as `key=value` pairs rather than
`--key value` flags.

#### Destroy an add-on

```term
$ heroku addons:remove redistogo
```

```term
$ hk addon-remove redistogo
```

Next, a Heroku Postgres addon:

```term
$ heroku addons:remove heroku-postgresql:dev
```

```term
$ hk addon-remove heroku-postgresql-blue
```

Add-ons in hk are referenced by their `name`. Usually this is just the addon
provider's name, but for Heroku Postgres, it's of the form:
`heroku-postgresql-color`. In either case, the name matches what's displayed in
`hk addons`.

This command can permanently destroy data, so it prompts for confirmation.

### Access Control (sharing with collaborators)

#### Add access for a user

```term
$ heroku sharing:add user@test.com
```

```term
$ hk access-add user@test.com
```

#### Remove access for a user

```term
$ heroku sharing:remove user@test.com
```

```term
$ hk access-remove user@test.com
```

### Releases

#### View an app's releases

```term
$ heroku releases
```

```term
$ hk releases
```

#### View release info

```term
$ heroku releases:info v123
```

```term
$ hk release-info v123
```

#### Rollback to a previous release

```term
$ heroku rollback
```

```term
$ hk rollback v122
```

The Toolbelt attempts to rollback by one version, but hk requires you to specify
the version you want to rollback to.

## Full Command List

```term
$ hk help commands
```

```
hk access [-a <app>]                                              # list access permissions (extra)
hk access-add [-a <app>] [-s] <email>                             # give a user access to an app (extra)
hk access-remove [-a <app>] <email>                               # remove a user's access to an app (extra)
hk account-feature-disable <feature>                              # disable an account feature (extra)
hk account-feature-enable <feature>                               # enable an account feature (extra)
hk account-feature-info <feature>                                 # show info for an account feature (extra)
hk account-features                                               # list account features (extra)
hk addon-add [-a <app>] <service>[:<plan>] [<config>=<value>...]  # add an addon
hk addon-destroy [-a <app>] <name>                                # destroy an addon
hk addon-open [-a <app>] <name>                                   # open an addon (extra)
hk addon-plan [-a <app>] <name> <plan>                            # change an addon's plan (extra)
hk addon-plans <service>                                          # list addon plans (extra)
hk addon-services                                                 # list addon services (extra)
hk addons [-a <app>] [<service>:<plan>...]                        # list addons
hk api <method> <path>                                            # make a single API request (extra)
hk apps [<name>...]                                               # list apps
hk create [-r <region>] [<name>]                                  # create an app
hk creds                                                          # show credentials (extra)
hk destroy <name>                                                 # destroy an app
hk domain-add [-a <app>] <domain>                                 # add a domain
hk domain-remove [-a <app>] <domain>                              # remove a domain
hk domains [-a <app>]                                             # list domains
hk drain-add [-a <app>] <url>                                     # add a log drain (extra)
hk drain-info [-a <app>] <id or url>                              # show info for a log drain (extra)
hk drain-remove [-a <app>] <id or url>                            # remove a log drain (extra)
hk drains [-a <app>]                                              # list log drains (extra)
hk dynos [-a <app>] [<name>...]                                   # list dynos
hk env [-a <app>]                                                 # list env vars
hk feature-disable [-a <app>] <feature>                           # disable an app feature (extra)
hk feature-enable [-a <app>] <feature>                            # enable an app feature (extra)
hk feature-info [-a <app>] <feature>                              # show info for an app feature (extra)
hk features [-a <app>]                                            # list app features (extra)
hk get [-a <app>] <name>                                          # get env var (extra)
hk help [<topic>]                                                 # 
hk info [-a <app>]                                                # show app info
hk key-add [<public-key-file>]                                    # add ssh public key (extra)
hk key-remove <fingerprint>                                       # remove an ssh public key (extra)
hk keys                                                           # list ssh public keys (extra)
hk log [-a <app>] [-n <lines>] [-s <source>] [-d <dyno>]          # stream app log lines
hk login                                                          # log in to your Heroku account (extra)
hk logout                                                         # log out of your Heroku account (extra)
hk maintenance [-a <app>]                                         # show app maintenance mode (extra)
hk maintenance-disable [-a <app>]                                 # disable maintenance mode (extra)
hk maintenance-enable [-a <app>]                                  # enable maintenance mode (extra)
hk open [-a <app>]                                                # open app in a web browser (extra)
hk pg-info [-a <app>] <dbname>                                    # show Heroku Postgres database info (extra)
hk pg-list [-a <app>]                                             # list Heroku Postgres databases (extra)
hk pg-unfollow [-a <app>] <dbname>                                # stop a replica postgres database from following (extra)
hk psql [-a <app>] [-c <command>] [<dbname>]                      # open a psql shell to a Heroku Postgres database (extra)
hk regions                                                        # list regions (extra)
hk release-info [-a <app>] <version>                              # show release info
hk releases [-a <app>] [-n <limit>] [<version>...]                # list releases
hk rename <oldname> <newname>                                     # rename an app
hk restart [-a <app>] [<type or name>]                            # restart dynos
hk rollback [-a <app>] <version>                                  # roll back to a previous release
hk run [-s <size>] [-d] <command> [<argument>...]                 # run a process in a dyno
hk scale [-a <app>] <type>=[<qty>]:[<size>]...                    # change dyno quantities and sizes
hk set [-a <app>] <name>=<value>...                               # set env var
hk ssl [-a <app>]                                                 # show ssl endpoint info
hk ssl-cert-add [-a <app>] <certfile> <keyfile>                   # add a new ssl cert
hk ssl-cert-rollback [-a <app>]                                   # add a new ssl cert
hk ssl-destroy [-a <app>]                                         # destroy ssl endpoint
hk status                                                         # display heroku platform status (extra)
hk transfer [-a <app>] <email>                                    # transfer app ownership to a collaborator (extra)
hk transfer-accept [-a <app>]                                     # accept an inbound app transfer (extra)
hk transfer-cancel [-a <app>]                                     # cancel an outbound app transfer (extra)
hk transfer-decline [-a <app>]                                    # decline an inbound app transfer (extra)
hk transfers [-a <app>]                                           # list existing app transfers (extra)
hk unset [-a <app>] <name>...                                     # unset env var
hk update                                                         # 
hk url [-a <app>]                                                 # show app url (extra)
hk version                                                        # show hk version
hk which-app [-a <app>]                                           # show which app is selected, if any (extra)
```

## Feedback

hk is a beta product in active development. We'd love to hear about any issues
you run into, or any feedback you have. Please email us at
[hk-feedback@heroku.com](mailto:hk-feedback@heroku.com)

[hk]: https://github.com/heroku/hk "hk on Github"
[ruby-cli]: https://github.com/heroku/heroku "Heroku Ruby CLI"
