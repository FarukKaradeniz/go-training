# What does this app do?

1. Asks for user input to set the array's values by manual or random.
2. If the user chooses by random, values are set with the "crypto/rand" package.
3. After the creation of the array, the user can select a sort algorithm to sort the created array.

## Using Docker

To run the application:
``
docker run -it farukkaradeniz/go-training
``

To run the tests:
``
docker run farukkaradeniz/go-training sh -c "go test"
``

## TODO

- [x] Bubble Sort
- [x] Selection Sort
- [x] Merge Sort
- [x] Quick Sort
- [ ] Heap Sort
- [ ] Insertion Sort
- [ ] Shell Sort
- [ ] Radix Sort
