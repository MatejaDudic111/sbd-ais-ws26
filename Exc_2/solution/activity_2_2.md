--step 1
| Rows in table |      Join result       |   Time  |
|    50,000     |      27,501,822        | 1.181 s |
|    100,000    |      109,946,508       | 1.953 s |
|    200,000    |      439,395,606       | 7.596 s |

when the input doubles the join result grows by around 4 times meaning it is a quadratic growth pattern
if we run the same prediction for 1 million rows it is going to be around 11 billion join results
runtime also increases but not at a stable predictable rate.

--step 2
The index made the query 2 times slower while the number of results stayed the same, the execution plan used
parallel hash join and parallel index-only scans. the index did not improve performance.

--step 3

query without join
select sum(num * num) from (select country, count(*) as num from people_100k group by country) as counts;

results: 109,946,508 in 9.673 ms 
   
 
query with full table now
select sum(num * num) from (select country, count(*) as num from people_big group by country) as counts;

results: 10,983,941,260 in 100.987 ms 

The first query is the same as the original 100k query in terms of pairs but way faster

The second query is close to the predicted 11 billion pairs and is much faster than the original
because it counts rows per country and squares every count instead of using join on every pair

--step 4
indexes and hardware: they can improve perforamce but in our specific example redoing the query had a way bigger effect

if we needed pairs would computing power matter: because pairs grow in a quadratic way i dont think computers would solve the issue.
Computing power of course would help but it would be more of a bandaid than an actual solution, they would help the process but the pairs grow too fast

limitations of oltp: oltps are designed for smaller every day transactions and using it on large databases could slow down other users by using too many resources.
for very big databases it is better to uses something like a warehouse