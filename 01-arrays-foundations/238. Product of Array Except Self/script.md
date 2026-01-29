answer[i] = (numbers before i) × (numbers after i)
(before me) × (after me)

Step 1: Write out what you need for EACH index (the dumb way)

i=0: (nothing before) × (2 × 3 × 4) = 1 × 24 = 24
i=1: (1) × (3 × 4) = 1 × 12 = 12
i=2: (1 × 2) × (4) = 2 × 4 = 8
i=3: (1 × 2 × 3) × (nothing) = 6 × 1 = 6

Step 2: Look at the "left" part only. Write them all out.
i=0: nothing before → we say 1 (neutral for multiplication)
i=1: 1
i=2: 1 × 2 = 2
i=3: 1 × 2 × 3 = 6
prefix array = [1, 1, 2, 6]

Step 3: Stare at the prefix array. Do you see a pattern?
prefix[0] = 1 (base case, nothing before)
prefix[1] = 1 = prefix[0] × nums[0] = 1 × 1
prefix[2] = 2 = prefix[1] × nums[1] = 1 × 2
prefix[3] = 6 = prefix[2] × nums[2] = 2 × 3

Pattern: prefix[i] = prefix[i-1] × nums[i-1]

Step 4: Do the same for "right" part. Write them all out.
i=0: 2 × 3 × 4 = 24
i=1: 3 × 4 = 12
i=2: 4
i=3: nothing after → we say 1
Right array = [24, 12, 4, 1]

Step 5: Stare at the suffix array. Pattern?
suffix[3] = 1 (base case, nothing after)
suffix[2] = 4 = suffix[3] × nums[3] = 1 × 4
suffix[1] = 12 = suffix[2] × nums[2] = 4 × 3
suffix[0] = 24 = suffix[1] × nums[1] = 12 × 2

Pattern: suffix[i] = suffix[i+1] × nums[i+1]

Step 6: Multiply prefix × suffix for each index
i=0: prefix[0] × suffix[0] = 1 × 24 = 24 ✅
i=1: prefix[1] × suffix[1] = 1 × 12 = 12 ✅
i=2: prefix[2] × suffix[2] = 2 × 4 = 8 ✅
i=3: prefix[3] × suffix[3] = 6 × 1 = 6 ✅
