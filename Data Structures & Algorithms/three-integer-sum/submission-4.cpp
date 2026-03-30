class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> res ;
        int n = nums.size(); 
        sort(nums.begin(), nums.end());
        for(int i = 0 ; i < n ; i++){
            if (nums[i]>0) break ;
            if(i > 0 && nums[i] == nums[i-1]) continue ;
            int l = i +1 ;
            int r = n-1 ;
            
            while(l < r){
                int total = nums[i] + nums[l] + nums[r];
                if(total<0){
                    l++;
                }else if(total>0){
                    r--;
                }else{
                    res.push_back({nums[i],nums[l],nums[r]});
                    l++;
                    r--;
                    while(l < r && nums[l]==nums[l-1]){
                        l++;
                    }
                }
            }
        }
        return res ;
    }
};
