using System.Text.RegularExpressions;

namespace _065_ValidNumber {
    public class Solution {
        private static Regex reg = new Regex(@"^([-+](?=[0-9.]))?\d*((?<=\d)\.|\.(?=\d)|(?<=\d)\.(?=[0-9e]))?\d*((?<=[0-9.])e(?=[-+0-9.])([-+](?=[0-9.]))?\d+)?$", RegexOptions.Compiled | RegexOptions.Singleline);
        public bool IsNumber(string s) {
            s = s.Trim();
            if(s == string.Empty) {
                return false;
            }
            Match match = reg.Match(s);
            return match.Value == s;
        }
    }
}
