-- A) highest unit by price
select * from orders order by price_per_unit desc limit 1;
-- B) 3 highest categories by quantity
select product_category, sum(quantity) as total_quantity from orders group by product_category order by total_quantity desc limit 3;
-- C) the total revenue per each category
select product_category, sum(price_per_unit*quantity) as total_revenue from orders group by product_category order by total_revenue desc;
-- D) top 5 customers by total spending
select customer_name, sum(price_per_unit*quantity) as total_spending from orders group by customer_name order by total_spending desc limit 5;
-- E) here  we find the number of orders of the top 5 highest spending so we can extract data for the answer
select customer_name, count(*) as order_count from orders group by customer_name order by sum(price_per_unit*quantity) desc limit 5;