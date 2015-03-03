#!/bin/bash

NAME="neurogo"
APP_DIR="/srv/sites/neurogo"
SOCKFILE=/srv/sites/neurogo/run/server.sock

echo "Starting $NAME"

# Create the run directory if it doesn't exist
RUNDIR=$(dirname $SOCKFILE)
test -d $RUNDIR || mkdir -p $RUNDIR

go build $APP_DIR/application.go

source $APP_DIR/application