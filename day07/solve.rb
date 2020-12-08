def what_colors_can_contain_this_color(color, a_big_hash)
  all_colors = []
  if a_big_hash[color]
    first_order_colors = a_big_hash[color]
    puts "First order colors found: " + first_order_colors.to_s
  else
    return []
  end
  all_colors << first_order_colors
  first_order_colors.each do |this_color|
    # second_order_colors = what_colors_can_contain_this_color(this_color, a_big_hash)
    if what_colors_can_contain_this_color(this_color, a_big_hash).include?("no other")
      return all_colors
    else
      all_colors << what_colors_can_contain_this_color(this_color, a_big_hash)
    end
  end
end

rules = []
File.open("test_input").each_line do |line|
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
ans_colors = what_colors_can_contain_this_color(color_to_find, big_hash)
puts "ans_colors is " + ans_colors.to_s
ans = ans_colors.length()
puts "Part 1 ans is " + ans.to_s