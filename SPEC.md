# kidbase10 Encoding Specification

The "kid Base 10" (`kidbase10`) encoding is designed to represent arbitrary text containing a
subset of English characters as a series of 9 digit numbers written in base-10 format. It is
intended to be simple enough for a child aged 10 or younger to implement with paper and pencil.
The protocol can be used to teach computer science concepts in introductory education settings. It
is not meant to process arbitrary text or binary input.

This specification uses formal language found in industry standard documents, such as RFC-4648 [1].
It is not recommended for use in primary education environments.

## Character Input

For input, the following characters are allowed in any sequence and any length:

```
A, E, H, I, L, N, R, S, T, (whitespace)
```

The characters above are drawn from the most common vowels and consonants in written English [2].
Future versions or variants of this specification MAY consider a different set of allowed
characters. The maximum number of allowed characters is 10.

Input characters MAY have their lowercase forms as input. Lowercase input letters MUST be converted
to uppercase form before any encoding begins. Case preservation is not required when encoding and
decoding text.

Whitespace characters are allowed for the purpose of separating words and encoding phrases of
arbitrary length. The underscore character `_` MAY be used as a substitute for a whitespace input
character. Preserving specific whitespace values is NOT required when encoding and decoding text.


## Encoding Alphabet

The encoding alphabet has 10 entries, one per base-10 digit. The input letters from above are placed in
alphabetical order on the table. Zero (0) is reserved for the whitespace character.

| base10 Digit | Character |
| ------ | --------- |
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

The encoding alphabet MAY be presented in a transposed format for written or presentation materials.
This can help visual learners understand the relationship between numbers and letters when encoding
or decoding data.

## Encoding Procedure

The encoding process represents groups of 9 characters ("9-char") as their corresponding 9 digit
base-10 number. At the beginning of the procedure, all characters must be transformed to their
uppercase format. Whitespace characters of any form MUST be replaced with the standard whitespace/
space character. The underscore `_` character MUST also be replaced with the standard whitespace/
space character.

After the initial processing above, each character in the resulting "9-char" is replaced with its
corresponding base 10 digit in the encoding alphabet. If a character is not present in the encoding
alphabet, the procedure should stop and report an "Invalid Character" error. If the input character
group has fewer than 9 characters, the 0 digit is added on the right until the resulting number has
9 digits in base-10 representation.

If a whitespace character is encountered at the beggining of a 9-char group, it should be encoded
as a leading 0 digit in the resulting base-10 number.

Encoding output MAY be one of three forms:

1. Single line of text characters representing each digit. This format MAY NOT be interpreted as an
   integer by consuming systems.
2. Multiple lines of 9-character text digits. In this format, leading whitespace characters MUST be
   represented as a leading 0 digit. Each line MAY NOT be interpreted as an integer by consuming
   systems.
3. An array of 8 or 9-digit base-10 integers. This format MAY be interpreted as an array of
   unsigned integers by consuming systems. The leading 0 for whitespace MAY be omittied in this
   format.

Trailing 0 digits may not be omitted under any circumstance.

## Decoding Procedure

Decoding proceeds as follows:

1. If the input is a single stream of text digits, the stream MUST be broken up into an array
   of 9-char text digits first.
2. If the input is multiple lines of 9-char text, each newline MUST be interpreted as a new entry
   in an array of 9-char text digits.

Each entry in the text array is then parsed to its corresponding unsigned base-10 integer. Leading
"0" characters MUST NOT result in an error when parsing. The process above MAY be skipped when the
decoding input is an array of integers.

When decoding integers within the array input, each integer MUST BE greater than or equal to 0, and
MUST BE less than 1 billion (1000000000). Decoding an integer < 0 or >= 1000000000 MUST result in
an error and halt the decoding procedure.

For each base-10 integer, the base digit is replaced with its corresponding text character in the
encoding alphabet. For integers less than 100 million (100000000), a leading whitespace character
MUST be added first. Whitespace characters can be represented either as the standard space character `" "`
OR as the underscore character `"_"`. The resulting text SHOULD represent a 9-character string of
text characters (9-char).

This process above MUST be repeated for each base-10 integer in the input array, in order. Each
resulting 9-char MUST be appended to the 9-char generated from the previous integer in the input
array. After decoding the entire array, any trailing whitespace characters MAY be removed when
rendering output as text.

## Examples of Encoding and Decoding

| Input | Encoded | Decoded |
| ----- | ------ | ------ |
| "Lane" | `516200000` | `LANE` |
| "I stare at the hill" | `408917201909320345500000000` | `I STARE AT THE HILL` |

```sh
$ kidbase10 "I stare at the hill"
408917201
909320345
500000000
$ kidbase10 -D 516200000
LANE
```

[1] https://datatracker.ietf.org/doc/html/rfc4648#section-4

[2] https://www.thesaurus.com/e/ways-to-say/most-common-letter/