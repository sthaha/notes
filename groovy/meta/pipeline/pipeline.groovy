
class d {
 static listeners = [:]
}


static def on(String event, Closure c) {
  on([event], c)
}

static def on(List events, Closure c) {
  events.each { e ->
    d.listeners[e] = d.listeners[e] ?: [] as Set
    d.listeners[e] << c
    println "... registered for $e ${d.listeners[e]}"
  }
}

static def emit(List events, Object... args) {
  events.each { e ->
    if (!d.listeners[e]) {
      return
    }
    d.listeners[e].each { c -> c.call([name: e], args) }
  }
}


static def emit(String event, Object... args) {
  emit([event], args)
}
