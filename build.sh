#! /bin/bash

libraries="curses"
binaries=""

# Parse Parameters
usage() {
    cat << EOF
usage: $0 [OPTIONS]

Builds the gozer libraries and binaries

OPTIONS:
   -h      Show this message
   -u      Update the dependences
   -f      Format the sources
   -t      Run the tests
   -l      Rebuild the libraries
   -b      Rebuild the binaries
   -j      Rebuild the json encoders
EOF
}

doUpdate=0
doFormat=0
doTests=0
buildLibraries=0
buildBinaries=0
generateJsonEncoders=0
while getopts “huftlb” OPTION
do
    case $OPTION in
        h)
            usage
            exit 1
            ;;
        u)
            doUpdate=1
            ;;
        f)
            doFormat=1
            ;;
        t)
            doTests=1
            ;;
        l)
            buildLibraries=1
            ;;
        b)
            buildBinaries=1
            ;;
        ?)
            usage
            exit
            ;;
    esac
done

buildBanner() {
    echo
    echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
    echo -e " " $1 ":" $2
    echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
}

buildStepStart() {
    echo -e "\t" $1 ":" $2
    echo " - - output - - - - - - - - - - - - - -"
}

buildStepEnd() {
    echo " - - result - - - - - - - - - - - - - -"
    echo -e "\t" $1 ":" $2
    echo "========================================"
}

buildBanner "Preparing Build"

buildStepStart "Configuring GOPATH"
if [[ -z $GVP_NAME ]]; then

    # GVP not run, therefore need to run it
    gvp init
    . gvp "in"
    export GOPATH="$GOPATH$(pwd):"

fi
buildStepEnd "GOPATH" $GOPATH

buildStepStart "Updating Dependencies"
if [ 1 -eq $doUpdate ]; then
    if !(gpm); then
        buildStepEnd "update" "FAILED"
        exit 1
    else
        buildStepEnd "update" "SUCCESS"
    fi
else
    buildStepEnd "update" "SKIPPED"
fi

formatCode() {
    buildStepStart "Formatting Code" $1
    if [ 1 -eq $doFormat ]; then
        if (go fmt "$1"); then
            buildStepEnd "formatting" "SUCCESS"
        else
            buildStepEnd "formatting" "FAILED"
        fi
    else
        buildStepEnd "formatting" "SKIPPED"
    fi
}

buildLibrary() {
    buildStepStart "Building Library" $1
    if (go build $1); then
        buildStepEnd "build" "SUCCESS"
    else
        buildStepEnd "build" "FAILED"
        exit 1
    fi
}

testLibrary() {
    buildStepStart "Testing Library" $1
    if [ 1 -eq $doTests ]; then
        if (go test $1); then
            buildStepEnd "tests" "SUCCESS"
        else
            buildStepEnd "tests" "FAILED"
            exit 1
        fi
    else
        buildStepEnd "tests" "SKIPPED"
    fi
}

installLibrary() {
    buildStepStart "Installing Library" $1
    if (go install $1); then
        buildStepEnd "install" "SUCCESS"
    else
        buildStepEnd "install" "FAILED"
        exit 1
    fi
}

buildBinary() {
    buildStepStart "Building Binary" $1
    binaryName="${1##*/}"
    if (go build -o "bin/$binaryName" $1); then
        buildStepEnd "build" "SUCCESS"
    else
        buildStepEnd "build" "FAILED"
        exit 1
    fi
}

if [ 1 -eq $buildLibraries ]; then
    for library in $libraries; do
        buildBanner "Building" $library

        formatCode $library

        buildLibrary $library

        testLibrary $library

        installLibrary $library

        echo
    done
fi

if [ 1 -eq $buildBinaries ]; then
    for binary in $binaries; do
        buildBanner "Building" $binary

        buildBinary $binary

        echo
    done
fi


