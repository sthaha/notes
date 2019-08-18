require "./vm_helper"


describe Vm::Machine do
  it "accepts an expression" do
    x = Expr::Add.new( Expr::Number.new(5), Expr::Number.new(9),)
    expr = Expr::Add.new(x, Expr::Add.new(x, x))

    vm = Vm::Machine.new(expr, Vm::EMPTY_ENV)
    vm.expression.reducible?.should be_true
    vm.run
    # NOTE: second time it doesn't rerun the expression
    vm.expression.reducible?.should be_false
    vm.run
    vm.expression.as(Expr::Number).value.should eq(42)

  end

  it "can only evaluate left to right" do
    # 2 + 3 x 5 + 4 should be 19 but gets 45 due to left -> right eval
    expr =  Expr::Multiply.new(
      Expr::Add.new( Expr::Number.new(2), Expr::Number.new(3),),
      Expr::Add.new( Expr::Number.new(5), Expr::Number.new(4),),
    )

    vm = Vm::Machine.new(expr, Vm::EMPTY_ENV)
    vm.run
    vm.expression.as(Expr::Value).value.should_not eq(19)
  end

  it "can evaluate variables" do
    expr =  Expr::Add.new(
      Expr::Add.new( Expr::Var.new("x"), Expr::Number.new(3)),
      Expr::Multiply.new( Expr::Number.new(5), Expr::Var.new("y")),
    )

    env = {} of String => Expr::Any
    env["x"] = Expr::Number.new(2)
    env["y"] = Expr::Number.new(8)

    vm = Vm::Machine.new(expr, env)
    vm.run
    vm.expression.as(Expr::Value).value.should_not eq(19)
  end
end

