require "./spec_helper"

describe Groups do

  it "can find 0 block" do
    m = [
      [[0],
       [0]],

      [[0, 0],
       [0, 0]],

      [[0, 0],
       [0, 0],
       [0, 0]],

      [[0, 0, 0],
       [0, 0, 0],
       [0, 0, 0]],
    ]
    m.each { |mx|  Groups.find(mx).should eq(0) }


  end

  it "can find whole block" do
    m = [
      [[1]],

      [[0, 0, 0],
       [0, 1, 0],
       [0, 0, 0]],

      [[0, 0, 0],
       [0, 0, 0],
       [0, 0, 1]],

      [[1, 1, 1],
       [1, 1, 1],
       [1, 1, 1]],

      [[1, 1, 1],
       [1, 0, 1],
       [1, 0, 1]],

      [[1, 0, 1],
       [1, 0, 1],
       [1, 1, 1]],

      [[1, 0, 1],
       [0, 1, 0],
       [1, 0, 1]]
    ]

    m.each { |mx|  Groups.find(mx).should eq(1) }
  end

  it "can find 2 groups" do
    m = [
      [[1, 0, 1],
       [1, 0, 1],
       [1, 0, 1]],

      [[1, 1, 1],
       [0, 0, 0],
       [1, 1, 1]],

      [[1, 1, 0],
       [0, 0, 0],
       [0, 1, 1]],

      [[1, 0, 1],
       [1, 0, 0],
       [1, 0, 0]],

      [[1, 1, 1],
       [1, 0, 0],
       [1, 0, 1]],

      [[1, 0, 1],
       [1, 0, 0],
       [1, 1, 1]]
    ]

    m.each { |mx|  Groups.find(mx).should eq(2) }
  end

  it "can find 3 groups" do
    m = [
      [[1, 0, 0],
       [0, 0, 1],
       [1, 0, 0]],

      [[1, 0, 1],
       [1, 0, 0],
       [1, 0, 1]],

      [[1, 0, 1],
       [0, 0, 1],
       [1, 0, 1]]
    ]

    m.each { |mx|  Groups.find(mx).should eq(3) }
  end

  it "can find 4x4 groups" do
    m = [
      [[1, 0, 0, 0],
       [0, 1, 0, 1],
       [0, 0, 1, 1],
       [0, 0, 0, 0]],

    ]

    m.each { |mx|  Groups.find(mx).should eq(1) }
  end
end
