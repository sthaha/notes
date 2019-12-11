
class Statement::Noop < Statement::Any
  def reducible(); false; end
end
