
abstract class Expr::Any
  abstract def reducible? : Bool
  abstract def reduce

  abstract def to_s(io : IO)
  abstract def inspect(io : IO)
end

abstract class Expr::Value < Expr::Any
  abstract def value : Int32

  def to_s(io : IO)
    io << value
  end

  def inspect(io : IO)
    io << "<< " << "#{self}" << " >>"
  end

  def reducible?
    return false
  end

  def reduce
    return self
  end

end

class Expr::Number < Expr::Value
  getter value : Int32

  def initialize(@value : Int32); end
end


abstract class Expr::Op < Expr::Any
  abstract def reduce : Expr::Op | Expr::Value

  def reducible?
    true
  end

end

alias OpOrNum = Expr::Op | Expr::Number


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

  def reduce() : Expr::Add | Expr::Number
    return Add.new(left.reduce, @right) if left.reducible?
    return Add.new(left, right.reduce) if right.reducible?

    l = @left.as Expr::Value
    r = @right.as Expr::Value
    return Number.new(l.value + r.value)
  end

end
