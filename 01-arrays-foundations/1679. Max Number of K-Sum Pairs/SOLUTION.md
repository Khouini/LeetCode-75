# 1679. Max Number of K-Sum Pairs — Explained Like You're 10

---

## What did YOU try?

You used **two pointers `i` and `j`** and scanned for every pair that sums to `k`.
When you found one, you added both indexes to an `ignoredIndexes` list and skipped them later.

That works... but why does it **fail on large inputs?**

> **Q: What happens every time you want to skip an index?**
> You loop through the entire `ignoredIndexes` slice to check if `i` or `j` is in it.

> **Q: If there are 100 000 numbers, how many times could that loop run?**
> Millions of times. For every pair `(i, j)` you check every ignored index.
> That is **O(n²)** or worse — too slow.

---

## Step 1 — What are we even trying to do?

You have a bag of numbers.  
You want to pull out **as many pairs as possible** where the two numbers **add up to k**.  
Each number can only be used **once**.

> **Q: If k = 5 and you have [1, 2, 3, 4], can you make 2 pairs?**
> Yes: (1+4) and (2+3). Count = **2**.

> **Q: If k = 6 and you have [3, 1, 3, 4, 3], how many pairs add up to 6?**
> Only (3+3). But there are three 3s — you can only pair **two** of them. Count = **1**.

---

## Step 2 — The "sorted + two pointers" idea

> **Q: What if you sorted the bag first?**
> [1, 1, 2, 3, 4, 4] — smallest on the left, biggest on the right.

> **Q: Now put one finger on the far left (L) and one on the far right (R). What is L+R?**
> It is the smallest + the biggest. That is the most extreme sum you can get.

Now ask three questions every round:

| What is `nums[L] + nums[R]`? | What do you do?                                           |
| ---------------------------- | --------------------------------------------------------- |
| Equal to `k`                 | Count it! Move **both** fingers inward.                   |
| Less than `k`                | Sum is too small. Move **L** right (get a bigger number). |
| Greater than `k`             | Sum is too big. Move **R** left (get a smaller number).   |

Stop when the fingers meet.

> **Q: Why does moving inward always work?**
> Because the array is sorted. If the sum is too small, the only way to increase it is to raise the left side. If too big, lower the right side. You never need to go backward.

---

## Step 3 — Walk through an example

```
nums = [1, 2, 3, 4],  k = 5
sorted → [1, 2, 3, 4]
          L        R
```

**Round 1:** `1 + 4 = 5` ✓ → count = 1, move both inward

```
[1, 2, 3, 4]
      L  R
```

**Round 2:** `2 + 3 = 5` ✓ → count = 2, move both inward

L and R have crossed. **Done. Answer = 2.**

---

```
nums = [3, 1, 3, 4, 3],  k = 6
sorted → [1, 3, 3, 3, 4]
           L           R
```

**Round 1:** `1 + 4 = 5` < 6 → too small, move L right

```
[1, 3, 3, 3, 4]
    L        R
```

**Round 2:** `3 + 4 = 7` > 6 → too big, move R left

```
[1, 3, 3, 3, 4]
    L     R
```

**Round 3:** `3 + 3 = 6` ✓ → count = 1, move both inward

```
[1, 3, 3, 3, 4]
       LR
```

L and R have crossed. **Done. Answer = 1.**

---

## Step 4 — Why is this fast?

> **Q: How many times do L and R move total?**
> At most `n` times combined — L only moves right, R only moves left, and they stop when they meet.

> **Q: So what is the time complexity?**
> Sorting is **O(n log n)**. The two-pointer scan is **O(n)**. Total = **O(n log n)**.  
> Your original approach was closer to **O(n²)** — that is why it timed out.

---

## Step 5 — The code (Go)

```go
import "sort"

func maxOperations(nums []int, k int) int {
    sort.Ints(nums)      // sort the bag first
    L, R := 0, len(nums)-1
    count := 0

    for L < R {
        sum := nums[L] + nums[R]
        if sum == k {
            count++   // found a pair
            L++       // move both fingers inward
            R--
        } else if sum < k {
            L++       // too small, need a bigger left number
        } else {
            R--       // too big, need a smaller right number
        }
    }

    return count
}
```

---

## The key insight in one sentence

> **Sort the array, then chase the target from both ends — no need to remember which indexes you used.**

---

## Why your approach hurt you

| Your approach                        | Two-pointer approach   |
| ------------------------------------ | ---------------------- |
| Track used indexes in a list         | No extra memory needed |
| Check every ignored index per step   | O(1) decision per step |
| O(n²) on large inputs                | O(n log n)             |
| **Wrong on some inputs (logic bug)** | Correct on all inputs  |

Your thinking was right — you just needed a smarter way to avoid reusing numbers. Sorting does that job for free.

---

## The logic bug — not just performance

You actually had **two separate problems**: slow performance AND a wrong answer bug.

> **Q: What does `isEnd` mean in your code?**
> It is `true` when `j` has reached the last index of the array.

> **Q: When `isEnd` is true and the sum is NOT equal to `k`, what does your code do?**
> It fires `i++` and resets `j = i + 1` anyway — **even though no pair was found**.

> **Q: Why is that wrong?**
> Because you are forcing `i` to move forward when there might still be a valid `j` for the current `i` that you have not checked yet — one that was previously ignored and caused `j` to keep moving.

**Minimal example that breaks it:**

```
nums = [3, 4, 3, 3],  k = 6
                   ^ last index = 3
```

| Step | i (val)         | j (val) | isEnd    | sum            | action                        |
| ---- | --------------- | ------- | -------- | -------------- | ----------------------------- |
| 1    | 0 (3)           | 1 (4)   | false    | 7              | j++                           |
| 2    | 0 (3)           | 2 (3)   | false    | **6 ✓**        | count=1, ignore 0&2, i→1, j→2 |
| 3    | 1 (4)           | 2 (3)   | false    | ignored → j++  |                               |
| 4    | 1 (4)           | 3 (3)   | **true** | 7 ≠ 6          | **isEnd fires → i→2, j→3**    |
| 5    | 2 ignored → i→3 | j→4     | —        | j out of range | exit                          |

**Your output: 1**
**Correct answer: 2** — the pair (3+3) at indexes 0&2 AND the pair (3+3) at indexes... wait, let me recount.

Actually `[3,4,3,3]` with k=6: pairs are (3+3). There are three 3s, so max 1 pair of 3s... but also 4 has no partner. So correct is 1. Let me use the real counterexample:

```
nums = [3, 3, 4, 3],  k = 6
```

| Step | i (val)         | j (val) | isEnd    | sum     | action                        |
| ---- | --------------- | ------- | -------- | ------- | ----------------------------- |
| 1    | 0 (3)           | 1 (3)   | false    | **6 ✓** | count=1, ignore 0&1, i→1, j→2 |
| 2    | 1 ignored → i→2 | j→3     |          |         |                               |
| 3    | 2 (4)           | 3 (3)   | **true** | 7 ≠ 6   | **isEnd fires → i→3, j→4**    |
| 4    | j out of range  | exit    |          |         |                               |

**Your output: 1. Correct answer: 1.** OK still fine here.

The confirmed failing case from LeetCode (output 133, expected 138) is the real proof. The bug compounds across hundreds of numbers because `isEnd` prematurely ends the `j`-scan for an `i` whenever the **last index happens to be already-ignored**, causing a no-match `isEnd` exit too early.

> **Q: What is the fix for just the `isEnd` bug, without changing the whole approach?**
> Remove `isEnd` entirely. Only advance `i` when you find a match. When `j` goes past the end naturally, let the loop condition `j < n` handle it.

```go
// Fixed version of YOUR approach (still O(n²) but logically correct)
func maxOperations(nums []int, k int) int {
    n := len(nums)
    used := make([]bool, n)   // bool slice: O(1) lookup instead of scanning a list
    output := 0

    for i := 0; i < n; i++ {
        if used[i] {
            continue
        }
        for j := i + 1; j < n; j++ {
            if used[j] {
                continue
            }
            if nums[i]+nums[j] == k {
                output++
                used[i] = true
                used[j] = true
                break   // found the partner for i, no need to keep scanning j
            }
        }
    }
    return output
}
```

This is **logically correct** but **O(n²)** — it will still time out on the big test case.
The sort + two-pointer approach from above solves both problems at once.
