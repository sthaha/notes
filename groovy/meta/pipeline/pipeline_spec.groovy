import pipeline
import sh

shx = new sh()

shx("starting spec ...")


pipeline.on("build") { e, Object... args ->
  println ">> Got event $e: ${args}"
  shx( ">> ls /home/jenkins - $e" )
}

pipeline.on(["build.start", "build.end"]) { e, Object... args ->

  // sh "Hello"
  println ">> Got event $e: ${args}"
  shx( ">> ls /tmp - $e.name" )
}


println "Emiting build ..."
pipeline.emit(
  ["build", "build.start"],
  [repo: "https://github.com/foobar", sha: "deadbeef"])
pipeline.emit("build.foobar")

pipeline.emit("build.pass")
pipeline.emit("build.fail")

println "Emiting build.end ..."
pipeline.emit("build.end")
