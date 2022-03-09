#!/usr/bin/env sh

# run confd to render out the config
#confd -onetime -backend env -log-level debug

sleep 2

# run
/app/${APP_NAME}
