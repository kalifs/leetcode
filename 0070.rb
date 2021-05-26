# @param {Integer} n
# @return {Integer}
def climb_stairs(n)
    ones = n
    twos = 0
    total = 0
    while true
        break if ones <= 1
        total += combinations(ones,twos)
        twos += 1
        ones -= 2
    end
    total += combinations(ones,twos)
    # puts("total: #{total}")
    return total
end

def combinations(ones, twos)
    size = ones + twos
    most = [ones, twos].max
    least = [ones,twos].min
    n = 1
    size.downto(most+1) do |i|
         n = n * i
    end
    r = 1
    1.upto(least) do |i|
         r = r * i
    end
    total = n / r
    # puts("ones: #{ones} twos: #{twos} r: #{r} n: #{n} total: #{total}")
    return total
end



# RECURSIVE VARIANT
# @param {Integer} n
# @return {Integer}
def climb_stairs(n)
    return combinations(n,0)
end

def combinations(ones, twos)

    size = ones + twos
    most = [ones, twos].max
    least = [ones,twos].min
    n = 1
    size.downto(most+1) do |i|
         n = n * i
    end
    r = 1
    1.upto(least) do |i|
         r = r * i
    end
    total = n / r
    # puts("ones: #{ones} twos: #{twos} r: #{r} n: #{n} total: #{total}")
    return total if ones <= 1
    return total + combinations(ones - 2, twos + 1)
end
