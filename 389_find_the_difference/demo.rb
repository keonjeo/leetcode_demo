#!/bin/env ruby

# @param {String} s
# @param {String} t
# @return {Character}
def find_the_difference(s, t)
  (t.bytes.sum - s.bytes.sum).chr
end

class Array
  def sum
    inject(0) { |sum, x| sum + x }
  end
end

a = "abcd"
b = "abcde"

# a="a"
# b="aa"

difference = find_the_difference(a, b)

puts "=================="
puts difference
puts "=================="