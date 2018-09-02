package foobar

// to run: groovy log_extend.groovy
// somehow inclues log dir

import java.util.logging.Logger
import java.util.logging.Level
import log.*


Logger.metaClass.methodMissing = {String name, args ->
  println "Creating an new method: $name"

    def warning = Level.WARNING.intValue()
    def severe = Level.WARNING.intValue()

    int val = warning + ( severe - warning) * Math.random()
    def level = new log.CustomLevel(name.toUpperCase(), val)
    def impl = {Object... vargs ->
    println "calling $name: $vargs"
    delegate.log(level, vargs[0])
    }
    Logger.metaClass."$name" = impl
    impl args
}


def log = Logger.getLogger(this.class.name)

log.wtf "what is this dude?", "see you"
log.wtf "what is this dude - again?"

log.ok "alls well is this dude?"
log.ok "See you all then"
