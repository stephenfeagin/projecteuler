# First get the multiples of 3 below 1000
threes <- seq(3, 999, by = 3)
# Then the multiples of 5
fives <- seq(5, 995, by = 5)
# Take only the unique values to eliminate those that are multiples of 15
unique_vals <- unique(c(threes, fives))
sum(unique_vals)
