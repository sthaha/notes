class Dog{}

println Dog
println Dog.metaClass // it is HandleMetaClass here
println Dog.metaClass.class.name


println "adding ... to metaclass"
Dog.metaClass.name = 'fido'
Dog.metaClass.speak =  { "$name says wooooof!" }

println Dog
println Dog.metaClass // now this becomes an ExpandoMetaClass[class Dog]


println new Dog().speak() // works fine!
