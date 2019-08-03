class Expr::Multiply < Expr::Op
  getter left
  getter right

  def initialize(@left : OpOrNum, @right : OpOrNum); end

  def to_s(io : IO)
    io << left << " x " << right
  end


  def reduce() : self | Expr::Number
    return Multiply.new(left.reduce, @right) if left.reducible?
    return Multiply.new(left, right.reduce) if right.reducible?

    l = @left.as Expr::Value
    r = @right.as Expr::Value
    return Number.new(l.value * r.value)
  end

end


