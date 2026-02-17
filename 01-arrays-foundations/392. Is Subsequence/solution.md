This problem is **LeetCode 392 – Is Subsequence**.

You are given two strings `s` and `t`.
You must check if **all characters of `s` appear in `t` in the same order**.

---

# ✅ Brute Force Solution (Simple & Step-by-Step)

The brute force idea is very straightforward:

👉 Try to match each character of `s` one by one inside `t`.

We scan `t` from left to right and try to match characters from `s`.

---

## 🧠 Step-by-Step Idea

1. Start with two pointers:
   - `i` for string `s`
   - `j` for string `t`

2. Compare `s[i]` with `t[j]`
   - If they match → move `i` forward (we found one character)
   - Always move `j` forward

3. If we reach the end of `s`, it means we found all characters in order → return `true`

4. If we finish scanning `t` but not `s` → return `false`

---

## 🔎 Example 1

```
s = "abc"
t = "ahbgdc"
```

| s pointer | t pointer | comparison | action    |
| --------- | --------- | ---------- | --------- |
| a         | a         | match      | move both |
| b         | h         | no         | move t    |
| b         | b         | match      | move both |
| c         | g         | no         | move t    |
| c         | d         | no         | move t    |
| c         | c         | match      | move both |

We matched all letters of `s` ✅
Return **true**

---

## 🔎 Example 2

```
s = "axc"
t = "ahbgdc"
```

We find `a` ✔
But we never find `x` ❌

Return **false**

---

# 💻 Brute Force Code (C++)

```cpp
bool isSubsequence(string s, string t) {
    int i = 0; // pointer for s
    int j = 0; // pointer for t

    while (i < s.length() && j < t.length()) {
        if (s[i] == t[j]) {
            i++;  // move s pointer if match
        }
        j++;      // always move t pointer
    }

    return i == s.length();
}
```

---

# ⏱ Time Complexity

- We scan `t` once.
- **O(n)** where n = length of `t`

Even though it's called "brute force", this is already efficient.

---

# 🎯 Why This Works

We only care about:

- Order must stay the same
- Characters do not need to be consecutive

So we just check if we can "build" `s` while walking through `t`.

---
