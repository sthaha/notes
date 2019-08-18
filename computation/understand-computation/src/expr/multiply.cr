class Expr::Multiply < Expr::Op
  getter left
  getter right

  def initialize(@left : OpOrNum, @right : OpOrNum); end

  def to_s(io : IO)
    io << left << " * " << right
  end


  def reduce(env) : self | Expr::Number
    return Multiply.new(left.reduce(env), @right) if left.reducible?
    return Multiply.new(left, right.reduce(env)) if right.reducible?

    l = @left.as Expr::Number
    r = @right.as Expr::Number
    return Number.new(l.value * r.value)
  end

end


