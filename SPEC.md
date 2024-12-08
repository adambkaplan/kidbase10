# kidbase10 Encoding Specification

The "kid Base 10" (`kidbase10`) encoding is designed to represent
arbitrary text containing a subset of text characters as a series of
numbers written in base-10 format. It is intended to be simple enough
for a child aged 10 or younger to implement with paper and pencil. The
protocol can be used to teach computer science concepts such as
encoding, cryptographic ciphers, and base number systems in primary
education settings. It is not meant to process arbitrary text or binary
input.

This specification document uses formal language found in industry
standards, such as RFC-4648 [1]. It is not recommended for use in
primary education environments, but may be suitable for an
undergraduate programming or Computer Science course.

## Character Input

For input, the following characters are allowed in any sequence and any
length:

```
A, E, H, I, L, N, R, S, T, (whitespace)
```

The characters above are drawn from the most common vowels and
consonants in written English [2]. A whitespace character, or a
character that serves the same function as a whitespace character, MUST
be present in the set of allowed characters. The set of allowed input
characters MUST have length ten (10).

### Whitespace Character

A whitespace character MUST be included in the set of allowed inputs
for the purpose of separating words and encoding phrases of arbitrary
length. The ASCII space character (" ") is the canonical whitespace
character for encoding input and decoding output.

## Encoding Alphabet

The encoding alphabet MUST have 10 entries, one per base-10 digit. The
input letters from above are placed in alphabetical order on the table.
Zero (0) is reserved for the whitespace character:

| base-10 Digit | Character |
| ------------- | --------- |
| 0 | Space |
| 1 | A |
| 2 | E |
| 3 | H |
| 4 | I |
| 5 | L |
| 6 | N |
| 7 | R |
| 8 | S |
| 9 | T |

The encoding alphabet MAY be presented in a transposed format for
written or presentation materials. This can help visual learners
understand the relationship between numbers and letters when encoding
or decoding data.

## Encoding Procedure

The encoding process represents groups of N characters ("N-char") as
their corresponding N digit base-10 number. The canonical
implementation processes 8 characters at a time ("8-char"). Variants
MAY pick any value of N that is suitable for integer arithmetic
operations within a given application.

During encoding, each character in the N-char group is replaced with
its corresponding base-10 digit in the encoding alphabet. If a
character is not present in the encoding alphabet, the procedure MUST
stop and report an "invalid character" error.

If the input character group has fewer than N characters, the 0 digit
MUST be appended on the right until the resulting number has N digits
in base-10 representation.

If a whitespace character is encountered at the beggining of an N-char
group, it MUST be encoded as a leading 0 digit in the resulting base-10
number.

### Encoding Output

Encoding output MAY be one of three forms:

1. "Character stream": a single line of text characters representing
   each digit. This format MAY NOT be interpreted as an integer by
   consuming systems.
2. "Multi-line text": multiple lines of N-character text digits. In
   this format, leading whitespace characters MUST be represented as a
   leading 0 digit. Each line MAY NOT be interpreted as an integer by
   consuming systems.
3. "Integer array": an array of base-10 integers. Encoders MAY choose
   any formatting scheme that produces arrays of unsigned integers,
   such as JSON and YAML. This format MAY be interpreted as an
   array of unsigned integers by consuming systems. Leading 0 digits
   for whitespace MAY be omittied in this format.

Trailing 0 digits MUST NOT be omitted under any circumstance.

## Decoding Procedure

Decoding proceeds as follows:

1. If the input is a single stream of text digits, the stream MUST be
   broken up into an array of N-char text digits first.
2. If the input is multiple lines of N-char text, each newline MUST be
   interpreted as a new entry in an array of N-char text digits.

Each entry in the text array is then parsed to its corresponding
unsigned base-10 integer. Leading "0" characters MUST NOT result in an
error when parsing. The process above MAY be skipped when the decoding
input is an array of unsigned integers.

When decoding integers within the array input, each integer MUST BE
greater than or equal to 0, and MUST BE less than 10^N. Decoding an
integer < 0 or >= 10^N MUST result in an "invalid format" error and
halt the decoding procedure.

For each base-10 integer, the base digit when read from left to right
MUST BE replaced with its corresponding text character in the encoding
alphabet. For integers less than 10^(N-1), leading whitespace
characters MUST be added for any higher base-10 order of magnitude.

This process above MUST be repeated for each base-10 integer in the
input array, in order. Each resulting N-char MUST be appended to the
N-char generated from the previous integer in the array. After decoding
the entire array, any trailing whitespace characters MAY be removed
when rendering output as text.

## Appendix: Examples of Encoding and Decoding

### Canonical: 8-char, Latin Alphabet, Case Sensitive

The canonical procedure uses N=8 (8-char) for encoding and decoding,
the limited Latin alphabet described in this document, and the ASCII
space character for whitespace. Only uppercase Latin characters are
allowed for input and output.

#### Input: `LANE`

Encoded forms:
- Character stream and multi-line: `51620000`
- Integer array (JSON): [51620000]

Decoded: `LANE`

#### Input: `I STARE AT THE HILL`

Encoded forms:
- Character stream: `408917201909320345500000`
- Multi-line text:
  ```
  40891720
  19093203
  45500000
  ```
- Integer array (JSON): [40891720, 19093203, 45500000]

Decoded: `I STARE AT THE HILL`

### Alternative: 12-char, Case Insensitive, Underscore

The following examples use the following scheme:

- N=12 (12-char) for grouping
- Underscore "_" for the whitespace character
- Latin character set as described in this document, with lowercase and
  uppercase treated as equivalent. Uppercase is chosend for decoded
  output.

#### Input: `Lane`

Encoded forms:
- Character stream: `516200000000`
- Integer array (JSON): 516200000000

Decoded: `LANE`

#### Input: `I_Stare_At_The_Hill`

Encoded forms:
- Character stream: `408917201909320345500000`
- Multi-line text:
  ```
  408917201909
  320345500000
  ```
- Integer array (JSON): [408917201909, 320345500000]

Decoded: `I_STARE_AT_THE_HILL`

[1] https://datatracker.ietf.org/doc/html/rfc4648#section-4

[2] https://www.thesaurus.com/e/ways-to-say/most-common-letter/