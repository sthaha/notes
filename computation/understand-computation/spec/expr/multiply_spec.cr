require "./expr_helper"


describe Expr::Multiply do
  it "prints fine" do
    x = Expr::Multiply.new(
      Expr::Number.new(11),
      Expr::Number.new(31),
    )
    "#{x}".should eq("11 * 31")
  end

  it "can be reduced" do
    x = Expr::Multiply.new(
      Expr::Number.new(11),
      Expr::Number.new(31),
    )

    x.reduce(EMPTY_ENV).as(Expr::Number).value.should eq(341)
  end

  it "can be reduced multiple times" do
    x = Expr::Multiply.new(
      Expr::Number.new(10),
      Expr::Number.new(11),
    )
    x.reducible?.should eq(true)

    expr = Expr::Multiply.new(x, x)

    "#{expr}".should eq("10 * 11 * 10 * 11")
    expr.reducible?.should eq(true)

    expr.reduce(EMPTY_ENV).reducible?.should eq(true)
    expr.reduce(EMPTY_ENV).reduce(EMPTY_ENV).reducible?.should eq(true)
    res = expr.reduce(EMPTY_ENV).reduce(EMPTY_ENV).reduce(EMPTY_ENV)
    res.reducible?.should eq(false)

    res.as(Expr::Number).value.should eq(12_100)

  end

  it "can be reduced until it is a number" do
    x = Expr::Multiply.new( Expr::Number.new(5), Expr::Number.new(8))
    x.reducible?.should eq(true)

    expr = Expr::Multiply.new(x, Expr::Multiply.new(x, x))
    while expr.reducible?
      expr = expr.reduce(EMPTY_ENV)
    end

    # 40 * 40 * 40
    expr.as(Expr::Number).value.should eq(64_000)
  end

end

