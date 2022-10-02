# Running total sum
# Start with 2 because it's already in the first two terms
total <- 2
# Keep track of the last two terms
last_two <- c(1, 2)

# Start loop
while (TRUE) {
    # Find the value of the next term
    next_term <- sum(last_two)
    # If we get above 4 million, stop the loop and don't add it to the total
    if (next_term >= 4e6) {
            break
        }
    # Re-record the last two values for use in the next iteration
    last_two <- c(last_two[2], next_term)
    # If the new term is even, add it to the running total
    if (next_term %% 2 == 0) {
        total <- total + next_term
    }
}

print(total)
