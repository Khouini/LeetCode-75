# 11. Container With Most Water - Step by Step Explanation

---

## 🎯 What's The Problem?

You have vertical lines (bars) at different positions. Each line has a height.
You need to find **two lines** that together with the x-axis form a container that holds the **most water**.

**Visual:**

```
Height: [1, 8, 6, 2, 5, 4, 8, 3, 7]

     8 |    █              █
     7 |    █              █     █
     6 |    █  █           █     █
     5 |    █  █     █     █     █
     4 |    █  █     █  █  █     █
     3 |    █  █     █  █  █  █  █
     2 |    █  █  █  █  █  █  █  █
     1 | █  █  █  █  █  █  █  █  █
       +---------------------------
         0  1  2  3  4  5  6  7  8
```

**The container between index 1 (height 8) and index 8 (height 7):**

```
     8 |    █ ≈ ≈ ≈ ≈ ≈ ≈ ≈|█
     7 |    █ ≈ ≈ ≈ ≈ ≈ ≈ ≈|█ ≈ ≈ █
     6 |    █ ≈ █ ≈ ≈ ≈ ≈ ≈|█ ≈ ≈ █
     5 |    █ ≈ █ ≈ ≈ █ ≈ ≈|█ ≈ ≈ █
     ...
         Water area = width × height
                    = (8-1) × min(8,7)
                    = 7 × 7 = 49
```

---

## 🔑 Key Insight: How Do We Calculate Water?

```
Water = width × height
      = (right_index - left_index) × min(left_height, right_height)
```

**Why min()?** Water spills over the shorter line!

```
   █
   █  ~~~~█    ← Water level = shorter bar
   █  ~~~~█
   █  ~~~~█
   L     R
```

---

## 🧠 Approach 1: Brute Force (Check Every Pair)

### Think Like This:

> "What if I just try EVERY possible pair of lines?"

### How?

1. Pick left line (position i)
2. Pick right line (position j, must be after i)
3. Calculate water between them
4. Keep track of maximum

### Visual:

```
[1, 8, 6, 2, 5, 4, 8, 3, 7]
 i  j                       → width=1, height=min(1,8)=1, area=1
 i     j                    → width=2, height=min(1,6)=1, area=2
 i        j                 → width=3, height=min(1,2)=1, area=3
 ...
    i                    j  → width=7, height=min(8,7)=7, area=49 ✅ MAX!
```

### Code:

```go
func maxArea(height []int) int {
    maxWater := 0
    n := len(height)

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            width := j - i
            h := min(height[i], height[j])
            water := width * h
            maxWater = max(maxWater, water)
        }
    }
    return maxWater
}
```

### Problem?

- **Time: O(n²)** - Too slow for big arrays!
- But it helps us understand the problem. Now let's think smarter.

---

## 🧠 Approach 2: Two Pointers (The Smart Way)

### Ask Yourself:

> "Do I really need to check EVERY pair?"

### The Key Observation:

Start with the **widest** container (left pointer at start, right pointer at end).

```
[1, 8, 6, 2, 5, 4, 8, 3, 7]
 L                       R

Width is MAXIMUM here (8).
Area = 8 × min(1, 7) = 8 × 1 = 8
```

Now, **which pointer should we move?**

### Think About It:

- If we move ANY pointer inward, **width decreases**
- To get MORE water, we need **height to increase**
- The height is limited by the **SHORTER** line
- So... **move the shorter one!** (hoping to find a taller line)

### Why Move The Shorter One?

```
Case: Left is shorter
    █              █
    █  ← shorter   █ ← taller
    L              R

If we move R inward:
- Width gets smaller ❌
- Height stays same (still limited by L) ❌
- Area MUST decrease!

If we move L inward:
- Width gets smaller ❌
- Height MIGHT increase ✅ (if we find taller line)
- Area MIGHT increase!
```

**Moving the taller one can NEVER help. Moving the shorter one MIGHT help.**

---

### Walk Through Example:

```
Array: [1, 8, 6, 2, 5, 4, 8, 3, 7]
        0  1  2  3  4  5  6  7  8

Step 1: L=0, R=8
        height[L]=1, height[R]=7
        Area = 8 × min(1,7) = 8
        1 < 7, so move L →
        maxArea = 8

Step 2: L=1, R=8
        height[L]=8, height[R]=7
        Area = 7 × min(8,7) = 49
        7 < 8, so move R ←
        maxArea = 49 ✅

Step 3: L=1, R=7
        height[L]=8, height[R]=3
        Area = 6 × min(8,3) = 18
        3 < 8, so move R ←
        maxArea = 49

Step 4: L=1, R=6
        height[L]=8, height[R]=8
        Area = 5 × min(8,8) = 40
        Equal! Move either (let's move L →)
        maxArea = 49

Step 5: L=2, R=6
        height[L]=6, height[R]=8
        Area = 4 × min(6,8) = 24
        6 < 8, so move L →
        maxArea = 49

Step 6: L=3, R=6
        height[L]=2, height[R]=8
        Area = 3 × min(2,8) = 6
        2 < 8, so move L →
        maxArea = 49

Step 7: L=4, R=6
        height[L]=5, height[R]=8
        Area = 2 × min(5,8) = 10
        5 < 8, so move L →
        maxArea = 49

Step 8: L=5, R=6
        height[L]=4, height[R]=8
        Area = 1 × min(4,8) = 4
        4 < 8, so move L →
        maxArea = 49

Step 9: L=6, R=6 → L meets R, STOP!

Answer: 49
```

---

### Visual Summary:

```
[1, 8, 6, 2, 5, 4, 8, 3, 7]
 L                       R   Area=8,  move L (1<7)
    L                    R   Area=49, move R (7<8) ⭐ MAX
    L                 R      Area=18, move R (3<8)
    L              R         Area=40, move L (8=8)
       L           R         Area=24, move L (6<8)
          L        R         Area=6,  move L (2<8)
             L     R         Area=10, move L (5<8)
                L  R         Area=4,  move L (4<8)
                L=R          DONE!
```

---

### Code:

```go
func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxWater := 0

    for left < right {
        // Calculate current area
        width := right - left
        h := min(height[left], height[right])
        water := width * h
        maxWater = max(maxWater, water)

        // Move the shorter pointer
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxWater
}
```

### Complexity:

- **Time: O(n)** - One pass through the array!
- **Space: O(1)** - Just two pointers!

---

## 🤔 Why Does Two Pointers Work? (Proof Intuition)

### The Worry:

> "What if we skip a good pair by moving a pointer?"

### The Answer:

When we move the shorter pointer, we're **NOT skipping any useful pairs**.

**Example:** If `height[L] < height[R]`, and we move L:

- Any pair with current L and some R' (where R' < R) would have:
  - Smaller width (R' < R)
  - Height still limited by L (the short one)
  - So area would be **smaller**!

We already found the BEST pair involving the current short pointer!

---

## 📊 Summary

| Approach     | Time  | Space | Idea                             |
| ------------ | ----- | ----- | -------------------------------- |
| Brute Force  | O(n²) | O(1)  | Try all pairs                    |
| Two Pointers | O(n)  | O(1)  | Start wide, move shorter pointer |

---

## 🎯 The Problem-Solving Journey:

1. **Understand the formula** → Area = width × min(heights)
2. **Start dumb** → Try all pairs (brute force)
3. **Ask: Can I skip pairs?** → Yes! Moving taller pointer never helps
4. **Two pointers** → Start wide, greedily move shorter pointer

---

## 🧩 Pattern Recognition:

This is a **Two Pointer** problem where:

- You have two ends to consider
- Moving one end might help, the other won't
- You can make a **greedy choice** about which to move

Similar problems: Trapping Rain Water, Two Sum (sorted array)

---
