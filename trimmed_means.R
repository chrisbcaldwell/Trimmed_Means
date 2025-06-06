# start the clock
begin <- Sys.time()

# set working directory to same as script:
setwd(dirname(parent.frame(2)$ofile))



df <- read.csv('random.csv')

q <- 0.05 # trimming quantile (from both ends)

print("Integer trimmed mean:")
print(mean(df$integers, trim = q))
print("Float trimmed mean:")
print(mean(df$floats, trim = q))
elapsed <- Sys.time() - begin
print("Total run time:")
print(elapsed)


