class Expr::Var < Expr::Any
  getter name

  def initialize(@name : String)
  end

  def reducible?
    true
  end

  def reduce(env)
    return env[name]
  end

  def to_s(io : IO)
    io << name
  end

end


