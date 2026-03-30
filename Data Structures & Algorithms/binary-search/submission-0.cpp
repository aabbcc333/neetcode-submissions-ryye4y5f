int targetInNums(const vector<int>& nums, int low, int high, int target) {
    if (low > high) return -1;

    int mid = low + (high - low) / 2;

    if (nums[mid] == target) return mid;

    if (nums[mid] < target)
        return targetInNums(nums, mid + 1, high, target);

    return targetInNums(nums, low, mid - 1, target);
}

class Solution {
public:
    int search(vector<int>& nums, int target) {
        return targetInNums(nums, 0, nums.size() - 1, target);
    }
};