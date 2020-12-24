namespace _003_LongestSubstringWithoutRepeatingCharacters {
    public class Solution {
        public int LengthOfLongestSubstring(string s) {
            int[] sn = new int[s.Length];
            for(int i = 0; i < s.Length; i++) {
                int n = s.IndexOf(s[i], i + 1);
                sn[i] = n < 0 ? s.Length : n;
            }
            int max = 0;
            for(int i = 0; i < sn.Length; i++) {
                for(int j = sn[i] - 1; j > i; j--) {
                    if(sn[j] < sn[i]) {
                        sn[i] = sn[j];
                    }
                }
                if(sn[i] - i > max) {
                    max = sn[i] - i;
                }
            }
            return max;
        }
    }
}
