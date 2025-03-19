# Go Sort
Go solution to Coding Challenges [build your own sort](https://codingchallenges.fyi/challenges/challenge-sort).

# Testing

## Step 0 
Set up the testing by running":

```bash
curl https://www.gutenberg.org/cache/epub/132/pg132.txt -o test.txt
tr -s '[[:punct:][:space:]]' '\n' < test.txt |sed '/^[0-9]/d' > words.txt
```

## Step 1
```bash
% ./gosort words.txt | uniq | head -n5
A
ACTUAL
AGREE
AGREEMENT
AND
```

## Step 2
```bash
% ./gosort -u words.txt | head -n5
A
ACTUAL
AGREE
AGREEMENT
AND
```

## Step 3

### Quicksort
```bash
 % ./gosort --qsort words.txt | uniq | head -n5
A
ACTUAL
AGREE
AGREEMENT
AND
```

### Mergesort
```bash
 % ./gosort --mergesort words.txt | uniq | head -n5
A
ACTUAL
AGREE
AGREEMENT
AND
```

## Step 4
```bash
 % ./gosort --random-sort words.txt | uniq | head -n5
A
ACTUAL
AGREE
AGREEMENT
AND
```