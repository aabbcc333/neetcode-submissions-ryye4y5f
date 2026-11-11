class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
      std::sort(nums.begin(), nums.end()); 

    int i = 0 ;
    while(i+1 < nums.size()){
        if(nums[i] == nums[i+1]){
        return true ;
        }
        i++;
    }
      return false ;
    }

};