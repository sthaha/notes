class Expr::Add < Expr::Op
  getter left
  getter right

  def initialize(@left : OpOrNum, @right : OpOrNum); end

  def to_s(io : IO)
    io << left << " + " << right
  end

  def inspect(io : IO)
    io << "<< " << self << " >>"
  end

  def reduce() : self | Expr::Number
    return Add.new(left.reduce, @right) if left.reducible?
    return Add.new(left, right.reduce) if right.reducible?

    l = @left.as Expr::Value
    r = @right.as Expr::Value
    return Number.new(l.value + r.value)
  end

end

