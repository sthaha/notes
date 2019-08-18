require "./expr_helper"

describe Expr::Number do
  it "prints fine" do
    x = Expr::Number.new(10)
    "#{x}".should eq("10")
  end
end


describe Expr::Add do
  it "prints fine" do
    x = Expr::Add.new(
      Expr::Number.new(11),
      Expr::Number.new(31),
    )
    "#{x}".should eq("11 + 31")
  end

  it "can be reduced" do
    x = Expr::Add.new(
      Expr::Number.new(11),
      Expr::Number.new(31),
    )

    x.reduce(EMPTY_ENV).as(Expr::Number).value.should eq(42)
  end

  it "can be reduced multiple times" do
    x = Expr::Add.new(
      Expr::Number.new(10),
      Expr::Number.new(11),
    )
    x.reducible?.should eq(true)

    expr = Expr::Add.new(x, x)

    "#{expr}".should eq("10 + 11 + 10 + 11")
    expr.reducible?.should eq(true)

    expr.reduce(EMPTY_ENV).reducible?.should eq(true)
    expr.reduce(EMPTY_ENV).reduce(EMPTY_ENV).reducible?.should eq(true)
    res = expr.reduce(EMPTY_ENV).reduce(EMPTY_ENV).reduce(EMPTY_ENV)
    res.reducible?.should eq(false)

    res.as(Expr::Number).value.should eq(42)

  end

  it "can be reduced until it is a number" do
    x = Expr::Add.new(
      Expr::Number.new(5),
      Expr::Number.new(9),
    )
    x.reducible?.should eq(true)

    expr = Expr::Add.new(x, Expr::Add.new(x, x))

    puts "\n\n\texpr: #{expr}"
    while expr.reducible?
      puts "\t    : #{expr}"
      expr = expr.reduce(EMPTY_ENV)
    end

    expr.as(Expr::Number).value.should eq(42)

  end

end
