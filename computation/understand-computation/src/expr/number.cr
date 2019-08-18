class Expr::Number < Expr::Value
  getter value : Int32

  def initialize(@value : Int32); end
end

class Expr::Bool < Expr::Value
  getter value : Bool

  def initialize(@value : Bool); end
end
