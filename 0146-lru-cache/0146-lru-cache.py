# doubly linked lists
class dnode:
    def __init__(self,key=0,value=0):
        self.prev=None
        self.next=None
        self.value=value
        self.key=key

# use hashtable, store node, and double linked list can help decrease the complexity
class LRUCache:
    def __init__(self, capacity: int):
        self.cache=dict()
        self.head=dnode()
        self.tail=dnode()
        # use 2 dummy node
        self.head.next=self.tail
        self.tail.prev=self.head
        self.cap=capacity
        self.size=0
    
    def add(self,node):
        # insert to head
        node.prev=self.head
        node.next=self.head.next
        self.head.next.prev=node
        self.head.next=node
    
    def delete(self,node):
        # move a node
        node.prev.next=node.next
        node.next.prev=node.prev
    
    def movehead(self,node):
        self.delete(node)
        self.add(node)

    def removetail(self):
        node=self.tail.prev # tail is dummy node, real tail is tail.prev
        self.delete(node)
        return node

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        node=self.cache[key]
        self.movehead(node)
        return node.value

    def put(self, key: int, value: int) -> None:
        if key not in self.cache:
            node=dnode(key,value)
            self.cache[key]=node
            self.add(node)
            self.size+=1
            if self.size>self.cap:
                remove=self.removetail()
                self.cache.pop(remove.key)
                self.size-=1
        else:
            node=self.cache[key]
            node.value=value
            self.movehead(node)



# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)