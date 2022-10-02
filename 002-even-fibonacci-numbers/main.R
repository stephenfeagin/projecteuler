# Keep track of the most recent two terms
last_two <- c(1, 2)
# Running total sum
# Start with 2 because it's already in the first two terms
total <- 2
# Get the next term in the sequence
next_term <- sum(last_two)
while (next_term < 4e6) {
    if (next_term %% 2 == 0) {
        total <- total + next_term
    }
    last_two <- c(last_two[2], next_term)
    next_term <- sum(last_two)
}
print(total)
