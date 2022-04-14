'''
I heard that summing numbers in [0,1] until the total is >1 will
require an average of e numbers.

This code puts that to the test.

Hopefully pretty graphs, other visualisations, and optimisations
to be added later.
'''

import random
import math
# import matplotlib.pyplot as plt

# Generate a number in the range [0,1] inclusive
def generate_number():
    return random.uniform(0, 1)

# One iteration of the theory; add numbers until their sum is >1, then
# return the amount of numbers needed
def sum_until_1():
    total, count = 0, 0
    while total < 1:
        total += generate_number()
        count += 1
    return count

if __name__ == "__main__":
    n = 1000 # Number of iterations
    total_count = 0 # Used to calculate average
    for i in range(n):
        total_count += sum_until_1()

    print("{} iterations: avg={} e={}".format(n, total_count/n, math.e))
