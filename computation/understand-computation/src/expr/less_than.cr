
class Expr::LessThan < Expr::Op
  getter left
  getter right

  def to_s(io : IO)
    io << left << " < " << right
  end

  def initialize(@left : OpOrNum, @right : OpOrNum); end

  def reduce(env) : self | Expr::Boolean
    puts "less than: #{left} < #{right}"
    return LessThan.new(left.reduce(env), @right) if left.reducible?
    return LessThan.new(left, right.reduce(env)) if right.reducible?

    l = @left.as Expr::Number
    r = @right.as Expr::Number
    return Boolean.new(l.value < r.value)
  end

end
