Watch
=====

Sometimes you find yourself on a system without `watch` (\**cough\** windows \**cough\**).

This is a quick and dirty replacement

    go get github.com/MindTwister/watch
    go install github.com/MindTwister/watch

You now have a quick and dirty watch replacement.

##Usage:

    watch -interval=4 <command>


###Arguments

    -interval=[int]

**Default:** 2

Defines how often the given command should be run, any further arguments will be passed to the command
