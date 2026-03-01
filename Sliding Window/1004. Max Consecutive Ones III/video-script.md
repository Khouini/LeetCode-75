# 1004. Max Consecutive Ones III — YouTube Video Script

---

## Intro (30 seconds)

**SAY:**

> "Given a binary array — only 0s and 1s — and a number k, you're allowed to flip at most k zeros into ones. Find the longest streak of consecutive 1s you can get."

**WRITE on whiteboard:**

```
nums = [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0]   k = 2

flip these two ─────────┐  ┐
                         ▼  ▼
       [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0]
                      ▲  ▲
        flip these two┘  ┘

Best choice: flip index 5 and 10
→ [1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1]
                  ───────────────────
                  6 consecutive 1s → answer: 6
```

**SAY:**

> "The trick is choosing WHICH zeros to flip. Let's figure that out step by step."

---

## Part 1 — Reframe the Problem (1 minute)

**SAY:**

> "Before jumping into code, let's reframe the problem. Flipping zeros sounds complicated. But notice — we're looking for a stretch of the array that already has a lot of 1s, and only a FEW zeros — at most k zeros."

**WRITE:**

```
"Flip at most k zeros"

  is the SAME as

"Find the longest subarray with at most k zeros"
```

**SAY (slow and clear):**

> "Read that again. We don't actually flip anything. We just need to find the longest window in the array that contains at most k zeros. If a window has at most k zeros, we COULD flip them all, and the entire window becomes all 1s."

**WRITE example:**

```
[1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0]   k = 2

window [0, 0, 1, 1, 1, 1] → 2 zeros ≤ k → valid! length = 6
window [0, 0, 0, 1, 1, 1] → 3 zeros > k → invalid!
```

**SAY:**

> "This reframing is the first key insight. Now the problem is: find the longest subarray with at most k zeros. Much cleaner to think about."

---

## Part 2 — Brute Force (2 minutes)

**SAY:**

> "Let's try the obvious approach first. Check EVERY possible subarray, count its zeros, keep the longest valid one."

**WRITE:**

```
[1, 0, 0, 1, 1]   k = 1

Start at i=0:
  [1]             → 0 zeros ✓  length 1
  [1, 0]          → 1 zero  ✓  length 2
  [1, 0, 0]       → 2 zeros ✗  STOP
Start at i=1:
  [0]             → 1 zero  ✓  length 1
  [0, 1]          → 1 zero  ✓  length 2
  [0, 1, 1]       → 1 zero  ✓  length 3  ← best so far!
Start at i=2:
  ...and so on
```

**WRITE pseudocode:**

```
for i → 0 to n:
  zeros = 0
  for j → i to n:
    if nums[j] == 0: zeros++
    if zeros > k: break
    max = max(max, j - i + 1)
```

**SAY:**

> "This works, but it's O(n²). For every starting position, we extend right counting zeros. Can we do better?"

**SAY (key observation):**

> "Watch what happens when we move start from i=0 to i=1. We recount almost the same elements. We're throwing away information we already have. What if instead of restarting, we just ADJUSTED the window?"

---

## Part 3 — Sliding Window (4 minutes)

### Step A — The window idea

**DRAW a long array on the whiteboard:**

```
[1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0]   k = 2
```

**SAY:**

> "Imagine a stretchy window on the array. It has a left edge L and a right edge R. The rule is simple: the window is allowed to contain AT MOST k zeros."

**DRAW the window as a bracket:**

```
[1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0]
 L─────────────R
 window: [1,1,1,0,0]  → 2 zeros → valid ✓
```

### Step B — Two moves only

**WRITE:**

```
Two moves:

1. EXPAND  → move R right (grow the window)
   "Let's try to include one more element"

2. SHRINK  → move L right (shrink the window)
   "We have too many zeros, let go of the leftmost element"
```

**SAY:**

> "R always moves right, one step at a time — we're always trying to grow. But if the window has MORE than k zeros, it's invalid. So we shrink from the left until it's valid again."

### Step C — Why shrink from the left?

**SAY:**

> "This is the part that confused me at first. Why move L forward? Why not move R back?"

> "Think about it — R just added a new zero that broke our window. Everything to the LEFT of R was already explored. The only way to FIX the window is to push L forward, hoping to drop a zero off the left side."

**WRITE:**

```
Window too many zeros?
→ Move R back?  NO — we'd undo progress
→ Move L forward? YES — drop elements from the left
   until we drop a zero, restoring the balance
```

### Step D — Full walk-through

**WRITE on whiteboard:**

```
nums = [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0]   k = 2
```

**DRAW (or show slide) the table — fill it in row by row:**

Size = R - L + 1

| R         | nums[R] | zeros | L   | Size  | Max     | Note                                                |
| --------- | ------- | ----- | --- | ----- | ------- | --------------------------------------------------- |
| 0         | 1       | 0     | 0   | 1     | 1       | expand: nums[0]=1, window=[1]                       |
| 1         | 1       | 0     | 0   | 2     | 2       | expand: nums[1]=1, window=[1,1]                     |
| 2         | 1       | 0     | 0   | 3     | 3       | expand: nums[2]=1, window=[1,1,1]                   |
| 3         | 0       | 1     | 0   | 4     | 4       | expand: nums[3]=0, zeros=1 ≤ k → ok                 |
| 4         | 0       | 2     | 0   | 5     | 5       | expand: nums[4]=0, zeros=2 ≤ k → ok                 |
| 5         | 0       | **3** | 0   | —     | —       | expand: nums[5]=0, zeros=3 **> k → SHRINK!**        |
| ↳ shrink  |         | 3     | 1   | —     | —       | nums[L=0]=1, not a zero → L++ (no change to zeros)  |
| ↳ shrink  |         | 3     | 2   | —     | —       | nums[L=1]=1, not a zero → L++                       |
| ↳ shrink  |         | 3     | 3   | —     | —       | nums[L=2]=1, not a zero → L++                       |
| ↳ shrink  |         | **2** | 4   | —     | —       | nums[L=3]=**0** → zeros-- → zeros=2 ≤ k, L++ → stop |
| 5 (done)  |         | 2     | 4   | **2** | 5       | size = R-L+1 = 5-4+1 = 2, no new max                |
| 6         | 1       | 2     | 4   | 3     | 5       | expand: nums[6]=1, window=[0,0,0,1]                 |
| 7         | 1       | 2     | 4   | 4     | 5       | expand: nums[7]=1, window=[0,0,0,1,1]               |
| 8         | 1       | 2     | 4   | 5     | 5       | expand: nums[8]=1, window=[0,0,0,1,1,1]             |
| 9         | 1       | 2     | 4   | 6     | **6** ✓ | expand: nums[9]=1, size=6 → new best!               |
| 10        | 0       | **3** | 4   | —     | —       | expand: nums[10]=0, zeros=3 **> k → SHRINK!**       |
| ↳ shrink  |         | **2** | 5   | —     | —       | nums[L=4]=**0** → zeros-- → zeros=2 ≤ k, L++ → stop |
| 10 (done) |         | 2     | 5   | **6** | 6       | size = R-L+1 = 10-5+1 = 6, max stays 6              |

**SAY:**

> "Done. The answer is 6. The two rows where zeros goes '3 → 2' are the shrink steps — L jumps forward until it drops off a zero. Notice that L only moved forward 5 times total, and R moved forward 11 times. Each element is visited at most twice — once by R, once by L."

### Step E — Show the code

**WRITE:**

```go
func longestOnes(nums []int, k int) int {
    L := 0
    zerosCount := 0
    max := 0

    for R := 0; R < len(nums); R++ {
        if nums[R] == 0 {
            zerosCount++
        }

        // Shrink window until we have at most k zeros
        for zerosCount > k {
            if nums[L] == 0 {
                zerosCount--
            }
            L++
        }

        // Update best window size
        if R - L + 1 > max {
            max = R - L + 1
        }
    }

    return max
}
```

**WRITE complexity:**

```
Time:  O(n)  ← each pointer moves at most n times
Space: O(1)  ← just a few variables
```

---

## Part 4 — Why Is This O(n) and Not O(n²)? (1 minute)

**SAY:**

> "You might look at the nested loop — a for inside a for — and think it's O(n²). It's not. Here's why."

**WRITE:**

```
R moves right:  0 → 1 → 2 → ... → n     (n steps total)
L moves right:  0 → ... → at most n      (n steps total)

L NEVER moves left. It only goes forward.
Total work = n + n = O(n)
```

**SAY:**

> "Both pointers only move forward, and each can move at most n times. So the total number of operations is at most 2n, which is O(n). The inner while loop doesn't restart from scratch — it picks up where it left off."

---

## Part 5 — The Sliding Window Pattern (1 minute)

**SAY:**

> "This is the classic sliding window pattern. It shows up everywhere. Let me give you the template."

**WRITE:**

```
Sliding Window Template:

1. Expand R to include a new element
2. Update window state (count zeros, sum, etc.)
3. While window is INVALID → shrink from L
4. Update the answer with current window size

Works when:
  - You need the longest/shortest SUBARRAY
  - With some constraint (at most k zeros, sum ≤ target, etc.)
```

**SAY:**

> "Anytime you see 'longest subarray with at most...' or 'shortest subarray with at least...' — think sliding window."

---

## Wrap Up (30 seconds)

**WRITE summary:**

```
Key insights:

1. REFRAME: "flip k zeros" = "find longest subarray with ≤ k zeros"
2. Brute force: try every subarray → O(n²)
3. Sliding window: expand R, shrink L when invalid → O(n)
4. Both L and R only move forward → total work is O(n)
```

**SAY:**

> "The hardest part of this problem is the reframing — realizing you don't actually flip anything. Once you see it as 'find the longest window with at most k zeros,' the sliding window writes itself."

---

## Timing Summary

| Section                                       | Duration    |
| --------------------------------------------- | ----------- |
| Intro                                         | ~30s        |
| Reframe the problem                           | ~1 min      |
| Brute Force + observation                     | ~2 min      |
| Sliding Window (concept + walkthrough + code) | ~4 min      |
| Why O(n) not O(n²)                            | ~1 min      |
| Sliding Window pattern                        | ~1 min      |
| Wrap up                                       | ~30s        |
| **Total**                                     | **~10 min** |
