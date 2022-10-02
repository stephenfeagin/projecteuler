# Generator expressions make this easy
if __name__ == "__main__":
    seq = range(1, 101)  # We need to include 100 in the sequence
    sum_of_squares = sum(i**2 for i in seq)
    square_of_sum = sum(seq)**2
    print(square_of_sum - sum_of_squares)
    