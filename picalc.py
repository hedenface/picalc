from math import fabs

pi = 3.14159265358979
max_num = 99999

def main():
    numerator = 22.0
    denominator = 7.0

    approx = (numerator / denominator)

    pi_diff = fabs(pi - approx)

    print(f"pi = {pi}")

    print(f"({numerator} / {denominator}) = ({approx}), difference from pi = {pi_diff}")

    n = 0

    for d in range(1, max_num):

        if (n / d) > pi:
            continue

        for n in range(1, max_num):

            approx = (n / d)

            if approx > pi:
                break

            this_diff = fabs(pi - approx)

            if this_diff <= pi_diff:
                numerator = n
                denominator = d
                pi_diff = this_diff

    print(f"({numerator} / {denominator}) = ({approx}), difference from pi = {pi_diff}")

main()
