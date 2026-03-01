# 1679. Max Number of K-Sum Pairs — YouTube Video Script

---

## Intro (30 seconds)

**SAY:**

> "Given an array of integers and a target k, find the maximum number of pairs that sum to k. Each number can only be used once."

**WRITE on whiteboard:**

```
nums = [1, 2, 3, 4],  k = 5
→ (1,4) (2,3) → answer: 2

nums = [3, 1, 3, 4, 3],  k = 6
→ (3,3) → answer: 1
```

**SAY:**

> "Let's start with the obvious approach and see what it teaches us."

---

## Part 1 — Brute Force (2 minutes)

**WRITE on whiteboard:**

```
[1, 2, 3, 4]   k = 5

i=0: check 1+2, 1+3, 1+4 ✓ → pair (1,4)
i=1: check 2+3 ✓ → pair (2,3)
i=2: no partner left
```

**SAY:**

> "For every number, scan the rest of the array looking for its complement — the number that adds up to k."

**WRITE below:**

```
for i → 0 to n:
  for j → i+1 to n:
    if nums[i] + nums[j] == k → count++
```

**SAY:**

> "This works, but for each element we scan the entire remaining array. That's n times n — O(n²). For 100,000 numbers, that's 10 billion operations."

**CIRCLE the inner loop and WRITE next to it:**

```
← this is a SEARCH for (k - nums[i])
```

**SAY (key moment):**

> "Notice what's happening: the inner loop is just **searching** for a value. Whenever you're searching, ask yourself — can I make this search faster?"

**WRITE:**

```
Speed up a search:
  1. Hash map → O(1) lookup
  2. Sorting → eliminate chunks at once
```

**SAY:**

> "Two tools for this. Let's do sorting first — it reveals a beautiful pattern."

---

## Part 2 — Sort + Two Pointers (3 minutes)

### Step A — Sort and observe

**ERASE previous work. WRITE:**

```
[3, 1, 3, 4, 3]   k = 6

sorted → [1, 3, 3, 3, 4]
          ↑              ↑
       smallest       largest
```

**SAY:**

> "After sorting, smallest is on the left, largest on the right. Now watch what the sum of the two extremes tells us."

### Step B — The elimination logic

**WRITE three scenarios side by side:**

```
Case 1: sum == k       Case 2: sum < k        Case 3: sum > k
  ✓ pair found!         L is too small          R is too big
  move both inward      move L right →          ← move R left
```

**SAY (slow and clear):**

> "If sum equals k — great, we found a pair, move both inward."
>
> "If sum is too SMALL — L is already paired with the LARGEST number, and it still wasn't enough. So L can NEVER be in any valid pair. Throw it away. Move L right."
>
> "If sum is too BIG — R is already paired with the SMALLEST number, and it's still too much. So R can NEVER be in any valid pair. Throw it away. Move R left."

**SAY:**

> "This is the key insight — the sorted extremes tell you which end to eliminate."

### Step C — Walk-through

**WRITE and trace step by step (use arrows for L and R):**

```
[1, 3, 3, 3, 4]    k = 6
 L           R      1+4=5 < 6  → L right

[1, 3, 3, 3, 4]
    L        R      3+4=7 > 6  → R left

[1, 3, 3, 3, 4]
    L     R         3+3=6 == 6 → pair! count=1, both inward

[1, 3, 3, 3, 4]
       LR           L meets R → STOP
```

**SAY:**

> "Each step, one or both pointers move inward. They start n apart and meet in the middle. That's O(n) moves."

### Step D — Show the code

**WRITE (or show slide):**

```go
sort.Ints(nums)
L, R := 0, len(nums)-1
count := 0

for L < R {
    sum := nums[L] + nums[R]
    if sum == k {      count++; L++; R--
    } else if sum < k { L++
    } else {            R--
    }
}
```

**WRITE complexity:**

```
Time:  O(n log n)  ← sorting
Space: O(1)        ← just two variables
```

**SAY:**

> "Sorting costs n log n, the scan is n. No extra memory. Clean."

---

## Part 3 — Hash Map / Waiting Room (3 minutes)

**SAY:**

> "Now let's go back to the other idea — making the search O(1) with a hash map."

### Step A — The waiting room analogy

**DRAW a box on the whiteboard labeled "Waiting Room".**

**SAY:**

> "Imagine a waiting room at a train station. Each number walks in and says: I need someone who adds up to k with me. They check if their complement is already sitting there."

**WRITE the two outcomes next to the box:**

```
Number arrives:
  → complement in room?  YES → pair up, both leave
  → complement in room?  NO  → sit down and wait
```

### Step B — Walk-through

**WRITE the array and trace with the waiting room box:**

```
nums = [3, 1, 3, 4, 3]   k = 6
```

**For each step, WRITE the number arriving, update the waiting room box:**

```
3 arrives → need 3 → room: {}        → nobody → sits    room: {3:1}
1 arrives → need 5 → room: {3:1}     → nobody → sits    room: {3:1, 1:1}
3 arrives → need 3 → room: {3:1,1:1} → FOUND! → pair!   room: {3:0, 1:1}  count=1
4 arrives → need 2 → room: {3:0,1:1} → nobody → sits    room: {3:0, 1:1, 4:1}
3 arrives → need 3 → room: {3:0,...}  → nobody → sits    room: {3:1, 1:1, 4:1}
```

**SAY after the trace:**

> "Answer is 1. Each element does one O(1) map lookup. Single pass through the array."

### Step C — The special case that isn't special

**SAY:**

> "Notice 3 + 3 = 6 — a number pairing with itself. We didn't need any special case. The first 3 sits, the second 3 finds it and pairs. The waiting room handles it automatically."

### Step D — Show the code

**WRITE (or show slide):**

```go
freq := make(map[int]int)
count := 0

for _, num := range nums {
    complement := k - num
    if freq[complement] > 0 {
        count++
        freq[complement]--
    } else {
        freq[num]++
    }
}
```

**WRITE complexity:**

```
Time:  O(n)   ← single pass
Space: O(n)   ← the map
```

---

## Part 4 — Comparison & Wrap Up (1 minute)

**DRAW a comparison table:**

```
              Brute Force    Two Pointers    Hash Map
Time:         O(n²)          O(n log n)      O(n)
Space:        O(n)           O(1)            O(n)
Key idea:     scan all       sort+eliminate  remember seen
```

**SAY:**

> "Brute force searches for every complement — O(n²). Two pointers organizes the data by sorting, then eliminates from the ends — O(n log n) time, O(1) space. Hash map remembers what it's seen, making every search instant — O(n) time but O(n) space."

> "Both optimized approaches start from the same observation: the brute force bottleneck is searching for a complement. One speeds it up with order, the other with memory."

---

## Timing Summary

| Section                   | Duration      |
| ------------------------- | ------------- |
| Intro                     | ~30s          |
| Brute Force + observation | ~2 min        |
| Sort + Two Pointers       | ~3 min        |
| Hash Map / Waiting Room   | ~3 min        |
| Comparison & wrap up      | ~1 min        |
| **Total**                 | **~9-10 min** |
