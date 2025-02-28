"""
This test module tests the functionality of fenwick tree.
"""

from fenwick_tree import FenwickTree

def test_add_query_correctness():
    """
    Tests the correctness of the cumulative sum tracked by the fenwick tree.
    """
    tree = FenwickTree(capacity=20)
    # Input
    test_input = [(10, 3), (0, 1), (0, 1), (5, 2)]
    test_output = [(0,2), (5,4), (10, 7)]
    for idx, delta in test_input:
        tree.add(idx, delta)
    # Testing
    for idx, delta in test_output:
        assert delta == tree.query(idx)

def test_capacity_extension():
    """
    Tests the capacity extension feature of the fenwick tree.
    """
    small_tree = FenwickTree(capacity=5, expansion_multiplier=2)
    big_tree = FenwickTree(capacity=20)

    # Test
    test_input = [(0, 1), (5, 2), (3,0), (4,100), (6,1), (9,3), (12,1)]
    for idx, delta in test_input:
        small_tree.add(idx, delta)
        big_tree.add(idx, delta)

    # Validate extension: 5*2*2
    assert small_tree.capacity == 20

    # Trees must be similar
    for i in range(20):
        assert small_tree.query(i) == big_tree.query(i)
        assert small_tree.bits[i] == big_tree.bits[i]
        assert small_tree.tree[i] == big_tree.tree[i]
