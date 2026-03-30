class Solution {
public:
    int maxArea(vector<int>& heights) {
        const int n = heights.size();
        int L = 0;
        int R = n-1 ;
        int area = 0 ;
        while(L<R){
            if(heights[L]<=heights[R]){
            area = max(area,(R-L)*heights[L]);
            L++;
            }
            else{
                area = max(area, (R-L)*heights[R]);
            R--;
            }
        }
        return area ;
        
    }
};
