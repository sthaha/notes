require "./expr_helper"

describe Expr::LessThan do
  it "prints fine" do
    x = Expr::LessThan.new(
      Expr::Number.new(11),
      Expr::Number.new(31),
    )
    "#{x}".should eq("11 < 31")
  end

  it "can be reduced" do
    x = Expr::LessThan.new(
      Expr::Number.new(11),
      Expr::Number.new(31),
    )
    x.reduce(EMPTY_ENV).as(Expr::Boolean).value.should be_true
  end

  it "can be reduced multiple times" do
    expr = Expr::LessThan.new(
      Expr::Number.new(10),
      Expr::Add.new(
        Expr::Number.new(10),
        Expr::Number.new(5),
      ),
    )
    expr.reducible?.should eq(true)

    "#{expr}".should eq("10 < 10 + 5")
    expr.reducible?.should be_true

    expr.reduce(EMPTY_ENV).reducible?.should be_true
    expr.reduce(EMPTY_ENV).reduce(EMPTY_ENV).reducible?.should be_false
    res = expr.reduce(EMPTY_ENV).reduce(EMPTY_ENV)
    res.reducible?.should be_false
    res.as(Expr::Boolean).value.should be_true

  end

  it "can be reduced until it is a number" do
    expr = Expr::LessThan.new(
      Expr::Multiply.new(Expr::Number.new(5), Expr::Number.new(8)),
      Expr::Add.new(Expr::Number.new(5), Expr::Number.new(8)),
    )
    expr.reducible?.should be_true
    while expr.reducible?
      expr = expr.reduce(EMPTY_ENV)
    end

    expr.as(Expr::Boolean).value.should be_false
  end

end

