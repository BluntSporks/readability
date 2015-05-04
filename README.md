# readability
Reusable Golang library to provide readability scores

## Status
Under development; not ready for use

## Scores
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

The apostrophe s at the the of possessives is ignored because it does not increase syllable count in the general case.

## Short forms
See [short-names](https://www.github.com/BluntSporks/short-names) for a list of abbreviations used.
