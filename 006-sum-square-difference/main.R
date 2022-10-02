# R makes calculations like these easy with builtin mathematical functions
# and out-of-the-box vectorization of many functions

sum_of_squares <- function(vec) {
    # The sum of the squared values in a vector
    sum(vec**2)
}
square_of_sum <- function(vec) {
    # The square of the sum of the values in a vector
    sum(vec)**2
}

vec <- 1:100
print(square_of_sum(vec) - sum_of_squares(vec))
