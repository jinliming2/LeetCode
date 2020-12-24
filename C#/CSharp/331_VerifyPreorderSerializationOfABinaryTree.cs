using System.Collections.Generic;

namespace _331_VerifyPreorderSerializationOfABinaryTree {
    public class Solution {
        private struct node {
            public string character;
            public bool left;
            public bool right;
        }
        public bool IsValidSerialization(string preorder) {
            if(preorder == "#") {
                return true;
            }
            bool root = false;
            string[] list = preorder.Split(',');
            Stack<node> stack = new Stack<node>();
            foreach(string str in list) {
                if(stack.Count == 0) {
                    if(root) {
                        return false;
                    } else {
                        root = true;
                    }
                }
                while(stack.Count > 0) {
                    node lasts = stack.Pop();
                    if(!lasts.left || !lasts.right) {
                        stack.Push(lasts);
                        break;
                    }
                }
                if(str == "#") {
                    if(stack.Count <= 0) {
                        return false;
                    }
                    node last = stack.Pop();
                    if(!last.left) {
                        last.left = true;
                        stack.Push(last);
                    } else if(last.right) {
                        return false;
                    }
                } else {
                    if(stack.Count > 0) {
                        node last = stack.Pop();
                        if(!last.left) {
                            last.left = true;
                        } else {
                            last.right = true;
                        }
                        stack.Push(last);
                    }
                    node current;
                    current.character = str;
                    current.left = false;
                    current.right = false;
                    stack.Push(current);
                }
            }
            while(stack.Count > 0) {
                node lasts = stack.Pop();
                if(!lasts.left || !lasts.right) {
                    stack.Push(lasts);
                    break;
                }
            }
            return stack.Count == 0;
        }
    }
}
