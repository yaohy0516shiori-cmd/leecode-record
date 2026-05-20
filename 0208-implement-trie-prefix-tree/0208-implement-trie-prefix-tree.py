class treenode:
    def __init__(self):
        self.children={}
        self.end=False
# 不同于二叉树但是形式类似, 总之要有办法指向下一个节点
class Trie:

    def __init__(self):
        self.root=treenode()

    def insert(self, word: str) -> None:
        node=self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch]=treenode()
            node=node.children[ch]
        node.end=True
# 能出来不代表一定完全匹配
    def search(self, word: str) -> bool:
        node=self.root
        for ch in word:
            if ch not in node.children:
                return False
            node=node.children[ch]
        return node.end
# 只要能出来就可以返回搜索到了
    def startsWith(self, prefix: str) -> bool:
        node=self.root
        for ch in prefix:
            if ch not in node.children:
                return False
            node=node.children[ch]
        return True


# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)