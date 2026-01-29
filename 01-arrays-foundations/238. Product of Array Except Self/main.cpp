#include <iostream>
#include <vector>

using namespace std;

// class Solution {
// public:
//     vector<int> productExceptSelf(vector<int> &nums) {
//         vector<int> left(nums.size());
//         left[0] = 1;
//
//         vector<int> right(nums.size());
//         right[nums.size() - 1] = 1;
//
//         for (int i = 1; i < nums.size(); i++) {
//             left[i] = nums[i - 1] * left[i - 1];
//         }
//
//         for (int i = nums.size() - 2; i >= 0; i--) {
//             right[i] = nums[i + 1] * right[i + 1];
//         }
//
//         vector<int> output(nums.size());
//         for (int i = 0; i < nums.size(); i++) {
//             output[i] = left[i] * right[i];
//         }
//
//         return output;
//     }
// };

class Solution {
public:
    vector<int> productExceptSelf(vector<int> &nums) {
        vector<int> left(nums.size());
        left[0] = 1;

        vector<int> right(nums.size());
        right[nums.size() - 1] = 1;

        for (int i = 1; i < nums.size(); i++) {
            left[i] = nums[i - 1] * left[i - 1];
        }

        vector<int> output(nums.size());
        output[nums.size() - 1] = left[nums.size() - 1];

        for (int i = nums.size() - 2; i >= 0; i--) {
            right[i] = nums[i + 1] * right[i + 1];
            output[i] = left[i] * right[i];
        }

        return output;
    }
};

int main() {
    Solution sol;
    // vector<int> nums = {1, 2, 3, 4};
    vector<int> nums = {-1, 1, 0, -3, 3};
    vector<int> result = sol.productExceptSelf(nums);

    for (int i = 0; i < result.size(); i++) {
        cout << result[i] << " ";
    }
    return 0;
}
