# Binary Coded Decimal Arithmetic Experiments


https://grace.bluegrass.kctcs.edu/~kdunn0001/files/Arithmetic_Operations_and_Circuits/
0110 mentioned

https://tutorials.eeems.ca/ASMin28Days/lesson/day18.html
0110 ?

```
struct FP {
    byte  sign;           // Whether the number is positive or negative
    byte  exponent;       // Locates the decimal point
    byte  significand[7]; // The number itself
    byte  guard[2];       // Guard digits for mathematics
};
```

The magnitude of every real number except zero can be represented as
m × 10exp,
where exp is an integer designating the exponent and m is a real number
designating the significand such that 1 <= m < 10.

Sign
This byte determines if the number evaluates as positive or negative, and also if it is real or complex. For the uninitiated, a complex number is one of the form a+bi, where i is the square root of -1.

    %00000000 — Positive and real.
    %10000000 — Negative and real.
    %00001100 — Positive and complex.
    %10001100 — Negative and complex. 

Exponent
The exponent field reports the power of ten that the mantissa is to be raised.
The number format is not the usual two's complement, but rather biased to $80.
A value of $80 trasnslates as 100. $81 is 101. $7F is 10-1.

Significand
These are the digits of the number. Each nibble specifies one decimal digit, so
you can have a floating-point number with 14 digits. The first digit, and only
the first digit, is the characteristic (the whole part) with the remainder
being the mantissa (the decimal part).
