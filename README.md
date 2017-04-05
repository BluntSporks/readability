# readability
Reusable Golang library to provide readability scores

# Status
Ready to use.

However, this package has not been extensively tested. Results may differ in small amounts from official scoring of
these measures and should not be taken as official.

## Scores
You can learn about [readability tests](https://en.wikipedia.org/wiki/Readability_test) from Wikipedia.

The scores are:
* [Automated Readability Index](https://en.wikipedia.org/wiki/Automated_Readability_Index)
* [Coleman-Liau index](https://en.wikipedia.org/w/index.php?title=Coleman-Liau_Index)
* [Flesh-Kincaid](https://en.wikipedia.org/w/index.php?title=Flesch-Kincaid)
* [Flesch reading-ease](https://en.wikipedia.org/wiki/Flesch–Kincaid_readability_tests#Flesch_reading_ease)
* [Gunning fog index](https://en.wikipedia.org/wiki/Gunning_fog_index)
* [SMOG](https://en.wikipedia.org/wiki/SMOG)

To calculate these scores requires:
* Splitting sentences
* Counting syllables

An analysis of English syllable counts shows the following summary statistics:
* Avg chars per word: 7.410083
* Avg syls per word: 2.465334
* Avg multiple: 3.005711

So the best estimate of the number of syllables in a word is to take the number of characters and divide by 3. For
example, the word "syllables" has 9 characters and 3 syllables, so it fits this pattern. The word "polysyllabic" has 12
characters but 5 syllables, where the formula predicts only 4. However, such errors should cancel each other out in
general, so the estimate should be good enough.

The best results are yielded by using the formula s = round(float32(c)/3.0) where s is estimated syllable count and c
is character count. The round function is better than the floor function or ceil function for this purpose. The formula
is exactly correct 58% of the time, and is within +/-1 syllable of the correct count 98% of the time. On average, the
positive and negative deviations should cancel each other out, at least in some formulas.

## Performance
Testing of this program shows that it takes an average of 73 seconds to process one ebook on my desktop machine on 409
ebooks, yielding 409 clusters of 5 grade levels each.

Average grade level of each cluster ranged from a minimum of 5.24 to a maximum of 16.84. The average of the 409 average
grade levels was 10.7, a mid-high school reading level for the material I happen to have in my collection. That seems
fairly accurate.

Each cluster of 5 grades was also given a sample standard deviation. The stddevs ranged from a minimum of 0.17 grade
levels to a maximum of 3.41. The average of the 409 stddevs was 0.72 grade levels. Thus the program scores are
typically in fairly good agreement as to what constitutes the correct reading level. At a stddev of 0.72, we expect 95%
of estimated grades to fall between +/-1.42 grade levels.

An example of its output that happens to have exactly a stddev of 0.72 grade levels, and thus is typical, is the book
"The Formula: How Algorithms Solve All Our Problems . . . and Create More", by Luke Dormehl, which was evaluated as
follows:

    Automated Readability: 13.76
    Coleman-Liau: 11.82
    Flesch-Kincaid: 12.72
    Gunning fog: 12.50
    SMOG: 12.24
    Sorted scores: [11.82, 12.24, 12.50, 12.72, 13.76]
    Average score: 12.61
    Std Dev of scores: 0.72

This output shows that the scores, even if not perfectly accurate, are still in close enough agreement to be of
practical use.

## Programming notes
See [abbreviation](https://www.github.com/BluntSporks/abbreviation) for a list of abbreviations used.
