
abstract class Expr::Op < Expr::Any
  abstract def reduce(env) : Expr::Op | Expr::Value | Expr::Bool

  def reducible?
    true
  end
end

alias OpOrNum = Expr::Op | Expr::Number

