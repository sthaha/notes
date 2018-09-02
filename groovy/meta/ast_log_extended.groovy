package log

import java.util.logging.*

Logger.metaClass.methodMissing = {String name, args ->
  println "Creating an new method: $name"

    def warning = Level.WARNING.intValue()
    def severe = Level.WARNING.intValue()

    int val = warning + ( severe - warning) * Math.random()
    def level = new CustomLogLevel(name.toUpperCase(), val)
    def impl = { Object... args -> delegate.log(level, args[0]) }
    Logger.metaClass."$name" = impl
    impl args
}


