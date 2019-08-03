class Expr::Number < Expr::Value
  getter value : Int32

  def initialize(@value : Int32); end
end


