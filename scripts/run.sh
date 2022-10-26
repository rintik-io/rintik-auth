#!/bin/bash
#
# This script used to run the app based on cmd/<app name>

listCMD=(`ls ./cmd`)
config="./.configs.yml"
log="./log"

declare -A listApp

for cmd in "${listCMD[@]}"
do
   listApp[$cmd]=1
   # or do whatever with individual element of the array
done

showHelp() {
    echo ""
    echo "App Runner"
    echo ""
    echo -e "Usage: bash run.sh [Options]"
    echo ""
    echo -e "Options:"
    echo -e "   --app <app name>  the <app name> is based on directory name inside cmd/ directory"
    echo -e "   --list            show list of available app inside cmd/ directory"
    echo -e "   --config          set config file"
    echo -e "   --log             set log path"
    echo ""
    echo -e "Example: bash run.sh --app rintik-auth --config /home/app/configs/.configs.yaml --log /home/app/log"

    return 1
}

showList() {
    echo -e "Available App:"
    printf ' %s\n' "${listCMD[@]}" 

    return 1
}

while :; do
    case $1 in
        -h|-\?|--help)
            showHelp    # Display a usage synopsis.
            exit 1
            ;;
        -l|-list|--list)
            showList    # Display a usage synopsis.
            exit 1
            ;;
        -app|--app)       # Takes an option argument; ensure it has been specified.
            if [ "$2" ]; then
                app=$2
                shift
            else
                echo 'ERROR: "--app" requires a non-empty option argument.'
                exit 1
            fi
            ;;
        -config|--config)       # Takes an option argument; ensure it has been specified.
            if [ "$2" ]; then
                config=$2
                shift
            fi
            ;;
        -log|--log)       # Takes an option argument; ensure it has been specified.
            if [ "$2" ]; then
                log=$2
                shift
            fi
            ;;
        --)              # End of all options.
            shift
            break
            ;;
        -?*)
            showHelp
            printf 'ERROR: Unknown option : %s\n' "$1" >&2
            exit 1
            ;;
        *)               # Default case: No more options, so break out of the loop.
            break
    esac

    shift
done

doRun(){
    appName=("")
    if [ -z $app ]; then
        if [[ $APP_RUN != "" ]]; then
            if [ ! -n "${listApp[$APP_RUN]}" ]; then
                echo "ERROR: app name [ $APP_RUN ] not found in cmd/"
                exit 1
            fi
            appName=$APP_RUN
        else
            echo "ERROR: Define which app to run with --app flags or set env variable APP_RUN. See --help."
            exit 1
        fi
    else
        appName=($app)
    fi

    if [[ $APP_CONFIG != "" ]]; then
        config=$APP_CONFIG
    else
        config=$config
    fi

    if [[ $APP_LOG != "" ]]; then
        log=($APP_LOG/$appName.log)
    else
        log=($log/$appName.log)
    fi

    go run cmd/$appName/main.go -config=$config -log=$log
}

doRun