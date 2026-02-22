# 1679. Max Number of K-Sum Pairs — Hash Map (Waiting Room)

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

> **Q: What is the expensive operation here?**
> For each `i`, we scan through all remaining `j` values to find the complement `k - nums[i]`. That's O(n) per element → O(n²) total.

> **Q: What are we actually doing in that inner loop?**
> We're **searching** for a specific value: `k - nums[i]`.

Whenever you find yourself searching, ask:

> **"Can I make that search O(1) instead of O(n)?"**

---

## Observation 1 — A Hash Map Gives O(1) Lookup

> **Q: What data structure gives you O(1) lookup by value?**
> A **hash map**.

> **Q: What would we store in it?**
> Every number we've seen that hasn't been paired yet.

> **Q: Then how would pairing work?**
> When we see a new number, we check: _"is my complement already in the map?"_
>
> - Yes → pair found. Remove the complement from the map.
> - No → store ourselves in the map and keep going.

That's the entire algorithm. One pass, O(1) per step.

---

## Observation 2 — The "Waiting Room" Mental Model

Think of the map as a **waiting room** at a train station.

Each number that arrives says: _"I'm looking for someone who adds up to `k` with me."_

- They check the waiting room: **is my complement already sitting there?**
  - **Yes** → pair up, count it, both leave.
  - **No** → sit down in the waiting room and wait for a future partner.

> **Q: Can a number be used twice?**
> No — when a pair is found, we **remove** one copy of the complement from the map. The count goes down.

> **Q: What if there are multiple copies of the same number?**
> The map tracks **counts**, not just presence. So `{3: 2}` means two 3s are waiting. When a third 3 arrives looking for complement 3, it finds `{3: 2} > 0`, pairs with one, and decrements to `{3: 1}`.

---

## Walk-through

`nums = [3, 1, 3, 4, 3]`, `k = 6`

| step | num | complement (k - num) | waiting room before | match?  | action                       | count |
| ---- | --- | -------------------- | ------------------- | ------- | ---------------------------- | ----- |
| 1    | 3   | 3                    | `{}`                | no      | sit → `{3:1}`                | 0     |
| 2    | 1   | 5                    | `{3:1}`             | no      | sit → `{3:1, 1:1}`           | 0     |
| 3    | 3   | 3                    | `{3:1, 1:1}`        | **yes** | pair! `{3:1}` → `{3:0, 1:1}` | 1     |
| 4    | 4   | 2                    | `{3:0, 1:1}`        | no      | sit → `{3:0, 1:1, 4:1}`      | 0     |
| 5    | 3   | 3                    | `{3:0, 1:1, 4:1}`   | no      | sit → `{3:1, 1:1, 4:1}`      | 0     |

**Answer: 1** ✅

---

## Code

```go
func maxOperations(nums []int, k int) int {
    freq := make(map[int]int) // the waiting room: number → count of unpaired copies
    count := 0

    for _, num := range nums {
        complement := k - num

        if freq[complement] > 0 {
            // complement is waiting — pair up and both leave
            count++
            freq[complement]--
        } else {
            // no partner yet — sit and wait
            freq[num]++
        }
    }

    return count
}
```

---

## Why No Special Cases?

> **Q: What if `num == complement`? (e.g. `3 + 3 = 6`)**
> It's handled automatically. The first 3 arrives, finds `freq[3] == 0`, sits. The second 3 arrives, finds `freq[3] == 1`, pairs with it, decrements to 0. The third 3 arrives, finds `freq[3] == 0`, sits again.
> No special logic needed — the waiting room handles it naturally.

---

## Complexity

|           |                                                        |
| --------- | ------------------------------------------------------ |
| **Time**  | O(n) — single pass, one O(1) map operation per element |
| **Space** | O(n) — the map can hold up to n unpaired numbers       |

---

## Comparison with Two Pointers

|                    | Hash Map (this file)      | Sort + Two Pointers             |
| ------------------ | ------------------------- | ------------------------------- |
| **Time**           | O(n)                      | O(n log n)                      |
| **Space**          | O(n) for the map          | O(1) extra                      |
| **Needs sorting?** | No                        | Yes                             |
| **Core idea**      | Remember what you've seen | Exploit order to eliminate ends |

The hash map is faster in time, but uses more memory. The two-pointer approach uses no extra memory, but requires sorting first.

---

## The Thinking Chain in One View

```
Brute force is O(n²)
  ↓ why?
  Searching for complement is O(n) per element
  ↓ question:
  Can I make that search O(1)?
  ↓ answer:
  Yes — store seen numbers in a hash map
  ↓ observation:
  Map acts as a waiting room — each number either
  finds a partner or waits for a future one
  ↓ result:
  Single pass → O(n) total
```
