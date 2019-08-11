
alias Matrix = Array(Array(Int32))
alias Visited =  Hash(Int32, Hash(Int32, Int32))

class Groups::Finder
  getter matrix
  getter groups
  getter target

  def initialize(@matrix : Matrix, @target=1)
    @visited = Visited.new()
    @groups = 0
  end

  def rows
    @matrix.size
  end

  def cols
    @matrix[0].size
  end

  def run()
    # puts "#{rows} x #{cols}: #{matrix}"
    row, col = 0, 0
    while row < rows && col < cols
        x = matrix[row][col]
        #puts "[#{@groups}] :  [#{row}][#{col}]: #{x}"
        row, col, found = process(row, col)
        @groups += 1 if found
    end

    return @groups
  end


  def process(r, c)
    x = matrix[r][c]
    # puts " .. [#{@groups}]  :  [#{r}][#{c}]: #{x}"

    if visited?(r, c)
      r, c = nextCell(r, c)
      return r, c, false
    end

    if x != target
      visit r, c
      r, c = nextCell(r, c)
      return r, c, false
    end

    visit(r, c, @groups)
    neighbours(r,c).each { |(row, col)| process(row, col)}
    r, c = nextCell(r, c)
    return  r, c, true

  end

  def neighbours(row, col)

    ns = [
      {row-1, col-1},
      {row-1, col  },
      {row-1, col+1},

      {row,   col-1},
      {row,   col+1},

      {row+1, col-1},
      {row+1, col  },
      {row+1, col+1},
    ]

  return ns.select { |(r, c)|
    (0...rows).includes?(r) &&
    (0...cols).includes?(c) &&
    !visited?(r,c)
  }

  end

  def visit(r, c, group = -1)
    @visited[r] ||= {} of Int32 => Int32
    @visited[r][c] = group
  end


  def nextCell(r, c)
    while visited?(r, c) && r < rows
      c = c + 1
      if c >= cols
        r = r + 1
        c = 0
      end
    end
    return r, c
  end


  def visited?(r, c)
    return false unless @visited.has_key? r
    return false unless @visited[r].has_key? c
    return true
  end

end

module Groups
  def self.find(m : Matrix) : Int32
    f =  Finder.new(m)
    return f.run()
  end

end
