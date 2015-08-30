__apply_localenv () {
    for pair in $(localenv list); do
        eval "export $pair"
    done
}

__cd_with_localenv () {
    cd "$1"
    __apply_localenv
}

alias cd=__cd_with_localenv
__apply_localenv
