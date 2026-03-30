class Solution {
public:
    int trap(vector<int>& height) {
        const int length = height.size();
        if(length<3) return 0 ;
        int L =0;
        int R = length-1 ;
        int lmax = 0;
        int rmax = 0 ;
        long long total = 0;

        while(L<R){
            lmax = max(lmax,height[L]);
            rmax = max(rmax, height[R]);

         if(lmax<=rmax){
            total += lmax-height[L];
            L++;
         }
         else{
            total += rmax- height[R];
            R--;
         }
        }
        return (int)total ;
    }
};
