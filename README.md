# Good URLs  <!-- omit in toc -->

- [Introduction](#introduction)
- [Setting up fastText](#setting-up-fasttext)
- [Background Preprocessing the Input Text](#background-preprocessing-the-input-text)
- [Estimating Accuracy on Unseen](#estimating-accuracy-on-unseen)
- [Making Predictions Using the Neural Network](#making-predictions-using-the-neural-network)
- [Assembling the Annual Money Laundering Report](#assembling-the-annual-money-laundering-report)

## Introduction

Given a URL's text, can a neural network predict if it is going to be a
good URL?

We're going to find out. What is good or bad does not need to be specified.
All that is needs is examples, 100s, 1000s. What Good URL means is between
you and the trained neural network.

We use a neural network to predict status: good or bad, of a URL.
We retrieve the visible text from a URL and use that the input for network.
The library used is [fastText](https://fasttext.cc/). We use the text
classification feature and follows the dataflow, as shown like this
[tutorial](https://fasttext.cc/docs/en/supervised-tutorial.html) from the project.

## Setting up fastText

The input files need to be UTF-8 with Unix style line endings, CR.
This tutorial assumes that you're running on Windows but could be executed on
Linux. We are going to be working at the terminal. Here are some common options:

1. `PowerShell` - not recommended, it writes UTF-16 by default.
2. `DOS Command Prompt` - not recommended, it writes Windows-style line ending, CRLF.
3. `Git Bash` - recommended, it writes UTF-8 and Unix style line ending.

Copy the files in the `bin` directory to somewhere on your `$PATH`
environment variable. I recommend a `/c/Users/rpereda/bin` directory and
that directory is on my `$PATH` environment variable.

Here are the files:

1. `fasttext.exe` - the command-line interface to fastText
2. `fasttext.dll` - a supporting dynamic library for the .exe
3. `normalize.exe` - a line-by-line, character-by-character text normalizer
4. `shuffle.exe` - shuffler of lines, required for the learning process
5. `rowpaste.exe` - two CSV files row wise.
6. `rowcut.exe` - cut out columns of a CSV file row wise.

fasttext.exe and fasttext.dll are for the command-line interface to the library.

```bash
$ ls -1sh
... fill in later
```

| file                          | purpose                                                 |
| ----------------------------- | ------------------------------------------------------- |
| training.xlsx                 | unformatted training data used from the MLR 2018 Report |
| model_goodness.bin            | neural network created during training                  |
| model_goodness.vec            | dictionary created during training                      |
| prediction-model??.txt        | prediction on the training data; should be 99% accurate |

1. the initial human vetted training data: training.xlsx
2. created using Excel to narrow down to this file: training.prn
3. `cat training.prn | normalize > training-n.prn`
4. `shuffle training-n.prn > training-ns.prn`
5. `fasttext supervised -input training-ns.prn -output model_goodness -lr 1.0 -epoch 25 -wordNgrams 2`
   That creates two files: model_goodness.{bin, vec}.
6. `fasttext test model_goodness.bin training-ns.prn`
   This measures the accuracy of the model. 

## Background Preprocessing the Input Text

Here is an excerpt of the Go program that normalizes the text. Line by line, text
is transformed. The label, in this case the country, is prefaced by ```___label___```.
This prefix is how fasttext distinguishes labels in a large amount of text.
Spaces surround some punctuations. Double quotes are deleted. Semicolons and
colons are replaced by a space. Each white space runs is compressed to a single space.

Digits are mapped to the symbol @. The reason for doing this is for the neural network
to focus on learning the grammar of address based on the shape numbers, not a specific
number. By shape of numbers, we mean the number of digits. So, we know that address that
ends with 5 digits, represented by @@@@@, more readily see the pattern than if we focus
on specific zipcode. In effect 90803 and 92705 are treated a single word, namely a

```go
line := scanner.Text()
line = strings.ToLower(line)
line = "__label__" + line
line = strings.ReplaceAll(line, "'", " ' ")
line = strings.ReplaceAll(line, `"`, "")
line = strings.ReplaceAll(line, ".", " . ")
line = strings.ReplaceAll(line, "<br />", " ")
line = strings.ReplaceAll(line, ",", " , ")
line = strings.ReplaceAll(line, "(", " ( ")
line = strings.ReplaceAll(line, ")", " ) ")
line = strings.ReplaceAll(line, "!", " ! ")
line = strings.ReplaceAll(line, "?", " ? ")
line = strings.ReplaceAll(line, ";", " ")
line = strings.ReplaceAll(line, ":", " ")
space := regexp.MustCompile(`\s+`)
line = space.ReplaceAllString(line, " ")
line = strings.ReplaceAll(line, "0", "@")
line = strings.ReplaceAll(line, "1", "@")
line = strings.ReplaceAll(line, "2", "@")
line = strings.ReplaceAll(line, "3", "@")
line = strings.ReplaceAll(line, "4", "@")
line = strings.ReplaceAll(line, "5", "@")
line = strings.ReplaceAll(line, "6", "@")
line = strings.ReplaceAll(line, "7", "@")
line = strings.ReplaceAll(line, "8", "@")
line = strings.ReplaceAll(line, "9", "@")

```

## Estimating Accuracy on Unseen

The standard way to do this is to split off part of the training
data and save it for testing. 

`P@1` and `R@1` are precision and recall for one label, namely country.
Because there is only one label, precision and recall are the same.
We can more call it in this case accuracy.

```bash
$ head -13000 training-ns.prn > training-ns-head.prn
$ tail -358 training-ns.prn > training-ns-tail.prn

$ time fasttext test model_goodness.bin training-ns-head.prn
replace later

N       13000
P@1     0.996
R@1     0.996
Number of examples: 13000

real    0m1.908s
user    0m0.015s
sys     0m0.094s

$ time fasttext test model_mapping-.bin training-ns-tail358.prn
N       357
P@1     0.969
R@1     0.969
Number of examples: 357

real    0m1.604s
user    0m0.015s
sys     0m0.094s

$ time fasttext predict-prob model_mapping.bin training-ns-predict.prn > prediction-prod-on-model.txt

real    0m2.153s
user    0m0.062s
sys     0m0.108s

```

## Making Predictions Using the Neural Network

```bash
fasttext predict-prob model_goodness.bin goodness.prn > ???
```

Here sample of the output file:

```bash
todo
```

If the probability is larger than one, as in 1.00001, that is a minor rounding
error in fastText. Don't worry about it.

## Assembling the Annual Money Laundering Report

We will simply append two columns to the input file. 

```bash
$ rowcut -c=1 urls.csv | \                      # cuts out the 2nd column, url text
  normalize -l=0 | \                            # normalizes the text
  fasttext predict-prob model_goodness.bin - | \ # run fasttext
  cut -c 10- \                                  # removes the __label__ prefix
  | tr " " ","  > predictions.csv               # makes a valid csv

# this combines two input CSV and the predictions CSV
$ rowpaste urls.csv predictions.csv  > urls-with-predictions.csv

```
