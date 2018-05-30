Your boss wants to identify the successful projects running in your company, so he asked you to prepare a list of all the currently active projects and their average monthly income.

You have stored the information about these projects in a simple database with a single Projects table that has five columns:

internal_id: the company's internal identifier for the project;
project_name: the official name of the project;
team_size: the number of employees working on the project;
team_lead: the name of the project manager;
income: the average monthly income of the project.
Your boss says that internal project ids are irrelevant to him and that he isn't interested in how big the teams are. Since that's the case, he wants you to create another table by removing the internal_id and team_size columns from the existing Projects table. Return it sorted by internal_id in ascending order.

```sql
CREATE PROCEDURE projectList()
BEGIN
	select project_name, team_lead, income from Projects;
END
```

1. Just need to select value need from the tables Projects
