#!/bin/bash
#
# This script used to build the app based on cmd/<app name>

showHelp() {
    echo ""
    echo "App Builder"
    echo ""
    echo -e "Usage: bash build.sh [Options]"
    echo ""
    echo -e "Options:"
    echo -e "   --app <app name>   the <app name> is based on directory name inside cmd/ directory"
    echo -e "   --out <save path>  path/filename to save the bin file"
    echo ""
    echo -e "Example: bash build.sh --app rintik-auth --out /apps"
    echo -e "This will create a binary file named rintik-auth in the /apps directory"

    return 1
}

while :; do
    case $1 in
        -h|-\?|--help)
            showHelp    # Display a usage synopsis.
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
        -out|--out)       # Takes an option argument; ensure it has been specified.
            if [ "$2" ]; then
                out=$2
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

doBuild() {    
    if [ -z $app ]; then
        echo "ERROR: args --app must not empty"
        exit 1
    fi

    echo "Building bin file $app"
    go build ./cmd/$app

    if [ ! -z $out ]; then        
        if [ ! -n "$out" ]; then
            echo "ERROR: out $out not found"
            exit 1
        fi
        echo "Moving $app file to $out"
        mv ./$app $out
    fi
}

doBuild