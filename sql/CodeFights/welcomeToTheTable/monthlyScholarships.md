Students at your university get scholarships that are paid out throughout the year.

Information about the scholarships is stored in the table scholarships, which has the structure:

id: the unique student id;
scholarship: the amount of the annual scholarship the student has been awarded.
Now you need to calculate the amount of money each student should get per month. Given the table scholarships, build the resulting table as follows: The table should have the same columns as the initial table, but the scholarship column should contain the amount of the student's monthly scholarship payout. The rows should be ordered by the students' ids.

``` sql
CREATE PROCEDURE monthlyScholarships()
BEGIN
	SELECT id, (scholarship / 12) AS scholarship FROM scholarships;
END
```

1. We need to divide the scholarship by 12 to get the montly paiment and we want the return value to be call scholarship
