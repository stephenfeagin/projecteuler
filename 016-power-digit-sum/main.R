num <- format(2^1000, scientific = FALSE)
digits_char <- strsplit(num, "")[[1]]
digits_num <- as.numeric(digits_char)
sum(digits_num)