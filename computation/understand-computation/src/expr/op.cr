
abstract class Expr::Op < Expr::Any
  abstract def reduce(env) : Expr::Op | Expr::Value

  def reducible?
    true
  end
end

alias OpOrNum = Expr::Op | Expr::Number

