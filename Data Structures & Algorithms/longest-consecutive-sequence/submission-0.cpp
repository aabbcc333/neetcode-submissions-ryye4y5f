class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        int res = 0 ;
        std::unordered_set<int> numsSet(nums.begin(),nums.end());
        for(auto num : numsSet){
          if(numsSet.find(num-1) == numsSet.end()){
            int length = 1 ;
            while(numsSet.find(num+length) != numsSet.end()){
                length++ ;
            }
            res = max(res,length);
          }
        }
        return res;
    }
};
