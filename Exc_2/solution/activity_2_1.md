--A
select * from orders order by price_per_unit desc limit 1;
 customer_name | product_category | quantity | price_per_unit | order_date | country 
---------------+------------------+----------+----------------+------------+---------
 Emma Brown    | Automotive       |        3 |        2000.00 | 2024-10-11 | Italy
(1 row)

--B
select product_category, sum(quantity) as total_quantity from orders group by product_category order by total_quantity desc limit 3;
 product_category | total_quantity 
------------------+----------------
 Health & Beauty  |         300842
 Electronics      |         300804
 Toys             |         300598
(3 rows)

--C
select product_category, sum(price_per_unit*quantity) as total_revenue from orders group by product_category order by total_revenue desc;
 product_category | total_revenue 
------------------+---------------
 Automotive       |  306589798.86
 Electronics      |  241525009.45
 Home & Garden    |   78023780.09
 Sports           |   61848990.83
 Health & Beauty  |   46599817.89
 Office Supplies  |   38276061.64
 Fashion          |   31566368.22
 Toys             |   23271039.02
 Grocery          |   15268355.66
 Books            |   12731976.04
(10 rows)

--D
select customer_name, sum(price_per_unit*quantity) as total_spending from orders group by customer_name order by total_spending desc limit 5;
 customer_name  | total_spending 
----------------+----------------
 Carol Taylor   |      991179.18
 Nina Lopez     |      975444.95
 Daniel Jackson |      959344.48
 Carol Lewis    |      947708.57
 Daniel Young   |      946030.14
(5 rows)

--E
select customer_name, count(*) as order_count from orders group by customer_name order by sum(price_per_unit*quantity) desc limit 5;
\ customer_name  | order_count 
----------------+-------------
 Carol Taylor   |        1028
 Nina Lopez     |         980
 Daniel Jackson |        1033
 Carol Lewis    |         943
 Daniel Young   |         973
(5 rows)

All the names have similar spending and orders which are way higher than normal.
in dataset_generator.py we can see that there is 32 first names and 33 last names combined randomly.
What causes these high numbers is that there is no unique ID to sort each person by, so all the different customers, randomly generated with the same name and surname, get grouped as one 
person with one spending amount.