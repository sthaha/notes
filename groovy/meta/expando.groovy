e = new Expando()
e.name = 'foobar'
e.speak = { "$name says woof" }

println e
println e.speak()

// so only changes the instance and not all instances
e2 = new Expando()
println e2



