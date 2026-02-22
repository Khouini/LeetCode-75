# 1679. Max Number of K-Sum Pairs

You have an array of integers and a target `k`.
Remove **as many pairs as possible** where the two numbers add up to `k`.
Each number can only be used once. Return the count of pairs.

---

## The Brute Force — and What It Teaches Us

The first instinct: for every number, scan the rest of the array looking for a partner that completes the sum to `k`.

```
for each i:
    for each j > i:
        if nums[i] + nums[j] == k → count it, mark both as used
```

```go
func maxOperations(nums []int, k int) int {
    n := len(nums)
    used := make([]bool, n) // track which elements have already been paired
    count := 0

    for i := 0; i < n; i++ {
        if used[i] {
            continue // i was already paired in a previous round, skip it
        }
        for j := i + 1; j < n; j++ {
            if used[j] {
                continue // j was already paired, skip it
            }
            if nums[i]+nums[j] == k {
                count++
                used[i] = true // mark both as consumed
                used[j] = true
                break // found the one partner for i, no need to keep scanning j
            }
        }
        // if the inner loop finished without a break, no partner exists for i
    }

    return count
}
```

This works, but it's **O(n²)**. For every element, you potentially scan the entire remaining array.

But notice something important about what we're doing here:

> We're searching for a **specific value** — the complement `k - nums[i]`.

Whenever you find yourself **searching** for something, ask:

> **"Can I organize the data so the search becomes faster?"**

---

## Observation 1 — Searching Is the Bottleneck

In brute force, the expensive part is: _"given `nums[i]`, where is `k - nums[i]`?"_

Two classic ways to speed up a search:

1. **Hash map** — look up the complement in O(1)
2. **Sorting** — use structure to eliminate large chunks of the search space at once

Both work here. Let's follow the sorting path, because it reveals a powerful pattern.

---

## Observation 2 — What Does Sorting Give Us?

Once sorted, the array has **order**. And order gives us **information about sums**.

Take `nums = [1, 2, 3, 4]`, `k = 5`.

After sorting: `[1, 2, 3, 4]`

> **Key insight:** the **smallest** number is on the left, the **largest** is on the right.

If we add the smallest + largest:

- `1 + 4 = 5` → that's exactly `k`.

What if the sum was **too big**? Say `k = 4` and we try `1 + 4 = 5 > 4`.

> The largest number is too large. **No number** paired with 4 will give a smaller sum than `1 + 4`, because 1 is already the smallest. So 4 **can never be part of any valid pair**.

What if the sum was **too small**? Say `k = 6` and we try `1 + 4 = 5 < 6`.

> The smallest number is too small. **No number** paired with 1 will give a larger sum than `1 + 4`, because 4 is already the largest. So 1 **can never be part of any valid pair**.

This is the critical reasoning:

> **In a sorted array, the sum of the two extremes tells you which extreme to eliminate.**

---

## Observation 3 — This Leads Naturally to Two Pointers

Since we can always **eliminate one end**, we place a pointer at each end:

- `L` at the beginning (smallest)
- `R` at the end (largest)

Each step, exactly one of three things happens:

| `nums[L] + nums[R]` | What it means                           | Action                         |
| ------------------- | --------------------------------------- | ------------------------------ |
| `== k`              | Found a pair                            | Count it, move **both** inward |
| `< k`               | Sum too small, L can't pair with anyone | Move `L` right                 |
| `> k`               | Sum too big, R can't pair with anyone   | Move `R` left                  |

Every step, at least one pointer moves inward. They meet in the middle. **Done.**

---

## Walk-through

```
nums = [3, 1, 3, 4, 3], k = 6
sorted → [1, 3, 3, 3, 4]
          L           R
```

**Step 1:** `1 + 4 = 5 < 6` → too small → move L right

```
[1, 3, 3, 3, 4]
    L        R
```

**Step 2:** `3 + 4 = 7 > 6` → too big → move R left

```
[1, 3, 3, 3, 4]
    L     R
```

**Step 3:** `3 + 3 = 6 == k` → pair found! count = 1 → move both inward

```
[1, 3, 3, 3, 4]
       LR
```

L meets R → **stop**. Answer = **1**.

---

## Code (Go)

```go
func maxOperations(nums []int, k int) int {
    sort.Ints(nums)
    L, R := 0, len(nums)-1
    count := 0

    for L < R {
        sum := nums[L] + nums[R]
        if sum == k {
            count++
            L++
            R--
        } else if sum < k {
            L++
        } else {
            R--
        }
    }

    return count
}
```

---

## Complexity

|           |                                                              |
| --------- | ------------------------------------------------------------ |
| **Time**  | O(n log n) — sorting dominates; the two-pointer scan is O(n) |
| **Space** | O(1) extra — sorting is in-place, only two pointer variables |

---

## The Thinking Chain — Summary

1. Brute force works but is O(n²) because we **search** for a complement every time.
2. **"Can I organize the data to search faster?"** → Yes: sort it.
3. Sorting gives order. Order means the sum of the two extremes tells us **which end is useless**.
4. That elimination logic is exactly the **two-pointer** pattern: L and R walk inward, each step discards one element or finds a pair.
5. Each pointer moves at most n times total → O(n) scan on top of O(n log n) sort.
