using _003_LongestSubstringWithoutRepeatingCharacters;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace CSharpUnitTest {
    [TestClass]
    public class _003_LongestSubstringWithoutRepeatingCharactersTest {
        Solution solution;
        public _003_LongestSubstringWithoutRepeatingCharactersTest() {
            solution = new Solution();
        }
        [TestMethod]
        public void LengthOfLongestSubstringTest() {
            Assert.AreEqual(3, solution.LengthOfLongestSubstring("abcabcbb"));
            Assert.AreEqual(1, solution.LengthOfLongestSubstring("bbbbb"));
            Assert.AreEqual(3, solution.LengthOfLongestSubstring("pwwkew"));
        }
    }
}
