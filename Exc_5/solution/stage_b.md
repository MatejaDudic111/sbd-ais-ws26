stage b

-Options
    1. Keeping postgre and polling it better (several pollers for each machine part for example)
    2. gateways writing every reading into Kafka not postgre

-Comparison
    Option 1 (Postgre):
    -R1: not met, because all polls and writes go through one database
    -R2: not met, that gape from stage a between the commits and reads is still there even with several pollers
    -R3: maybe possible with enough pollers but hard
    -R4: not met every new consumer adds additional load to postgre

    Option 2:
    -R1: met, add partitions and detector instances
    -R2: met, there is machine id as the key so one machine always goes to the same partition in order
    -R3: met with enough instances of detector, the lag will show whether we are keeping up
    -R4: met, no extra load on postgre and every consumer gets his own group and offset

    -Decision: Options 2: the gateways write into kafka using the machine ids as key, and postgre is used for the machine registry. Option 2 meets all the Rs.