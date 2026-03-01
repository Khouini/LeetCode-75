# 283. Move Zeroes - Step by Step Explanation

---

## What's The Problem?

You have an array with some zeros mixed in. You need to:
1. Move all zeros to the **end**
2. Keep all non-zeros in their **original relative order**
3. Do it **in-place** (no extra array allowed!)

**Visual:**

```
Input:  [0, 1, 0, 3, 12]
         ↓     ↓
        zero  zero

Output: [1, 3, 12, 0, 0]
        ─────────  ────
        non-zeros  zeros
        (in order) (at end)
```

---

## Let's Think About This Step by Step

### Question 1: What does the final array look like?

Look at input and output:
```
Input:  [0, 1, 0, 3, 12]
Output: [1, 3, 12, 0, 0]
```

The output has **two parts**:
```
[1, 3, 12 | 0, 0]
 ────────   ────
   PART 1   PART 2
 non-zeros  zeros
```

**So our job is:** Split the array into non-zeros (left) and zeros (right).

---

### Question 2: What's the simplest way to do this?

> "If I had unlimited memory, what would I do?"

Easy! Make a new array:
1. First, add all non-zeros: `[1, 3, 12]`
2. Then, add zeros to fill: `[1, 3, 12, 0, 0]`
3. Copy back to original array

```go
func moveZeroes(nums []int) {
    temp := make([]int, 0, len(nums))
    
    // Step 1: Collect non-zeros
    for _, num := range nums {
        if num != 0 {
            temp = append(temp, num)
        }
    }
    
    // Step 2: Fill rest with zeros
    for len(temp) < len(nums) {
        temp = append(temp, 0)
    }
    
    // Step 3: Copy back
    copy(nums, temp)
}
```

**But wait!** The problem says "in-place" = no extra array allowed!

---

### Question 3: How do I do this WITHOUT an extra array?

> "I can't store things elsewhere... so I need to rearrange elements IN the same array."

Think about it:
- I need to **move non-zeros to the front**
- Zeros will naturally end up at the back

**But how do I know WHERE to put each non-zero?**

---

### Question 4: Where should each non-zero go?

Let's think about this manually:

```
[0, 1, 0, 3, 12]
```

- The first non-zero I find should go to position 0
- The second non-zero I find should go to position 1
- The third non-zero I find should go to position 2
- ... and so on

**I need to remember: "Where should the NEXT non-zero go?"**

Let's call this `insertPos` (the position where I'll insert the next non-zero).

```
insertPos = 0    means "next non-zero goes at index 0"
insertPos = 1    means "next non-zero goes at index 1"
insertPos = 2    means "next non-zero goes at index 2"
```

---

### Question 5: What is `insertPos` exactly?

Think of `insertPos` as a **bookmark** or **placeholder**.

It answers: **"If I find a non-zero right now, where do I put it?"**

```
[_, _, _, _, _]
 ↑
insertPos = 0

"The next non-zero I find goes HERE"
```

---

### Question 6: When do I move `insertPos`?

**Only when I successfully place a non-zero!**

- Found a zero? Don't move `insertPos` (we're not placing anything)
- Found a non-zero? Place it at `insertPos`, THEN move `insertPos` to the next spot

```
Before: insertPos = 0
        I find non-zero 1, put it at position 0
After:  insertPos = 1  (ready for the next non-zero)
```

---

### Question 7: What do I do when I see a zero?

**Nothing! Just skip it.**

Why? Because we only care about placing non-zeros. Zeros will be handled later.

```
See a zero → "Not interested, move on"
See a non-zero → "Ah! Put this at insertPos, then insertPos++"
```

---

## Let's Walk Through It VERY Slowly

```
Array: [0, 1, 0, 3, 12]
Index:  0  1  2  3  4

insertPos = 0  (first non-zero will go to index 0)
```

**Step 1: Look at index 0**
```
[0, 1, 0, 3, 12]
 ↑
 i=0, value is 0

Q: Is it zero? YES
A: Skip it! Don't touch insertPos.

insertPos still = 0
```

**Step 2: Look at index 1**
```
[0, 1, 0, 3, 12]
    ↑
    i=1, value is 1

Q: Is it zero? NO, it's a non-zero!
A: Put it at insertPos (which is 0)!

nums[insertPos] = nums[i]
nums[0] = 1

Array becomes: [1, 1, 0, 3, 12]
                ↑
            we wrote here

Now move insertPos: insertPos = 1
```

**Step 3: Look at index 2**
```
[1, 1, 0, 3, 12]
       ↑
       i=2, value is 0

Q: Is it zero? YES
A: Skip it!

insertPos still = 1
```

**Step 4: Look at index 3**
```
[1, 1, 0, 3, 12]
          ↑
          i=3, value is 3

Q: Is it zero? NO!
A: Put it at insertPos (which is 1)!

nums[1] = 3

Array becomes: [1, 3, 0, 3, 12]
                   ↑
               we wrote here

insertPos = 2
```

**Step 5: Look at index 4**
```
[1, 3, 0, 3, 12]
              ↑
              i=4, value is 12

Q: Is it zero? NO!
A: Put it at insertPos (which is 2)!

nums[2] = 12

Array becomes: [1, 3, 12, 3, 12]
                      ↑
                  we wrote here

insertPos = 3
```

**Step 6: Done scanning! But wait...**

```
Array is: [1, 3, 12, 3, 12]
                    ↑
               insertPos = 3

The part from insertPos to end still has old junk!
We need to fill it with zeros.
```

**Step 7: Fill the rest with zeros**
```
[1, 3, 12, 3, 12]
          ↑  ↑
          3  4   ← these positions need to be 0

[1, 3, 12, 0, 0]  ← DONE!
```

---

## The Code (Two-Pass Version)

```go
func moveZeroes(nums []int) {
    insertPos := 0
    
    // Pass 1: Move all non-zeros to the front
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {          // Found a non-zero?
            nums[insertPos] = nums[i]  // Put it at insertPos
            insertPos++                 // Move insertPos forward
        }
        // If it's zero, we do nothing (skip)
    }
    
    // Pass 2: Fill the rest with zeros
    for i := insertPos; i < len(nums); i++ {
        nums[i] = 0
    }
}
```

---

## Question 8: Can we do this in ONE pass?

> "I'm doing two loops. Can I avoid the second loop that fills zeros?"

**The problem:** After Pass 1, there's leftover junk at the end.

**The idea:** What if instead of COPYING, we SWAP?

When we swap:
- The non-zero moves to `insertPos` (good!)
- Whatever was at `insertPos` moves to where we are (it's probably a zero!)

---

## Let's See Swapping in Action

```
Array: [0, 1, 0, 3, 12]
insertPos = 0
```

**Step 1: i=0, value=0**
```
Zero → skip

[0, 1, 0, 3, 12]
 ↑
insertPos=0
```

**Step 2: i=1, value=1**
```
Non-zero! SWAP nums[insertPos] with nums[i]
SWAP nums[0] with nums[1]

Before: [0, 1, 0, 3, 12]
         ↑  ↑
      swap these

After:  [1, 0, 0, 3, 12]
            ↑
       insertPos=1
```

**Step 3: i=2, value=0**
```
Zero → skip

[1, 0, 0, 3, 12]
    ↑
insertPos=1
```

**Step 4: i=3, value=3**
```
Non-zero! SWAP nums[1] with nums[3]

Before: [1, 0, 0, 3, 12]
            ↑     ↑
         swap these

After:  [1, 3, 0, 0, 12]
               ↑
          insertPos=2
```

**Step 5: i=4, value=12**
```
Non-zero! SWAP nums[2] with nums[4]

Before: [1, 3, 0, 0, 12]
               ↑      ↑
            swap these

After:  [1, 3, 12, 0, 0]
                  ↑
             insertPos=3

DONE! No need for second pass!
```

---

## Why Does Swapping Work?

Think about what's at `insertPos`:
- Before any swaps: it's whatever was there originally
- After swaps start: it's always a ZERO (because we swapped a zero there earlier!)

```
[1, 0, 0, 3, 12]
    ↑
 insertPos points to a zero!

When we swap nums[insertPos] with nums[3]:
- 3 goes to insertPos (good!)
- 0 goes to position 3 (zeros moving right!)
```

The zeros naturally "bubble" to the right side!

---

## The Code (One-Pass with Swap)

```go
func moveZeroes(nums []int) {
    insertPos := 0
    
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            // Swap: put non-zero at insertPos, 
            //       whatever was at insertPos goes to i
            nums[i], nums[insertPos] = nums[insertPos], nums[i]
            insertPos++
        }
    }
}
```

---

## Question 9: What if there are no zeros?

```
Array: [1, 2, 3]
```

- i=0: non-zero, swap nums[0] with nums[0] (swap with itself = no change)
- i=1: non-zero, swap nums[1] with nums[1] (no change)
- i=2: non-zero, swap nums[2] with nums[2] (no change)

Result: `[1, 2, 3]` - works fine!

---

## Question 10: What if ALL are zeros?

```
Array: [0, 0, 0]
```

- i=0: zero, skip
- i=1: zero, skip
- i=2: zero, skip

insertPos never moved! It stayed at 0.

Result: `[0, 0, 0]` - works fine!

---

## Summary: The Mental Model

Think of TWO characters:

**READER (i):** Walks through the whole array, looking at each element.

**WRITER (insertPos):** Stays put until READER finds something good (non-zero). Then WRITER writes it and moves forward.

```
        READER
           ↓
[0, 1, 0, 3, 12]
 ↑
WRITER

READER: "Found 0, meh."         → READER moves, WRITER stays
READER: "Found 1, that's good!" → Swap to WRITER position, both move
READER: "Found 0, meh."         → READER moves, WRITER stays
READER: "Found 3, that's good!" → Swap to WRITER position, both move
READER: "Found 12, nice!"       → Swap to WRITER position, both move
```

---

## Complexity

| Approach         | Time | Space | Notes                    |
| ---------------- | ---- | ----- | ------------------------ |
| Extra Array      | O(n) | O(n)  | Not in-place             |
| Two Pointers     | O(n) | O(1)  | Copy then fill zeros     |
| Swap (Optimal)   | O(n) | O(1)  | Single pass              |

---

## Pattern Recognition

This is a **Two Pointer / Partition** problem:

- One pointer (READER) scans everything
- One pointer (WRITER) tracks where to place "good" elements
- We're partitioning: `[good stuff | bad stuff]`

Similar problems: 
- Remove Element (LeetCode 27)
- Remove Duplicates from Sorted Array (LeetCode 26)

---
