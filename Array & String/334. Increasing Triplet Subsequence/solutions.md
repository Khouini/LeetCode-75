# 334. Increasing Triplet Subsequence - Step by Step Explanation

Let me explain this like we're discovering the solution together, not memorizing anything.

---

## 🎯 What Are We Looking For?

We need **3 numbers** where:

- They appear in order (left to right): `i < j < k`
- Their values go up: `nums[i] < nums[j] < nums[k]`

**Visual:**

```
Array:  [2, 1, 5, 0, 4, 6]
         ↑        ↑  ↑
         1   <    4  < 6  ✅ Found!
```

---

## 🧠 Approach 1: Brute Force (The Obvious Way)

### Think Like This:

> "What if I just check EVERY possible combination of 3 numbers?"

### How?

1. Pick first number (position i)
2. Pick second number (position j, must be after i)
3. Pick third number (position k, must be after j)
4. Check if values go up

### Visual:

```
[2, 1, 5, 0, 4, 6]
 i  j  k           → 2 < 1? NO
 i  j     k        → 2 < 1? NO
 i  j        k     → 2 < 1? NO
 i     j  k        → 2 < 5? YES, 5 < 0? NO
 i     j     k     → 2 < 5? YES, 5 < 4? NO
 i     j        k  → 2 < 5? YES, 5 < 6? YES ✅ FOUND!
```

### Code:

```go
func increasingTriplet(nums []int) bool {
    n := len(nums)

    for i := 0; i < n-2; i++ {
        for j := i+1; j < n-1; j++ {
            for k := j+1; k < n; k++ {
                if nums[i] < nums[j] && nums[j] < nums[k] {
                    return true
                }
            }
        }
    }
    return false
}
```

### Problem?

- **Time: O(n³)** - Way too slow for big arrays!
- But it's a starting point. Now let's think smarter.

---

## 🧠 Approach 2: Think Backwards - What Do We Actually Need?

### Ask Yourself:

> "For each position j (the middle number), what do I need?"

**Answer:**

- I need to know: Is there ANY smaller number to my LEFT?
- I need to know: Is there ANY bigger number to my RIGHT?

### Visual:

```
Array: [2, 1, 5, 0, 4, 6]

For position 4 (value = 4):
  LEFT side: [2, 1, 5, 0] → smallest = 0 → 0 < 4? YES ✅
  RIGHT side: [6]         → biggest = 6  → 4 < 6? YES ✅

  So we found: 0 < 4 < 6 ✅
```

### Build Helper Arrays:

**Step 1:** For each position, what's the MINIMUM on the left?

```
Array:      [2, 1, 5, 0, 4, 6]
minLeft:    [2, 1, 1, 0, 0, 0]
             ↑  ↑  ↑  ↑  ↑  ↑
             2  1  1  0  0  0  (running minimum going right →)
```

**Step 2:** For each position, what's the MAXIMUM on the right?

```
Array:      [2, 1, 5, 0, 4, 6]
maxRight:   [6, 6, 6, 6, 6, 6]
             ↑  ↑  ↑  ↑  ↑  ↑
             6  6  6  6  6  6  (running maximum going left ←)
```

**Step 3:** Check each middle position:

```
Position 0: minLeft[0]=2, nums[0]=2, maxRight[0]=6 → 2 < 2? NO
Position 1: minLeft[1]=1, nums[1]=1, maxRight[1]=6 → 1 < 1? NO
Position 2: minLeft[2]=1, nums[2]=5, maxRight[2]=6 → 1 < 5 < 6? YES ✅
```

### Code:

```go
func increasingTriplet(nums []int) bool {
    n := len(nums)
    if n < 3 {
        return false
    }

    // Build minLeft array
    minLeft := make([]int, n)
    minLeft[0] = nums[0]
    for i := 1; i < n; i++ {
        minLeft[i] = min(minLeft[i-1], nums[i])
    }

    // Build maxRight array
    maxRight := make([]int, n)
    maxRight[n-1] = nums[n-1]
    for i := n-2; i >= 0; i-- {
        maxRight[i] = max(maxRight[i+1], nums[i])
    }

    // Check each position as middle
    for j := 1; j < n-1; j++ {
        if minLeft[j-1] < nums[j] && nums[j] < maxRight[j+1] {
            return true
        }
    }
    return false
}
```

### Complexity:

- **Time: O(n)** - Much better!
- **Space: O(n)** - We use extra arrays

---

## 🧠 Approach 3: The Clever Way (Optimal)

### Think Like This:

> "What if I walk through the array and keep track of just TWO things?"

**The key insight:**

- `first` = the **smallest** number I've seen so far (candidate for position i)
- `second` = the **smallest** number that's **bigger than first** (candidate for position j)

If I ever find a number **bigger than second**, I found my triplet!

### Walk Through Example:

```
Array: [2, 1, 5, 0, 4, 6]
Start: first = ∞, second = ∞

Step 1: num = 2
        2 < ∞? YES → first = 2
        first=2, second=∞

Step 2: num = 1
        1 < 2? YES → first = 1
        first=1, second=∞

Step 3: num = 5
        5 < 1? NO
        5 < ∞? YES → second = 5
        first=1, second=5

        Now we have: 1 < 5, just need something > 5

Step 4: num = 0
        0 < 1? YES → first = 0
        first=0, second=5

        ⚠️ Wait! first=0 came AFTER second=5 in array!
        That's OK! second=5 "remembers" there was a smaller number before it.

Step 5: num = 4
        4 < 0? NO
        4 < 5? YES → second = 4
        first=0, second=4

Step 6: num = 6
        6 < 0? NO
        6 < 4? NO
        6 > second? YES! ✅ FOUND TRIPLET!
```

### Why Does This Work?

The magic is: **`second` only gets a value if there was something smaller before it.**

So when we find `num > second`, we know:

- There exists some value < second (that's how second got its value)
- second < num
- Therefore: (something) < second < num ✅

### Code:

```go
func increasingTriplet(nums []int) bool {
    first := math.MaxInt32   // smallest so far
    second := math.MaxInt32  // smallest that's bigger than some previous number

    for _, num := range nums {
        if num <= first {
            first = num          // found new smallest
        } else if num <= second {
            second = num         // found new "middle" candidate
        } else {
            return true          // found third! (num > second > something)
        }
    }
    return false
}
```

### Complexity:

- **Time: O(n)** - One pass!
- **Space: O(1)** - Just two variables!

---

## 📊 Summary

| Approach      | Time  | Space | Idea                               |
| ------------- | ----- | ----- | ---------------------------------- |
| Brute Force   | O(n³) | O(1)  | Try all combinations               |
| Helper Arrays | O(n)  | O(n)  | Precompute min-left, max-right     |
| Two Variables | O(n)  | O(1)  | Track smallest and second-smallest |

---

## 🎯 The Problem-Solving Journey:

1. **Start dumb** → Check everything (brute force)
2. **Ask better questions** → What do I actually need at each position?
3. **Reduce memory** → Can I do this without storing everything?
