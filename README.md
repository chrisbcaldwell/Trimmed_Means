# Trimmed_Means
Exercise performed using the trim package (github.com/chrisbcaldwell/trim)

## Project Overview
A comparison of the `trim` packace I developed and the same function in R was conducted.

## Running the Project

To run the Go portion, download the files `random.csv' and 'trimmed_means.exe' to the same folder and simply run the `trimmed_means.exe` file from the command line.

To run the R version, run the script `trimmed_means.R` by your preferred method; I used RStudio.

## Project Data

A file `random.csv` was created using Excel.  5,000 random outputs were generated.  For the field `integers`

## Go version

output:

```
Integer data trimmed mean: 69.60
Float data trimmed mean: 69.58
Total run time: 4.144ms
```

## R version

output:

```
[1] "Integer trimmed mean:"
[1] 69.60067
[1] "Float trimmed mean:"
[1] 69.57587
[1] "Total run time:"
Time difference of 0.02674413 secs
```
