"""
An implementation of the Fenwick Tree which automatically extends upon exceeding limit.
Extention is O(nlogn)
"""

class FenwickTree:
    """
    Implementation for Fenwick Tree
    """
    def __init__(self, capacity: int = 10, expansion_multiplier: int = 2):
        self.capacity = capacity
        self.multiplier = expansion_multiplier
        self.tree = [0]*self.capacity
        self.bits = [0]*self.capacity

    def add(self, idx: int, delta: int):
        # Add delta to bits and tree.
        if idx<0:
            raise IndexError(f'Index out of bounds: {idx} for tree with capacity: {self.capacity}')
        # Resize tree if necessary
        while idx>=self.capacity:
            self.__extend()
        # Update bits
        self.bits[idx]+=delta
        self.__update(idx, delta)

    def __update(self, idx: int, delta: int):
        # Update fenwick tree
        while idx < self.capacity:
            self.tree[idx]+=delta
            idx=idx|(idx+1)

    def query(self, idx: int) -> int:
        if idx<0:
            raise IndexError(f'Index out of bounds: {idx} for tree with capacity: {self.capacity}')
        if idx>=self.capacity:
            return self.query(self.capacity-1)
        # Cumulative sum from valid nodes
        val = 0
        while idx >= 0:
            val += self.tree[idx]
            idx=(idx&(idx+1))-1
        return val

    def __extend(self):
        print("Extending mapping bits")
        # Resize bits
        self.capacity = self.capacity*self.multiplier
        self.bits.extend([0]*(self.capacity-len(self.bits)))

        self.__rebuild()

    def __rebuild(self):
        print("Rebuilding fenwick tree")
        self.tree = [0]*self.capacity
        for i, val in enumerate(self.bits):
            self.__update(i, val)
