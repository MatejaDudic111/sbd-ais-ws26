Stage A:
-Run 1 (20 machines, around 90s)
    -Rows in table readings: 18700
    -Rows read by the detector: 15959
    -2725 rows missing
    -last id that the detector remembered is 18700

-Are the missing readings in the table?
    Yes there is about 2700 rows in the table that the detector never read

-Why did the detector never see them?
    Because every row gets its ID when it is inserted while the session can see it only when commited. The detector asks for ID and remembers the highest that it has seen. for example if id 80 gets inserted but commit 200ms later, but at the same id 120 gets inserted and commits immediatelly what happens is that if the detector polls inbetween it only sees the id 120 and sets it as the highest id and the id 80 ,when it gets commited, is below 120 so the detector never reads them and loses them without realising

-Run 2 (full plant, 60 machines, around 60s)
    The plant sent 36,300 readings in 61 s
    -delay 12 s, backlog 7,000 rows
    -delay 23 s, backlog 13,500 rows
    -delay 30 s, backlog 17,900 rows

-What happens to delay, backlog and missing, and why?
    Delay and backlog kept growing while missing stayed at 41 after the first timestamp and didnt keep growing. what is happening here is that since the model needs to read 60 machines * 10 Hz which is 600 readings/s and for each it takes 2ms thats 1.8s for every 1s of data which causes backlog and delay to keep growing. R3 (anomalies within 5s) is violated here. for the missing staying at 41 what i think happened is that the detector falls very far behind and starts reading only old already commited rows so that gap from run 1 happens only in the beginning




run 1 and 2 evidence from both terminals:

[baseline] read   15,959  backlog       0 rows  delay   0.0 s  missing  2,725  anomalies 354

[baseline] stopped: read 15,959 readings, last id 18,700, 2,725 readings missing, 354 anomalies
(venv) PS C:\Users\admin\Desktop\sbd-ais-ws26\Exc_5\guardian> docker exec -it postgres psql -U postgres -d factory -c "SELECT count(*) FROM readings WHERE run = (SELECT max(run) FROM readings);"
 count 
-------
 18700
(1 row)

[baseline] stopped: read 16,674 readings, last id 37,233, 41 readings missing, 0 anomalies
  t=   90s     18,020 readings sent (200/s)

    t=   60s     36,060 readings sent (601/s)  faulty: M-029
run 1794411786: sent 36,300 readings in 61 s