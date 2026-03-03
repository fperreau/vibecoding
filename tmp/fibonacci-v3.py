import unittest
from functools import lru_cache


@lru_cache(maxsize=None)
def fibonacci_memo(n: int) -> int:
    """
    Compute the nth Fibonacci number using memoization (lru_cache).

    This is an optimized recursive approach with O(n) time complexity
    on first call and O(1) for subsequent calls to already-computed values.

    Args:
        n (int): The index (0-based) of the Fibonacci number to compute.

    Returns:
        int: The nth Fibonacci number.

    Raises:
        ValueError: If n is negative.
    """
    if n < 0:
        raise ValueError(f"n must be a non-negative integer, got {n}")
    if n <= 1:
        return n
    return fibonacci_memo(n - 1) + fibonacci_memo(n - 2)


def fibonacci(n: int):
    """
    Generate the Fibonacci sequence up to n terms using an optimized generator.

    Uses an iterative approach with O(n) time and O(1) space complexity,
    avoiding recursion overhead and stack limits.

    Args:
        n (int): The number of terms in the Fibonacci sequence to generate.

    Yields:
        int: The next term in the Fibonacci sequence.

    Raises:
        ValueError: If n is negative.
    """
    if n < 0:
        raise ValueError(f"n must be a non-negative integer, got {n}")
    a, b = 0, 1
    for _ in range(n):
        yield a
        a, b = b, a + b


def fibonacci_list(n: int) -> list[int]:
    """
    Return the first n Fibonacci numbers as a list.

    Args:
        n (int): The number of terms to return.

    Returns:
        list[int]: A list of the first n Fibonacci numbers.
    """
    return list(fibonacci(n))


def optimize(n: int) -> list[int]:
    """
    Optimized function returning the first `n` Fibonacci numbers as a list.

    Uses an iterative O(n) approach with O(1) additional space.

    Args:
        n (int): Number of terms to return.

    Returns:
        list[int]: The first n Fibonacci numbers.

    Raises:
        ValueError: If `n` is negative.
    """
    if n < 0:
        raise ValueError(f"n must be a non-negative integer, got {n}")
    a, b = 0, 1
    result: list[int] = []
    for _ in range(n):
        result.append(a)
        a, b = b, a + b
    return result


# ---------------------------------------------------------------------------
# Unit Tests
# ---------------------------------------------------------------------------

class TestFibonacciGenerator(unittest.TestCase):
    """Tests for the fibonacci() generator function."""

    def test_first_ten_terms(self):
        expected = [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
        self.assertEqual(list(fibonacci(10)), expected)

    def test_zero_terms(self):
        self.assertEqual(list(fibonacci(0)), [])

    def test_single_term(self):
        self.assertEqual(list(fibonacci(1)), [0])

    def test_two_terms(self):
        self.assertEqual(list(fibonacci(2)), [0, 1])

    def test_large_sequence(self):
        result = list(fibonacci(20))
        self.assertEqual(len(result), 20)
        self.assertEqual(result[-1], 4181)  # 20th Fibonacci number (0-indexed: F(19))

    def test_negative_raises(self):
        with self.assertRaises(ValueError):
            list(fibonacci(-1))

    def test_is_generator(self):
        import types
        self.assertIsInstance(fibonacci(5), types.GeneratorType)


class TestFibonacciMemo(unittest.TestCase):
    """Tests for the fibonacci_memo() memoized function."""

    def setUp(self):
        # Clear cache before each test to ensure isolation
        fibonacci_memo.cache_clear()

    def test_base_cases(self):
        self.assertEqual(fibonacci_memo(0), 0)
        self.assertEqual(fibonacci_memo(1), 1)

    def test_known_values(self):
        known = {2: 1, 3: 2, 4: 3, 5: 5, 6: 8, 7: 13, 10: 55, 15: 610}
        for n, expected in known.items():
            with self.subTest(n=n):
                self.assertEqual(fibonacci_memo(n), expected)

    def test_large_value(self):
        # F(50) is a well-known value
        self.assertEqual(fibonacci_memo(50), 12586269025)

    def test_negative_raises(self):
        with self.assertRaises(ValueError):
            fibonacci_memo(-5)

    def test_cache_is_used(self):
        fibonacci_memo(30)
        info = fibonacci_memo.cache_info()
        self.assertGreater(info.hits, 0)

    def test_consistency_with_generator(self):
        """Memoized results must match the generator output."""
        gen_results = list(fibonacci(15))
        for i, val in enumerate(gen_results):
            with self.subTest(i=i):
                self.assertEqual(fibonacci_memo(i), val)


class TestFibonacciList(unittest.TestCase):
    """Tests for the fibonacci_list() helper."""

    def test_returns_list(self):
        self.assertIsInstance(fibonacci_list(5), list)

    def test_correct_values(self):
        self.assertEqual(fibonacci_list(7), [0, 1, 1, 2, 3, 5, 8])

    def test_empty(self):
        self.assertEqual(fibonacci_list(0), [])


class TestOptimize(unittest.TestCase):
    """Tests for the optimize() helper."""

    def test_returns_list(self):
        self.assertIsInstance(optimize(5), list)

    def test_correct_values(self):
        self.assertEqual(optimize(7), [0, 1, 1, 2, 3, 5, 8])

    def test_empty(self):
        self.assertEqual(optimize(0), [])

    def test_negative_raises(self):
        with self.assertRaises(ValueError):
            optimize(-3)


if __name__ == "__main__":
    unittest.main(verbosity=2)
