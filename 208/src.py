#-*- coding: utf-8 -*-
"""
leetcode: 208.Implement Trie (Prefix Tree)
          https://leetcode.com/problems/implement-trie-prefix-tree/
"""


class TrieNode(object):
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.children = {}
        self.word_end = False


class Trie(object):

    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        """
        Inserts a word into the trie.
        :type word: str
        :rtype: void
        """
        node = self.root
        for w in word:
            if node.children.get(w) is None:
                new = TrieNode()
                node.children[w] = new
                node = new
            else:
                node = node.children.get(w)

        node.word_end = True


    def search(self, word):
        """
        Returns if the word is in the trie.
        :type word: str
        :rtype: bool
        """
        node = self.root
        for w in word:
            if node.children.get(w):
                node = node.children.get(w)
            else:
                return False

        return node.word_end


    def startsWith(self, prefix):
        """
        Returns if there is any word in the trie
        that starts with the given prefix.
        :type prefix: str
        :rtype: bool
        """
        node = self.root
        for w in prefix:
            if node.children.get(w):
                node = node.children.get(w)
            else:
                return False

        return True

# Your Trie object will be instantiated and called as such:
# trie = Trie()
# trie.insert("somestring")
# trie.search("key")
