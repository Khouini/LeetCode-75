# 394. Decode String – Stack

---

## What's the Problem?

You have a string like `"3[a2[c]]"`.
Return the fully decoded string: `"accaccacc"`.

The encoding rule is: `k[encoded_string]` — repeat `encoded_string` exactly `k` times.
Brackets can be **nested**.

---

## Step-by-Step Thinking

### Question 1: What makes this problem hard?

The brackets can be **nested**.

```
3[a2[c]]
     ↑
  inner bracket must be resolved first
  before the outer one can be repeated
```

You cannot decode left-to-right in one pass because you hit the outer `3[` before you know what's inside it.

> Observation: whenever you face a structure where you need to "come back" to something after finishing an inner part, that's a classic **stack** signal.

---

### Question 2: What do we need to remember when we open a `[`?

When we see `[`, we are about to enter a nested group. Two things belong to the _outer_ context that we must save:

1. The **string built so far** (before this group)
2. The **repeat count** for this group

Without saving these, we'd lose the outer context once we go deeper.

> Observation: push both onto a stack when you see `[`. Pop both when you see `]`.

---

### Question 3: How do we handle multi-digit numbers?

`k` can be more than one digit (e.g. `12[ab]`).

When we see a digit, we don't push immediately — we accumulate:

```
currentNum = currentNum * 10 + digit
```

This handles `12` as twelve, not `1` then `2`.

---

### Question 4: What happens when we hit `]`?

Three things:

1. Pop the **repeat count** `k` and the **string built before this group** `prevStr` from the stack
2. Repeat the **current string** `k` times
3. Prepend `prevStr` to it — that's the new current string

```
prevStr + repeat(currentStr, k)  →  new currentStr
```

---

### Question 5: What do we do with regular letters?

Just append them to `currentStr`. Nothing special.

---

## Full Walkthrough

Input: `"3[a2[c]]"`

| Step | Char  | Action                                  | `currentNum` | `currentStr`  | `stack`               |
| ---- | ----- | --------------------------------------- | ------------ | ------------- | --------------------- |
| 0    | —     | initial state                           | `0`          | `""`          | `[]`                  |
| 1    | `'3'` | accumulate digit: `0*10+3`              | `3`          | `""`          | `[]`                  |
| 2    | `'['` | push `("", 3)`, reset                   | `0`          | `""`          | `[("", 3)]`           |
| 3    | `'a'` | append letter                           | `0`          | `"a"`         | `[("", 3)]`           |
| 4    | `'2'` | accumulate digit: `0*10+2`              | `2`          | `"a"`         | `[("", 3)]`           |
| 5    | `'['` | push `("a", 2)`, reset                  | `0`          | `""`          | `[("", 3), ("a", 2)]` |
| 6    | `'c'` | append letter                           | `0`          | `"c"`         | `[("", 3), ("a", 2)]` |
| 7    | `']'` | pop `("a", 2)` → `"a" + repeat("c", 2)` | `0`          | `"acc"`       | `[("", 3)]`           |
| 8    | `']'` | pop `("", 3)` → `"" + repeat("acc", 3)` | `0`          | `"accaccacc"` | `[]`                  |

**Result: `"accaccacc"`**

---

## Code

```go
func decodeString(s string) string {
    type frame struct {
        str   string
        count int
    }

    stack := []frame{}
    currentStr := ""
    currentNum := 0

    for _, ch := range s {
        switch {
        case ch >= '0' && ch <= '9':
            currentNum = currentNum*10 + int(ch-'0')

        case ch == '[':
            stack = append(stack, frame{currentStr, currentNum})
            currentStr = ""
            currentNum = 0

        case ch == ']':
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            repeated := strings.Repeat(currentStr, top.count)
            currentStr = top.str + repeated

        default:
            currentStr += string(ch)
        }
    }

    return currentStr
}
```

---

## Key Insight

> A `[` means "save where I am, go deeper".
> A `]` means "come back, repeat, and continue".
>
> The stack is just a way to pause the outer context while you resolve the inner one.
