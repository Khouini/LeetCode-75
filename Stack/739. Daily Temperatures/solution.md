# 739. Daily Temperatures – Stack

---

## What's the Problem?

You are given an array of daily `temperatures`, like `[73, 74, 75, 71, 69, 72, 76, 73]`.
For each day, you need to find out **how many days you have to wait** until a warmer temperature occurs. If there is no future day that is warmer, you put `0`.

Return an array representing these waiting periods: `[1, 1, 4, 2, 1, 1, 0, 0]`.

---

## Step-by-Step Thinking

### Question 1: Why is the naive approach too slow?

The brute-force way is to stand at each day and scan the rest of the array to your right looking for a warmer temperature. In the worst-case scenario (e.g., temperatures keep dropping: `[100, 90, 80, 70]`), you will scan to the end for _every single day_. That is an $O(n^2)$ operation, which will time out for large arrays.

> **Observation:** We are looking for the **Next Greater Element**. Whenever a problem asks you to find the "next greater" or "next smaller" element efficiently, it is a very strong signal to use a **Monotonic Stack**.

---

### Question 2: How can a stack help us avoid scanning twice?

Imagine you are looking at temperatures day by day. If today is colder than yesterday, you still don't know the answer for yesterday. But you also don't know the answer for today! You are basically building a "waiting list" of days that are waiting for a warmer temperature.

If today suddenly gets really hot, it might resolve the wait time for several past days on your waiting list all at once.

> **Observation:** A stack is perfect for a "waiting list". We push temperatures onto the stack when they decrease. Because we only keep strictly decreasing temperatures waiting, it's called a **Monotonic Decreasing Stack**.

---

### Question 3: What exactly do we need to push onto the stack?

When we finally find a warmer day, we need to calculate _how many days have passed_.
If we only store the temperature value in the stack (e.g., `73`), we won't know how to calculate the distance between that past day and "today".

> **Observation:** We don't just need the past temperature; we need to know _when_ it happened. We should store the **index** of the day in the stack (which also allows us to look up the temperature in the original array).
> `wait_days = current_index - past_index`

---

### Question 4: What is the core logic when checking a new day?

For each new day (let's say index `i` and temperature `T`), we look at the top of our stack:

1. Is the stack empty? Just push the current day `i` onto the stack and wait.
2. Is `T` **greater** than the temperature of the day at the top of the stack?
   - Yes! We found a warmer day for the day at the top of the stack.
   - **Pop** that past day from the stack.
   - Calculate the distance: `i - pop_index`.
   - Save this distance in our result array at `result[pop_index]`.
   - **Repeat** this check. `T` might be warmer than the next day down in the stack too!
3. Is `T` **less than or equal to** the top of the stack? Just stack it on top and wait.

---

## Full Walkthrough

Input: `temperatures = [73, 74, 75, 71, 69, 72, 76, 73]`

| i   | Temp | Stack (stores indices) | Action / Logic                                                                                                                                    | Result Array `ans`  |
| --- | ---- | ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------- |
| `-` | `-`  | `[]`                   | Initial state. Result array is all `0`s.                                                                                                          | `[0,0,0,0,0,0,0,0]` |
| `0` | `73` | `[0]`                  | Stack is empty. Push index `0`.                                                                                                                   | `[0,0,0,0,0,0,0,0]` |
| `1` | `74` | `[1]`                  | `74 > temps[0]` (73). Pop `0`. <br>`ans[0] = 1 - 0 = 1`. Push `1`.                                                                                | `[1,0,0,0,0,0,0,0]` |
| `2` | `75` | `[2]`                  | `75 > temps[1]` (74). Pop `1`. <br>`ans[1] = 2 - 1 = 1`. Push `2`.                                                                                | `[1,1,0,0,0,0,0,0]` |
| `3` | `71` | `[2, 3]`               | `71 < temps[2]` (75). Just push `3`.                                                                                                              | `[1,1,0,0,0,0,0,0]` |
| `4` | `69` | `[2, 3, 4]`            | `69 < temps[3]` (71). Just push `4`.                                                                                                              | `[1,1,0,0,0,0,0,0]` |
| `5` | `72` | `[2, 3]`               | `72 > temps[4]` (69). Pop `4`. `ans[4] = 5 - 4 = 1`. <br>`72 > temps[3]` (71). Pop `3`. `ans[3] = 5 - 3 = 2`. <br>`72 < temps[2]` (75). Push `5`. | `[1,1,0,2,1,0,0,0]` |
| `6` | `76` | `[6]`                  | `76 > temps[5]` (72). Pop `5`. `ans[5] = 6 - 5 = 1`. <br>`76 > temps[2]` (75). Pop `2`. `ans[2] = 6 - 2 = 4`.<br>Push `6`.                        | `[1,1,4,2,1,1,0,0]` |
| `7` | `73` | `[6, 7]`               | `73 < temps[6]` (76). Push `7`.                                                                                                                   | `[1,1,4,2,1,1,0,0]` |

End of array. The remaining indices in the stack (`6` and `7`) never found a warmer day, so their answers correctly remain `0`.

**Result: `[1, 1, 4, 2, 1, 1, 0, 0]`**
