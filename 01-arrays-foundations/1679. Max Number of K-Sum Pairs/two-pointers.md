# 1679. Max Number of K-Sum Pairs — Two Pointers

---

## The Problem

You have an array of integers and a target `k`.
Remove **as many pairs as possible** where the two numbers add up to `k`.
Each number can only be used once. Return the count of pairs.

---

## Start from Brute Force

The first instinct: for every number, scan the rest of the array looking for a partner.

```go
for i := 0; i < n; i++ {
    for j := i + 1; j < n; j++ {
        if nums[i] + nums[j] == k → count it, mark both as used
    }
}
```

This works. But:

> **Q: What is the expensive operation here?**
> For each `i`, we scan through all remaining `j` values to find the complement `k - nums[i]`.

> **Q: If the array has 100 000 elements, how bad is this?**
> For each of the 100 000 elements, we scan up to 100 000 others. That's **O(n²)** — too slow.

> **Q: What is the actual work we're doing in that inner loop?**
> We're **searching** for a specific value: `k - nums[i]`.

Whenever you find yourself searching, ask:

> **"Can I organize the data so the search becomes faster or disappears entirely?"**

---

## Observation 1 — What Does Sorting Give Us?

> **Q: What happens if we sort the array first?**
> The smallest numbers are on the left, the largest on the right.

Take `nums = [1, 2, 3, 4]`, `k = 5`. After sorting: `[1, 2, 3, 4]`.

> **Q: What is the smallest possible sum we can make?**
> `1 + 2` — the two leftmost numbers.

> **Q: What is the largest possible sum?**
> `3 + 4` — the two rightmost numbers.

> **Q: What does `nums[left] + nums[right]` give us?**
> The sum of the current **extremes** — the smallest and the largest remaining numbers.

---

## Observation 2 — The Extremes Tell Us What to Eliminate

Let's say `L` points to the leftmost element and `R` to the rightmost.

> **Q: If `nums[L] + nums[R] > k`, what does that tell us about `R`?**
> The sum is too big. Since `L` is already the **smallest** possible left value, no matter what we put on the left, `nums[R]` will always overshoot. So **`R` can never be in a valid pair** — we can safely discard it.

> **Q: If `nums[L] + nums[R] < k`, what does that tell us about `L`?**
> The sum is too small. Since `R` is already the **largest** possible right value, no matter what we put on the right, `nums[L]` will always undershoot. So **`L` can never be in a valid pair** — we can safely discard it.

> **Q: If `nums[L] + nums[R] == k`?**
> We found a pair! Count it and discard both — move both pointers inward.

This gives us a rule for every step:

| `nums[L] + nums[R]` | Conclusion                                 | Action                         |
| ------------------- | ------------------------------------------ | ------------------------------ |
| `== k`              | Valid pair found                           | Count it, move **both** inward |
| `> k`               | R is too large, can never pair with anyone | Move `R` left                  |
| `< k`               | L is too small, can never pair with anyone | Move `L` right                 |

---

## Observation 3 — This Is O(n)

> **Q: How many times do L and R move total?**
> L only moves right, R only moves left. They start `n` positions apart and stop when they meet. So at most **n moves total**.

> **Q: So what's the total time complexity?**
> Sorting: **O(n log n)**. The two-pointer scan: **O(n)**. Total: **O(n log n)**.

Compare to brute force O(n²): for n = 100 000, that's the difference between ~1.7 million operations and ~10 billion.

---

## Walk-through

```
nums = [3, 1, 3, 4, 3],  k = 6
sorted → [1, 3, 3, 3, 4]
          L           R
```

**Step 1:** `1 + 4 = 5 < 6` → too small → `L` can never pair → move L right

```
[1, 3, 3, 3, 4]
    L        R
```

**Step 2:** `3 + 4 = 7 > 6` → too big → `R` can never pair → move R left

```
[1, 3, 3, 3, 4]
    L     R
```

**Step 3:** `3 + 3 = 6 == k` → pair found! count = 1 → move both inward

```
[1, 3, 3, 3, 4]
       LR
```

L and R have met → **stop**. Answer = **1**.

---

## Code

```go
func maxOperations(nums []int, k int) int {
    sort.Ints(nums)          // step 1: impose order
    L, R := 0, len(nums)-1  // step 2: start at both extremes
    count := 0

    for L < R {
        sum := nums[L] + nums[R]
        if sum == k {
            count++   // valid pair — consume both ends
            L++
            R--
        } else if sum < k {
            L++       // left is too small — push it right
        } else {
            R--       // right is too large — push it left
        }
    }

    return count
}
```

---

## Complexity

|           |                                         |
| --------- | --------------------------------------- |
| **Time**  | O(n log n) — sorting dominates          |
| **Space** | O(1) extra — just two pointer variables |

---

## The Thinking Chain in One View

```
Brute force is O(n²)
  ↓ why?
  Inner loop is searching for a complement
  ↓ question:
  Can I organize data to remove that search?
  ↓ idea:
  Sort — gives order to exploit
  ↓ observation:
  Sum of extremes reveals which extreme is useless
  ↓ leads to:
  Two pointers — one step eliminates one element
  ↓ result:
  O(n) scan after O(n log n) sort
```
