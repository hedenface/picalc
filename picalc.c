#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#define PI (3.14159265358979323846264338327950288419716939937510582097494459)
#define MAX_NUM 99999

int main();

int main() {
    float pi = PI;
    float numerator = 22.0;
    float denominator = 7.0;
    float approx = (numerator / denominator);
    float pi_diff = fabsf(pi - approx);

    float n, d = 0;

    printf("pi = %f\n", pi);
    printf("(%f / %f) = (%f), difference from pi = %f\n", numerator, denominator, approx, pi_diff);

    for (d = 1; d <= MAX_NUM; d++) {

        if ((n / d) > pi) {
            continue;
        }

        for (n = 1; n <= MAX_NUM; n++) {

            approx = (n / d);

            if (approx > pi) {
                break;
            }

            float this_diff = fabsf(pi - approx);

            if (this_diff <= pi_diff) {
                numerator = n;
                denominator = d;
                pi_diff = this_diff;
            }
        }
    }

    printf("(%f / %f) = (%f), difference from pi = %f\n", numerator, denominator, approx, pi_diff);
}
