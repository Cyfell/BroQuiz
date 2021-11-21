#!/bin/sh

if [[ ! -x "$(command -v inotifywait)" ]]; then
    echo >&2 "inotify-tools is required"
    exit 1
fi

function main {
    HOTRELOAD=true

    function stop_hotreload {
        HOTRELOAD=false
    }
    trap "stop_hotreload" SIGINT SIGTERM EXIT

    while ${HOTRELOAD}; do
        echo -e "[Testing]"
        go test -coverprofile=/dev/null $1/...
        inotifywait -e modify -e move -e create -e delete -r --include "\.(go)$" $1
    done
}

main $@