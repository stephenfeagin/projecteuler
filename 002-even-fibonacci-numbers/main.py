if __name__ == "__main__":
    # Start with the first two terms
    # The sum of the even numbers is already 2
    total = 2

    # Keep track of the most recent two numbers
    last_two = [1, 2]
    while True:
        # The next term in the sequence is the sum of the previous two
        next_term = sum(last_two)
        # Break the loop if we exceed 4 million
        if next_term >= 4000000:
            break
        # If the new term is even, add it to the running total
        if next_term % 2 == 0:
            total += next_term
        # Update the last_two list
        last_two = [last_two[1], next_term]
    print(total)