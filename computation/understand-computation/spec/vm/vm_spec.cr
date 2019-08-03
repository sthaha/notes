require "./vm_helper"


describe Vm::Machine do
  it "accepts an expression" do
    x = Expr::Add.new(
      Expr::Number.new(5),
      Expr::Number.new(9),
    )
    expr = Expr::Add.new(x, Expr::Add.new(x, x))

    vm = Vm::Machine.new(expr)
    vm.run
    # NOTE: second time it doesn't rerun the expression
    vm.run

  end

  it "can reduce an expression" do
  end
end

