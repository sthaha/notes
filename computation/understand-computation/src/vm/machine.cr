module Vm
  EMPTY_ENV = {} of String => Expr::Any
end

class Vm::Machine
  getter expression

  def initialize(@expression : Expr::Any, @env : Hash(String, Expr::Any)); end

  def step
    @expression = @expression.reduce(@env)
  end

  def run
    expr = expression

    puts "\n\n\t Env:"
    puts "#{@env.map{ |k, v| "\t    | #{k}: #{v}" }.join("\n")}"

    puts "\n\tcompute: #{expr}"
    while expression.reducible?
      puts "\t    | #{expression}"
      step
    end

    puts "\t-----------------------------"
    puts "\t#{expr} : #{expression}"
    puts "\t-----------------------------"
  end

end
