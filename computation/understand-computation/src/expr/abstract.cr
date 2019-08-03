abstract class Expr::Any
  abstract def reducible? : Bool
  abstract def reduce

  abstract def to_s(io : IO)
  abstract def inspect(io : IO)

  def inspect(io : IO)
    io << "<< " << self << " >>"
  end
end

abstract class Expr::Value < Expr::Any
  abstract def value : Int32

  def to_s(io : IO)
    io << value
  end

  def reducible?
    return false
  end

  def reduce
    return self
  end

end


abstract class Expr::Op < Expr::Any
  abstract def reduce : Expr::Op | Expr::Value

  def reducible?
    true
  end
end

alias OpOrNum = Expr::Op | Expr::Number
