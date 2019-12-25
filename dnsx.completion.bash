#!/bin/bash


## http://kodango.com/bash-competion-programming
# 例如我们定义这样一个补全函数_foo：
# $ function _foo()
# > {
# >     echo -e "\n"
# > 
# >     declare -p COMP_WORDS
# >     declare -p COMP_CWORD
# >     declare -p COMP_LINE
# >     declare -p COMP_WORDBREAKS
# > }
# $ complete -F _foo foo
# 假设我们在命令行下输入以下内容，再按下Tab键补全：

# $ foo b
# declare -a COMP_WORDS='([0]="foo" [1]="b")'
# declare -- COMP_CWORD="1"
# declare -- COMP_LINE="foo b"
# declare -- COMP_WORDBREAKS=" 	
# \"'><=;|&(:"


__dnsx_add_completions()
{
    # domainOpts="qq.com rockontrol querycap rktl"
    domainOpts="${*}"

    # domainOpts=$(dnsx configure domains ${COMP_LINE})
    typeOpts="A CNAME"

    local cur prev
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    case ${prev} in 
    add)  COMPREPLY=( $(compgen -W "${typeOpts}" -- ${cur}) ) ;;
    *)  COMPREPLY=( $(compgen -W "${domainOpts}" -- ${cur}) ) ;;
    esac

}

__dnsx__delete_completions()
{
    local
}

__dnsx_subcommand()
{

    local prev
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    case ${prev} {

    }
}


_dnsx_completions()
{
    local cur prev
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    # params
    case ${cur} in
    -*)
        COMPREPLY=( $(compgen -W "-p" -- ${cur}) ) 
        return 0 
        ;; 
    esac 
    case ${prev} in 
    "-p") 
        COMPREPLY=( $(compgen -W "$(dnsx configure list)" -- "${cur}") ) 
        return 0 
        ;; 
    esac

    # commands
    if [[ "${COMP_WORDS[1]}" == "-p" ]]; then
    {
        ## dnsx -p profile add ...
        [ -z "${COMP_WORDS[3]}" ] && {
            command=${COMP_WORDS[0]}
        } || {
            command=${COMP_WORDS[3]}
        }
    }
    elif [[ "${COMP_WORDS[2]}" == "-p" ]]; then
    {
        ## dnsx add -p profile ...
        command=${COMP_WORDS[1]}
    }
    else
    {
        ## dnsx add ...
        [ -n "${COMP_WORDS[1]}" ] && command=${COMP_WORDS[1]} || command=${COMP_WORDS[0]}
    }
    fi

    # main command
    subcmdOpts="add delete search switch configure help"
    domainOpts=$(dnsx configure domains ${COMP_LINE})
    case ${command} in
    "dnsx")
        COMPREPLY=( $(compgen -W "${subcmdOpts}" -- ${cur}) ) 
        return 0 
        ;; 
    "add") __dnsx_add_completions ${domainOpts} ;;
    esac 
}

complete -F _dnsx_completions dnsx
