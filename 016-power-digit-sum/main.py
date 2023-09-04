import math
num_char = format(math.pow(2, 1000), ".0f")
digits_num = [int(x) for x in num_char]
sum(digits_num)