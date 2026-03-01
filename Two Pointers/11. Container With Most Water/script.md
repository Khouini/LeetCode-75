water between two lines = width × height
width = how far apart they are
height = the SHORT one (water spills over the short one)

area = (R - L) × min(height[L], height[R])

heights = [1, 8, 6, 2, 5, 4, 8, 3, 7]

Step 1: Try every pair (the dumb way). Write them all out.

L=0, R=1: (1) × min(1,8) = 1
L=0, R=2: (2) × min(1,6) = 2
L=0, R=8: (8) × min(1,7) = 8
L=1, R=2: (1) × min(8,6) = 6
L=1, R=8: (7) × min(8,7) = 49 ← biggest
...36 pairs total. Too many. Can we be smarter?

Step 2: What makes area big? Two things: wide AND tall.

Where is the WIDEST pair? L at the start, R at the end.
Start there. That gives us the best possible width.

[1, 8, 6, 2, 5, 4, 8, 3, 7]
 L                       R

area = 8 × min(1, 7) = 8

Step 3: We have to shrink the width by 1. No choice. L or R must move inward.

Width is going down no matter what. So we need height to go UP.

Look at our two lines: L=1 (short), R=7 (tall).
The height is stuck at 1 because of L. L is the bottleneck.

If we move R inward: width shrinks, but height is STILL stuck at 1 (L is still there, still short).
  → guaranteed worse.

If we move L inward: width shrinks, but maybe we find a taller line than 1.
  → maybe better.

So move L. It's the short one. It's the bottleneck. Get rid of it.

Step 4: Just keep doing that. Always move the short one.

[1, 8, 6, 2, 5, 4, 8, 3, 7]
 L                       R   area = 8 × 1 = 8.   L is short, move L.
    L                    R   area = 7 × 7 = 49.  R is short, move R.
    L                 R      area = 6 × 3 = 18.  R is short, move R.
    L              R         area = 5 × 8 = 40.  equal, move either.
       L           R         area = 4 × 6 = 24.  L is short, move L.
          L        R         area = 3 × 2 = 6.   L is short, move L.
             L     R         area = 2 × 5 = 10.  L is short, move L.
                L  R         area = 1 × 4 = 4.   L is short, move L.
                LR           they met. stop.

max area = 49 ✅

Step 5: Why is this safe? Why don't we miss good pairs?

When we throw away the short pointer, think about what we're skipping:
all the pairs where that short line is one side, with a SMALLER width.
Same short height, less width → less area. Nothing good there.
