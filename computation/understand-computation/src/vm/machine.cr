

class Vm::Machine
  getter expression

  def initialize(@expression : Expr::Any); end
    def step
      @expression = @expression.reduce
    end

  def run
    expr = expression

    puts "\n\n\tcompute: #{expr}"
    while expression.reducible?
      puts "\t    | #{expression}"
      step
    end

    puts "\t-----------------------------"
    puts "\t#{expr} : #{expression}"
    puts "\t-----------------------------"
  end

end
