### **[TestGetMember](./dbt/arst/ok.sql#L1)**
* **Query**: select * from members where memid = $1;
* **Args**: [1]
* **Error**: 
* **Columns**: [memid surname firstname address zipcode telephone recommendedby joindate]
* **Results**: [[1 Smith Darren 8 Bloomsbury Close, Boston 4321 555-555-5555  2012-07-02T12:02:05Z]]

### **[TestGetMemberA](./dbt/arst/ok.sql#L4)**
* **Query**: select * from members where memid = $1;
* **Args**: [A]
* **Error**: pq: invalid input syntax for type integer: "A"
* **Columns**: []
* **Results**: []

### **[TestGetMember2](./dbt/arst/ok.sql#L7)**
* **Query**: select * from members where memid = $1;
* **Args**: [2]
* **Error**: 
* **Columns**: [memid surname firstname address zipcode telephone recommendedby joindate]
* **Results**: [[2 Smith Tracy 8 Bloomsbury Close, New York 4321 555-555-5555  2012-07-02T12:08:23Z]]

### **[TestFacilities](./dbt/arst/ok.sql#L13)**
* **Query**: select * from facilities where facid = any($1);
* **Args**: [{1,2,3}]
* **Error**: 
* **Columns**: [facid name membercost guestcost initialoutlay monthlymaintenance]
* **Results**: [[1 Tennis Court 2 5 25 8000 200] [2 Badminton Court 0 15.5 4000 50] [3 Table Tennis 0 5 320 10]]


