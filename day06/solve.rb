text = ""
File.open("input").each_line do |line|
  text += line
end

array_as_family = text.split("\n\n")
puts "Found " + array_as_family.length().to_s + " families"
puts "One of which is " + array_as_family[7].to_s

# part 1
families_as_uniq = []
array_as_family.each do |family|
  this_family = family.split("").uniq
  this_family.delete("\n")
  families_as_uniq << this_family
end

sum = 0
families_as_uniq.each do |family|
  sum = sum + family.length()
end

puts "Answer to part one is " + sum.to_s

families_as_sub_arrays = []
array_as_family.each do |family|
  families_as_sub_arrays << family.split("\n")
end

# on to part two
sum = 0
families_as_sub_arrays.each do |family_arr|
  family_arr[0].split("").each do |letter|
    if family_arr.all? { |person| person.include?(letter) }
      sum = sum + 1
    end
  end
end

puts "Answer to part two is " + sum.to_s
