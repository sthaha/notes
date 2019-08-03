require "./expr_helper"


describe Expr::Multiply do
  it "prints fine" do
    x = Expr::Multiply.new(
      Expr::Number.new(11),
      Expr::Number.new(31),
    )
    "#{x}".should eq("11 x 31")
  end

  it "can be reduced" do
    x = Expr::Multiply.new(
      Expr::Number.new(11),
      Expr::Number.new(31),
    )

    x.reduce.as(Expr::Number).value.should eq(341)
  end

  it "can be reduced multiple times" do
    x = Expr::Multiply.new(
      Expr::Number.new(10),
      Expr::Number.new(11),
    )
    x.reducible?.should eq(true)

    res = Expr::Multiply.new(x, x)

    "#{res}".should eq("10 x 11 x 10 x 11")
    res.reducible?.should eq(true)

    res.reduce.reducible?.should eq(true)
    res.reduce.reduce.reducible?.should eq(true)
    res.reduce.reduce.reduce.reducible?.should eq(false)

    res.reduce.reduce.reduce.as(Expr::Number).value.should eq(12_100)

  end

  it "can be reduced until it is a number" do
    x = Expr::Multiply.new(
      Expr::Number.new(5),
      Expr::Number.new(8),
    )
    x.reducible?.should eq(true)

    expr = Expr::Multiply.new(x, Expr::Multiply.new(x, x))
    while expr.reducible?
      expr = expr.reduce
    end

    # 40 x 40 x 40
    expr.as(Expr::Number).value.should eq(64_000)
  end

end

