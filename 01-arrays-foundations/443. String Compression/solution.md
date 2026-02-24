# 443. String Compression – Two Pointers

---

## What's the Problem?

You have an array of characters like `["a","a","b","b","c"]`.
Compress it **in-place** to `["a","2","b","2","c"]` and return the new length.

Rules:

- A group of 1 character → write just the character
- A group of N > 1 → write the character then each digit of N

---

## Step-by-Step Thinking

### Question 1: What do we need to track?

Two things happen as we scan:

- We **read** characters from the original array
- We **write** compressed output back into the same array

👉 Use two pointers: `read` and `write`.

```
chars = [a, a, a, b, b, c, c, c]
         ↑               ↑
        read            write
```

---

### Question 2: How do we count a group?

When `read` lands on a new character, keep moving it forward as long as the character stays the same. Count how many steps you moved.

```
chars = [a, a, a, b, b, c, c, c]
         ↑
        read

chars[read] == 'a'? yes → count=1, read++
chars[read] == 'a'? yes → count=2, read++
chars[read] == 'a'? yes → count=3, read++
chars[read] == 'b'? no  → stop
```

Group: `'a'` × 3

---

### Question 3: What do we write?

After counting the group:

1. Write the character at `write`, then `write++`
2. If `count > 1`, convert count to string and write each digit

```
write the char:   [a, ?, ?, ?, ?, ?, ?, ?]
                   ↑
                  write → now write=1

count=3, write "3": [a, 3, ?, ?, ?, ?, ?, ?]
                        ↑
                       write → now write=2
```

---

### Question 4: Is it safe to overwrite?

Yes. The `write` pointer always stays **at or behind** `read`.
You never overwrite a character before you've already read it.

---

## Full Walkthrough

```
Input: [a, a, a, b, b, c, c, c]
```

| step | group | write sequence | array so far             |
| ---- | ----- | -------------- | ------------------------ |
| 1    | a × 3 | 'a', '3'       | [a, 3, ?, ?, ?, ?, ?, ?] |
| 2    | b × 2 | 'b', '2'       | [a, 3, b, 2, ?, ?, ?, ?] |
| 3    | c × 3 | 'c', '3'       | [a, 3, b, 2, c, 3, ?, ?] |

Return `write = 6`.

---

## Code

```go
func compress(chars []byte) int {
    write, read := 0, 0

    for read < len(chars) {
        curChar := chars[read]
        count := 0

        // count the full group
        for read < len(chars) && chars[read] == curChar {
            read++
            count++
        }

        // write the character
        chars[write] = curChar
        write++

        // write the count digits (only if count > 1)
        if count > 1 {
            for _, c := range strconv.Itoa(count) {
                chars[write] = byte(c)
                write++
            }
        }
    }

    return write
}
```

---

## Key Insight

> Count the whole group first, then write. Never write while still counting.

This keeps the logic clean and avoids the bookkeeping needed when updating digits on-the-fly.
