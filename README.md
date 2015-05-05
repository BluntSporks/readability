# readability
Reusable Golang library to provide readability scores

# Status
Ready to use.

However, this package has not been extensively tested. Results may differ in small amounts from official scoring of
these measures and should not be taken as official.

## Scores
You can learn about [readability tests](https://en.wikipedia.org/wiki/Readability_test) from Wikipedia.

The scores are planned to be:
* [Automated Readability Index](https://en.wikipedia.org/wiki/Automated_Readability_Index)
* [Coleman-Liau index](https://en.wikipedia.org/w/index.php?title=Coleman-Liau_Index)
* [Flesh-Kincaid](https://en.wikipedia.org/w/index.php?title=Flesch-Kincaid)
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

## Short forms
See [short-names](https://www.github.com/BluntSporks/short-names) for a list of abbreviations used.
