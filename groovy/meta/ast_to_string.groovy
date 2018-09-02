
// class Person {
  // String first
  // String last
// }

// p = new Person(first: "John", last: "Doe")
// println p
//  > Person@7c469c48

// import groovy.transform.*

// @ToString
// class Person {
  // String first
  // String last
// }

// p = new Person(first: "John", last: "Doe")
// println p

// > Person(John, Doe)


import groovy.transform.*

@Canonical    // adds ToString as well
class Person {
  String first
  String last
}

p = new Person(first: "John", last: "Doe")
println p
