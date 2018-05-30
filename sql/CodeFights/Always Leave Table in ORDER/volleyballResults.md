You are creating a website that will help you and your friends keep track of the results of volleyball teams from all around the world. Your website regularly crawls the web searching for new games, and adds the results of these games to the results table stored in your local database. After each update, the table should be sorted in ascending order by the total number of games won. This year's results are quite marvelous - at any given moment there are no two teams that have won the same number of games!

The results table contains the following columns:

name - the unique name of the team;
country - the country of the team;
scored - the number of scored goals;
missed - the number of missed goals;
wins - the total number of games the team has won.
Your task is to sort the given results table in ascending order by the number of wins.

``` sql
CREATE PROCEDURE volleyballResults()
BEGIN
	SELECT * FROM results ORDER BY wins;
END
```
