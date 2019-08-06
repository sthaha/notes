require "./expr_helper"

describe Expr::Var do
  it "is a name" do
    v = Expr::Var.new("foobar")
    v.name.should be("foobar")
  end

  it "is reducible " do
    v = Expr::Var.new("foobar")
    v.reducible?.should be_true
  end


  it "can be reduced" do
    v = Expr::Var.new("x")

    expr = Expr::Add.new(Expr::Number.new(5), Expr::Number.new(9))
    v.reduce({x: expr}).should be(expr)
  end
end
