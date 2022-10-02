if __name__ == "__main__":
    # Get threes and fives
    threes = range(3, 1000, 3)
    fives = range(5, 1000, 5)
    # To eliminate duplicates, create a set of one and union the other with it
    unique_vals = set(threes).union(set(fives))
    print(sum(unique_vals))
