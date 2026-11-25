# Exercise 7: From the log to the lake

# Solution at the end

# Mateja Dudic



A batch job copies the Kafka topic `logs` into a data lake (Parquet files),
every few minutes. Each run **reads** the new records, **writes** them into
the lake, and **remembers** where it stopped.

**Your task:** choose how the job remembers where it stopped, implement it in
about five lines, and show that no record is ever missing or copied twice,
even if a run stops halfway.

## Part 0: setup

Stop the containers of Exercise 6 (`docker compose down -v` in that folder).
Then download the images (~5 GB, you need them next week too), start them,
and fill the topic with 2,000,000 log records:

```bash
docker compose --profile generator pull
docker compose up -d

docker exec kafka kafka-topics.sh --bootstrap-server localhost:9092 \
  --create --topic logs --partitions 4 --replication-factor 1

docker compose run --rm generator
```

Open <http://localhost:8888/?token=lab> and the notebook
`work/exercise7.ipynb`.

## Part 1 

1. **Run the example** (section 1 of the notebook). It copies the whole topic
   twice. Why does the checker fail?
2. **Choose an architecture:**
   - **A — State file:** the job saves where it stopped in a small file
     after writing.
   - **B — The lake is the truth:** the job finds where it stopped in the
     lake itself; it writes each run to a staging folder first, then moves
     it into the lake in one step.
   - **C — Remove duplicates when reading:** the lake may contain duplicates;
     readers use a view that keeps each record once.
3. **Implement** `run_once()` (section 3 of the notebook) with the building
   blocks listed there.
4. **Check** it (section 4): `check_all(run_once)` tests five situations,
   including runs that stop halfway. For C: `check_all(run_once, lake_sql=DEDUP_SQL)`.

The building blocks use Spark. You may use another technology if it meets
the same goal: no record missing, none twice, and the lake readable with
DuckDB (the checker in `tools/check_lake.py` works with any technology).

In `Exc_7/solution/README.md`, explain:

1. **Your decision:** which architecture and why
2. **Explain your result**
3. **The output of `check_all`.**

Add your notebook as `Exc_7/solution/exercise7.ipynb`. Be ready to present
your solution in about 5 minutes in the next exercise session (26.11.2026).

## Submission

Push `Exc_7/solution/` to your fork and add the link to the Moodle
submission. **Deadline: 25.11.2026, 18:00.**

<details>
<summary>Troubleshooting</summary>

- **Container name already in use / port already allocated:** the containers
  of Exercise 5 or 6 are still running: `docker compose down -v` in that
  folder.
- **The topic has 1 partition:** the generator ran before the topic was
  created. Delete it
  (`docker exec kafka kafka-topics.sh --bootstrap-server localhost:9092 --delete --topic logs`)
  and repeat Part 0.
- **`No module named ...` or `Failed to find data source: kafka`:** check
  your internet connection, restart the kernel and run the cells again.
- **Clean up:** `docker compose down` removes the containers and the topic.

</details>





# Results:


1. My decision: architecture B (the lake is the truth)

def run_once():
    # 1. where did the last run stop?

    start = lake_position()

    # 2. read what is new

    new = read_new(start)

    # 3. write it
    
    run = write_files(new, STAGING)

    # 4. remember where you stopped (if your architecture needs it)

    publish(run)

The idea is that we look where the job stopped by looking into the lake itself start = lake_position().

After that we read the new lines in the log new = read_new(start) 

and put them in a new staging folder run = write_files(new, STAGING). 

After which we write it from the staging folder into the lake publish(run).

The idea is that the staging folder makes all the files go into the lake together so that no bug or crash that leaves some files missing or duplicated can happen because if a crash happens nothing reaches the lake and the run gets discarded. 

Why i chose B: It is a consistent place of truth. if i chose A a crash between writing the data and saving the position would leave the two not in sync and would lead to duplicates or missing data compared to B where everything is in one place and changed together. C is more complex than B, B has no extra files or deduplication.

Trade off: We need to look into the lake every time to find its position. This gets slow the bigger the lake gets.



2. Explanation of results:

The example always read the whole topic from the beginning and thus every run copied again everything. Two runs made the lake have 4 million runs instead of 2 million. 2 million were duplicates. 

My method only read what came after the lake_position. 

- first run: all 2 million were copied since the lake position was empty

- second run: there was 100.000 new records that arrived and only those were copied so we had 2.100.000 now in the lake.

- third run: nothing new came nothing was copied and there were no duplicates.

- fourth and fifth run: run here stops while and after writing respectively. Here the run never reaches publishing and thus never reaches the lake. the 100.000 record are not loaded yet and the next run starts at the same position and copies them. results are correct no duplicates nothing missing again.

The idea is that the run either completely counts or its completely discarded.


3. Output of check_all():

situation                                 in lake  duplicates  missing  not loaded    run  check
-----------------------------------------------------------------------------------------------
1 first run                             2,000,000           0        0           0  16.6s  PASS

2 new records arrive, run again         2,100,000           0        0           0   5.3s  PASS

3 nothing new, run again                2,100,000           0        0           0   0.3s  PASS

4 run stops WHILE writing: right after  2,100,000           0        0     100,000         PASS

  ... then the next run                 2,200,000           0        0           0   4.7s  PASS

5 run stops AFTER writing: right after  2,200,000           0        0     100,000         PASS

  ... then the next run                 2,300,000           0        0           0   4.7s  PASS



