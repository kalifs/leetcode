# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
    (0...(nums.size)).each do |i|
       ((i+1)...nums.size).each do |j|
          return [i,j] if nums[i]+nums[j] == target
       end
    end
    return []
end
