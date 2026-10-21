-- step 1 smaller tables
create table people_50k as select * from people_big where id <= 50000;
create table people_100k as select * from people_big where id <= 100000;
create table people_200k as select * from people_big where id <= 200000;
-- step 1 join
select count(*) from people_50k p1 join people_50k p2 on p1.country = p2.country;

select count(*) from people_100k p1 join people_100k p2 on p1.country = p2.country;

select count(*) from people_200k p1 join people_200k p2 on p1.country = p2.country;
-- step 2 create index and analyze
create index index_people_100k_country on people_100k(country);
analyze people_100k;
-- step 2 explain analyze
explain analyze select count(*) from people_100k p1 join people_100k p2 on p1.country = p2.country;
-- step 3 query without join
select sum(num * num) from (select country, count(*) as num from people_100k group by country) as counts;
-- step 3 same with full table now
select sum(num * num) from (select country, count(*) as num from people_big group by country) as counts;
