
class Expr::Number
  getter value : Int32
  def initialize(@value : Int32); end

  def to_s(io : IO)
    io << value
  end

  def inspect(io : IO)
    io << "<< "
    to_s(io)
    io << " >>"
  end
end
