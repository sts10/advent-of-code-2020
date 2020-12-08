# def what_colors_can_contain_this_color(color, a_big_hash)
#   all_colors = []
#   if a_big_hash[color]
#     first_order_colors = a_big_hash[color]
#     puts "First order colors found: " + first_order_colors.to_s
#   else
#     return []
#   end
#   all_colors << first_order_colors
#   first_order_colors.each do |this_color|
#     # second_order_colors = what_colors_can_contain_this_color(this_color, a_big_hash)
#     if what_colors_can_contain_this_color(this_color, a_big_hash).include?("no other")
#       return all_colors
#     else
#       all_colors << what_colors_can_contain_this_color(this_color, a_big_hash)
#       puts "At this point, all_colors is " + all_colors.to_s
#     end
#   end
# end

def get_first_order_colors(color, a_big_hash)
  first_order_colors = []

  a_big_hash.each do |k, v|
    if v.include?(color)
      first_order_colors << k
    end
  end

  return first_order_colors.uniq
end

def what_colors_can_be_contained_by_this_color(color, a_big_hash)
  all_colors = []
  first_order_colors = get_first_order_colors(color, a_big_hash)

  puts "for " + color + ": first order colors found: " + first_order_colors.to_s

  # if first_order_colors == []
  #   return all_colors
  # end

  all_colors << first_order_colors

  first_order_colors.each do |this_color|
    second_order_colors = what_colors_can_be_contained_by_this_color(this_color, a_big_hash)
    if second_order_colors == []
      return all_colors
    else
      all_colors << second_order_colors
      puts "At this point, all_colors is " + all_colors.to_s
    end
  end
end

rules = []
File.open("input").each_line do |line|
  rules << line
end

big_hash = Hash.new

rules.each do |rule|
  key = rule.split("contain")[0].split(" ")[0..1].join(" ")
  puts "in rule " + rule + " found key of " + key
  if rule.split("contain")[1].include?("no other")
    values = ["no other"]
  else
    values = rule.split("contain")[1].split(",").map { |a| a.split(" ")[1..2].join(" ") }
  end
  puts "found values of " + values.to_s
  big_hash[key] = values
end

color_to_find = "shiny gold"
# ans_colors = self.what_colors_can_contain_this_color("shiny gold", big_hash)
ans_colors = what_colors_can_be_contained_by_this_color(color_to_find, big_hash).flatten
puts "ans_colors is " + ans_colors.to_s
ans = ans_colors.length()
puts "Part 1 ans is " + ans.to_s

puts "-------------------"

def given_array_find_all_first_order_colors(colors_to_find, big_hash)
  colors_to_find.map { |color_to_find| get_first_order_colors(color_to_find, big_hash) }.flatten
end

colors_to_find = ["shiny gold"]
c = []
f = given_array_find_all_first_order_colors(colors_to_find, big_hash)
puts "f is " + f.to_s
c = c + f
# puts c
loop do
  c_len_before = c.uniq.length()
  c.each do |color|
    f = get_first_order_colors(color, big_hash)
    c = c + f
  end
  c_len_after = c.uniq.length()
  if c_len_before == c_len_after
    break
  end
end
# end
puts c.uniq
puts "len is " + c.uniq.length().to_s
