require "./expr_helper"

describe Number do
  it "prints fine" do
    x = Expr::Number.new(10)
    "#{x}".should eq("10")
  end
end
