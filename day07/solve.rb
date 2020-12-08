def what_colors_can_contain_this_color(color, a_big_hash)
  all_colors = []
  arr = a_big_hash[color]
  all_colors << arr
  arr.each do |this_color|
    if what_colors_can_contain_this_color(this_color).include?("no other")
      return all_colors
    else
      all_colors << what_colors_can_contain_this_color(this_color)
    end
  end
end

rules = []
File.open("input").each_line do |line|
  rules << line
end

# big_hash = {}
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

# ans_colors = self.what_colors_can_contain_this_color("shiny gold", big_hash)
ans_colors = what_colors_can_contain_this_color("shiny gold", big_hash)
ans = ans_colors.length()
puts "Part 1 ans is " + ans.to_s
