# kidbase10: Simple Encoding a Child Can Do!

`kidbase10` is a simple encoding protocol that can be written with paper and pencil.
It uses a short encoding table to turn words and phrases into a series of 8-digit, base-10 numbers.

The protocol is meant solely for teaching purposes, specificially children aged 10 or younger.
Implementing the protocol is a suitable exercise for an introductory programming course.

## How it Works

### The Problem: Alice, Bob, and the Calculator

Imagine a student - Alice - wants to send a message to their classmate, Bob. Normally Alice would
write a note on a piece of paper to Bob. However, both of them recently got in trouble for passing
notes in class, and aren't allowed to pass paper any more. Alice and Bob are working on a math
project together, sharing a calculator with room for 8 digits. Can they continue passing their
messages to each other?

### The Solution: Encoding Words!

Alice and Bob can indeed do this if they _encode_ letters to numbers on their calculator, following
the `kidbase10` encoding procedure. `kidbase10` encodes letters by assigning a set of letters to a
corresponding base-10 digit. Because only 9 possible digits are allowed (0 is set aside for
spaces), the things Alice and Bob can say to each other are limited. Can you use your creativity to
help Alice and Bob send notes?

### The Procedure

`kidbase10` uses the following table as its "alphabet":

| Letter | Space | A | E | H | I | L | N | R | S | T |
| ------ | ----- | - | - | - | - | - | - | - | - | - |
| Number | 0     | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |

To encode a word to a number, each letter in the word is typed from left to right as its
corresponding digit on the calculator. `kidbase10` has a special rule that every place value on the
screen (8 in a typical school calculator) must be filled. If a word has less than 8 letters, the
number 0 is used to fill the rest of the screen.

For example:

- `LANE` encodes to 51620000
- `I STARE` encodes to 40891720

Longer words and phrases can be split into chunks of 8 characters/places. For example, the phrase
`I STARE AT THE HILL` encodes to the following three lines/entries:

```
40891720
19093203
45500000
```

A leading "0" is sometimes needed to represent a space. The calculator may not allow this, in which
case using numbers with 7 or fewer place values is allowed. Missing "0" digits on the left are
interpreted as spaces. Trailing "0"'s are still required to fill remaining space on the right.

For example, if ` A HILL` is the last portion of a message (leading space), the encoded value is
`01034550` to preserve the leading space and pad the encoding on the right.

### Decoding Numbers

To decode the numbers to words and phrases, the encoding "alphabet" table is used in reverse. For
each numeric digit, the corresponding letter (or space) is found on the table. Any "trailing" 0
digits used to fill the calculator screen are ignored when decoding. Leading "0" digits, or missing
digits if the number is less than 10 million, cannot be ignored and are treated as spaces.

Using our examples above:

- 51620000 decodes to `LANE`
- 40891720 decodes to `I STARE`

and

```
40891720
19093203
45500000
```

decodes to `I STARE AT THE HILL`.

## Specification

The full technical specification can be found in [SPEC.md](./SPEC.md). 

The specification is designed to be flexible - for example:

- Allowing alternate characters as input (including non-Latin
characters)
- Alternative forms of character input (ex: Latin uppercase and lowercase)
- Variable line length for encoding/decoding.

## License

```
   Copyright 2024 Adam B Kaplan

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```