# If we list all the natural numbers below 10 that are multiples of 3 or 5,
# we get 3, 5, 6, and 9. The sum of these multiples is 23.
#
# Find the sum of all the multiples of 3 or 5 below 1000

# First get the multiples of 3 below 1000
threes <- seq(3, 999, by = 3)
# Then the multiples of 5
fives <- seq(5, 995, by = 5)
# Take only the unique values to eliminate those that are multiples of 15
unique_vals <- unique(c(threes, fives))
sum(unique_vals)
