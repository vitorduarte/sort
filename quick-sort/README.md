# Quick Sort

It is a **in-place** sorting algorithm, this means that it don't use any auxiliar data to sort the objects. It is also a type of divide anconquer algorithm for sorting an array, based on a routine:

## Algorithm

1. If the range has fewer than two elements, return immediatelly as there is nothing to do.
2. Otherwise pick a value, called a *pivot*, that occurs in the range
3. Partiition the range: reorder its elements, while determining a point of division, so that all elements with values less than the pivot come before the division, while all elements with values greater than the pivot come after it; elements that are equal to the pivot can go either way.
4. Recursivelly apply the quicksort to the sub-range up to the point of division and to the sub-range after it, possibly excluding from both ranges the element equal to the pivot at the point of division.

The choice of partition routine (including the pivot selection) and other details not entirely specified above can affect the algorithm's performance, possibly to a great extent for specific input arrays. In discussing the efficiency of quicksort, it is therefore necessary to specify these choices first. Here we mention two specific partition methods.

## Time complexity

### Worst-case

The worst case occurs when one of the sublist returned by the partitioning routine is of size n-1 this may occur if the pivot happens to be the smallest or largest element in the list, or in some implementations (the Lomuto partition scheme) when all the elements are equal.

If this happens repeatedly in every partition, then each recursive call processes a list of size one less than the previous list. And ![sum](https://wikimedia.org/api/rest_v1/media/math/render/svg/8829d4203c5b6319b5752064f10812e9aa8e3b20), so in that case quicksort takes *O(n²)*

### Best-case

In the most balanced case, each time we perform a partition we divide the list into two nearly equal pieces. This means each recursive call processes a list of half the size. Consequently we can make only *log2 n* nested calls before we reach a list of size 1. This means that the depth of the call tree is *log2 n*. But each level of of the tree needs only *O(n)* time together (as long the array was splited and each subarray iterates one time on it). The result is that the algorith muses only *O(n log n)* time.

### Average case

The average case takes *O(n log n)* time.

## Space complexity

The space used by quicksort depends on the version used

The in-place version of quicksort has a space complexity of *O(log n)*, even in the worst case, when it is carefully implemented using the following strategies.

- In-place partitioning is used. This unstable partition requires O(1) space.
- After partitioning, the partition with the fewest elements is (recursively) sorted first, requiring at most O(log n) space. Then the other partition is sorted using tail recursion or iteration, which doesn't add to the call stack.

[Reference](https://en.wikipedia.org/wiki/Quicksort)
