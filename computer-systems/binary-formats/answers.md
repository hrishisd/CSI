# Binary representations of data

## Hexadecimal

### Simple converesion

0x9, 0x88 and 0xf6

### CSS colors

Each hexadecimal digit represents 4 bits. A 6 character hex string corresponds to 24 bits.
Both encodings can represent 2^24 colors.

### Say hello to hellohex

We expect the file to be 34 hexadecimal digits since each hex digit corresponds to half a byte.

The first 5 hex characters from xxd -p hellohex are 0x68656
In binary: 0110 1000 0110 0101 0110

## Integeres
### Basic conversion
4 is 0b100
65 is 0b1000001
105 is 64 + 32 + 8 + 1 = 0b1101000
255 is 256 - 1 = 128 + 64 + ... + 1 = 0b011111111 or 0x7f

10 is 2
11 is 3
1101100 is 4 + 8 + 32 + 64 = 108
1010101 is 1 + 4 + 16 + 64 = 85

### Unsigned binary addition

0b10000000  or 0x100
If your registers are only 8 bits wide, they will overflow and the result will be 0b0

### Two's complement conversions

127 -128 -1 1 -14
0x0e 0x80 0xff 0xf2

10000011 -125
11000100 -60

### Addition of two’s complement signed integers
-1

## Byte ordering

### 9001

big endian