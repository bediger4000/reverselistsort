# Daily Coding Problem: Problem #618 [Hard]

Given a list, sort it using this method: reverse(lst, i, j),
which reverses lst from i to j.

## Analysis

The problem is so loosely stated that I could merely
do a [bubble sort](bubble.go),
using `reverse(array, i, j)` to swap a smaller and a larger element.

If the point of this problem as an interview question is
to determine if a candidate knows a sorting algorithm,
and has a flash of insight (that `reverse(i, i+1)` works the same
as `swap(i, i+1)`), then this is a fine interview question.

If that's not the point, then all the interview has is the
experience of getting gamed by a candidate trying to figured out
what they can get away with.

The "[Hard]" annotation makes me think that's not what
this question was to provide.
I can't google up any other solutions than a bubble sort.
